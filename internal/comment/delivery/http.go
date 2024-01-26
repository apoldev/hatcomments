package delivery

import (
	"cloud_payments/internal/comment"
	"cloud_payments/internal/comment/usecase"
	"cloud_payments/internal/models"
	"cloud_payments/internal/utils"
	"cloud_payments/internal/utils/httputils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type CommentRepo interface {
	GetComments(limit, offset int) []models.Comment
	GetCommentByID(id uint) *models.Comment
	CreateComment(user models.User, room *models.Room, parent *models.Comment, text string, attachments []int) (*models.Comment, error)
	EditComment(id uint, text string, attachments []int) (*models.Comment, error)
	GetRoomOrCreate(int, string) *models.Room
	GetRoomByID(roomID int) *models.Room
	LoadCommentsByRoom(room *models.Room, limit, offset int) []models.Comment
	GetUploadedAttachment(file *utils.UploadFile) (*models.Attachment, bool)
	GetTenorAttachment(data *comment.TenorGIFData) *models.Attachment
	GetProject() *models.Project

	GetOrCreateUserVoteOnComment(user *models.User, comment2 *models.Comment, vote int) (*models.Vote, error)
	GetVotesByCommentID(id uint) (result []models.Vote, err error)

	DeleteComment(id int, user *models.User) error
	RestoreComment(int) error

	SaveCommentAction(action string, uid uint, old, newComment *models.Comment) error
}

type UserRepo interface {
	GetUserByID(int) (*models.User, error)

	CheckUserBan(u *models.User, req *http.Request) ([]models.UserBan, error)
}

type CommentHandler struct {
	CommentRepo   CommentRepo
	UserRepo      UserRepo
	CentrifugoApi *usecase.CentrifugoAPI
}

func (h *CommentHandler) CreateTenorGIFAttachment() func(c *gin.Context) {

	type TenorDataRequest struct {
		Title   string `json:"title"`
		ID      string `json:"id"`
		Gif     string `json:"gif"`
		NanoGif string `json:"nanogif"`
	}

	return func(c *gin.Context) {

		var req TenorDataRequest
		c.BindJSON(&req)

		if req.ID == "" {
			httputils.ErrorResponse(c, 200, httputils.Error{1007, "no id"})
			return
		}

		attachment := h.CommentRepo.GetTenorAttachment(&comment.TenorGIFData{
			ID:      req.ID,
			Title:   req.Title,
			Gif:     req.Gif,
			NanoGif: req.NanoGif,
		})

		c.JSON(200, attachment)
	}

}

func (h *CommentHandler) GetUploadHandler() func(c *gin.Context) {

	type ImageUploadForm struct {
		Name string                `form:"name" binding:"required"`
		File *multipart.FileHeader `form:"file" binding:"required"`
	}

	return func(c *gin.Context) {

		var req ImageUploadForm
		err := c.MustBindWith(&req, binding.FormMultipart)

		uploaded, err := utils.ReadUploadFile(req.File)

		if err != nil {
			c.Error(err)
			return
		}

		attachment, currentlyCreated := h.CommentRepo.GetUploadedAttachment(uploaded)
		// Если только что создано в базе - выполним запись в папку upload
		if currentlyCreated {
			utils.WriteUploadFile(uploaded)
		}

		if attachment == nil {
			c.JSON(404, gin.H{
				"error": "not found",
			})
			return
		}

		attachment.Format = uploaded.Format

		c.JSON(200, attachment)
	}
}

func (h *CommentHandler) GetUploadUrlHandler() func(c *gin.Context) {
	type Response struct {
		URL string `json:"url"`
	}
	return func(c *gin.Context) {
		c.JSON(200, Response{
			URL: os.Getenv("APP_URL") + "/upload",
		})
	}
}

func (h *CommentHandler) GetCommentHandler() func(c *gin.Context) {

	type Resp struct {
		Room     *models.Room              `json:"room"`
		Comments []*comment.ResultComments `json:"comments"`
	}

	return func(c *gin.Context) {

		idString := c.Param("id")
		ProjectId, err := strconv.Atoi(idString)
		if err != nil {
			c.JSON(400, gin.H{
				"error": "Bad request",
			})
			return
		}

		roomName := c.Request.URL.Query().Get("room")

		limit := utils.GetUint(c.Request.URL.Query().Get("limit"), 2)
		offset := utils.GetUint(c.Request.URL.Query().Get("offset"), 0)

		fmt.Println(limit, offset)

		if roomName == "" {
			c.JSON(400, gin.H{
				"error": "Bad request",
			})
			return
		}

		room := h.CommentRepo.GetRoomOrCreate(ProjectId, roomName)
		comments := h.CommentRepo.LoadCommentsByRoom(room, limit, offset)

		//t := time.Now()
		// 100-200 nanosecond
		result := []*comment.ResultComments{}
		for i := range comments {
			result = append(result, comment.CommentDTO{}.DTO(&comments[i]))
		}

		//fmt.Println("dto", time.Since(t).String())

		c.JSON(200, Resp{
			Room:     room,
			Comments: result,
		})
	}
}

func (h *CommentHandler) CommentsHandler() func(c *gin.Context) {

	return func(c *gin.Context) {

		var comments []models.Comment

		//s.store.Db.Model(&comment.Comment{}).Find(&comments)

		comments = h.CommentRepo.GetComments(50, 0)

		c.JSON(200, comments)
	}

}

func (h *CommentHandler) ConnectHandler() func(c *gin.Context) {

	return func(c *gin.Context) {

		ip := c.Request.Header.Get("X-Forwarded-For")
		fmt.Println(c.Request.Header, "ip:", ip)

		c.JSON(200, map[string]interface{}{
			"result": map[string]interface{}{
				"user": "anonymous_" + uuid.NewString(),
			},
		})

	}
}

func (h *CommentHandler) PublishHandler() func(c *gin.Context) {

	return func(c *gin.Context) {

		b, _ := io.ReadAll(c.Request.Body)

		var req comment.RequestProxy
		json.Unmarshal(b, &req)

		// todo сделать проверку куки и jwt

		fmt.Println("[headers]", c.Request.Header)

		// если анонимус
		if req.User == "" || strings.HasPrefix(req.User, "anonymous_") {
			httputils.ErrorResponse(c, 200, httputils.Error{1001, "unauthorized"})
			return
		}

		// Проверка юзера
		userID, _ := strconv.Atoi(req.User)
		if userID == 0 {
			httputils.ErrorResponse(c, 200, httputils.Error{1002, "error"})
			return
		}

		u, err := h.UserRepo.GetUserByID(userID)
		if err != nil {
			httputils.ErrorResponse(c, 200, httputils.Error{1003, "user not found"})
			return
		}

		userBans, err := h.UserRepo.CheckUserBan(u, c.Request)

		if len(userBans) > 0 {
			httputils.ErrorResponse(c, http.StatusOK, httputils.Error{9000, "banned"})
			return
		}

		action := ""

		d := strings.Split(req.Channel, ":")
		if len(d) == 2 {
			action = d[1]
		}

		if req.Channel == "send:restore" {
			var data comment.RequestDeleteComment
			json.Unmarshal(b, &data)

			com := h.CommentRepo.GetCommentByID(data.Data.CommentID)

			if com.UserID != u.ID && u.Role == "" {
				httputils.ErrorResponse(c, 200, httputils.Error{1004, "no access"})
				return
			}

			h.CommentRepo.RestoreComment(int(com.ID))
			h.CommentRepo.SaveCommentAction(action, u.ID, com, nil)

			go func() {

				room := h.CommentRepo.GetRoomByID(int(com.RoomID))

				finalData := comment.CentrifugoPublishData{
					Method: "restore",
					Data: comment.DeleteCommentPublishData{
						CommentID: com.ID,
					},
				}

				result, _ := json.Marshal(&finalData)
				h.CentrifugoApi.Publish("comments:room-"+room.Slug, result)

			}()

		}

		if req.Channel == "send:delete" {
			var data comment.RequestDeleteComment
			json.Unmarshal(b, &data)

			com := h.CommentRepo.GetCommentByID(data.Data.CommentID)

			if com.UserID != u.ID && u.Role == "" {
				httputils.ErrorResponse(c, 200, httputils.Error{1004, "no access"})
				return
			}

			h.CommentRepo.DeleteComment(int(com.ID), u)
			h.CommentRepo.SaveCommentAction(action, u.ID, com, nil)

			go func() {

				room := h.CommentRepo.GetRoomByID(int(com.RoomID))

				finalData := comment.CentrifugoPublishData{
					Method: "delete",
					Data: comment.DeleteCommentPublishData{
						CommentID:   com.ID,
						DeletedByID: u.ID,
					},
				}

				result, _ := json.Marshal(&finalData)
				h.CentrifugoApi.Publish("comments:room-"+room.Slug, result)

			}()

		}

		if req.Channel == "send:comment" {

			var data comment.RequestComment
			json.Unmarshal(b, &data)

			message := data.Data

			room := h.CommentRepo.GetRoomByID(message.RoomID)
			replyTo := h.CommentRepo.GetCommentByID(message.Parent)

			newComment, err := h.CommentRepo.CreateComment(*u, room, replyTo, message.Input, message.Attachments)

			if err != nil {
				httputils.ErrorResponse(c, 200, httputils.Error{1004, err.Error()})
				return
			}

			go func() {

				if replyTo != nil {
					// Найдем юзера родителя, если есть
					user, _ := h.UserRepo.GetUserByID(int(replyTo.UserID))

					if user != nil {
						replyTo.User = *user
						newComment.Parent = replyTo
					}
				}

				finalData := comment.CentrifugoPublishData{
					Method: "comment",
					Data:   comment.CommentDTO{}.DTO(newComment),
				}

				result, _ := json.Marshal(&finalData)

				h.CentrifugoApi.Publish("comments:room-"+room.Slug, result)

			}()

		}

		if req.Channel == "send:edit" {

			var data comment.RequestComment
			json.Unmarshal(b, &data)

			message := data.Data

			com := h.CommentRepo.GetCommentByID(message.CommentID)

			if com == nil {
				httputils.ErrorResponse(c, 200, httputils.Error{10012, "no access"})
				return
			}

			if com.UserID != u.ID {
				httputils.ErrorResponse(c, 200, httputils.Error{1009, "no access"})
				return
			}

			editedComment, err := h.CommentRepo.EditComment(message.CommentID, message.Input, message.Attachments)

			if err != nil {
				httputils.ErrorResponse(c, 200, httputils.Error{1010, err.Error()})
				return
			}

			h.CommentRepo.SaveCommentAction(action, u.ID, com, editedComment)

			room := h.CommentRepo.GetRoomByID(int(editedComment.RoomID))

			if err != nil {
				httputils.ErrorResponse(c, 200, httputils.Error{1011, err.Error()})
				return
			}

			go func() {

				dto := comment.CommentDTO{}.DTO(editedComment)

				finalData := comment.CentrifugoPublishData{
					Method: "edit",
					Data: comment.EditCommentPublishData{
						CommentID:   message.CommentID,
						Text:        editedComment.Text,
						Attachments: dto.Attachments,
					},
				}

				result, _ := json.Marshal(&finalData)

				h.CentrifugoApi.Publish("comments:room-"+room.Slug, result)

			}()

		}

		if req.Channel == "send:vote" {

			var voteReq comment.RequestCommentVote
			json.Unmarshal(b, &voteReq)

			room := h.CommentRepo.GetRoomByID(voteReq.Data.RoomID)
			com := h.CommentRepo.GetCommentByID(voteReq.Data.CommentID)

			newVote, err := h.CommentRepo.GetOrCreateUserVoteOnComment(u, com, voteReq.Data.Vote)

			if err != nil {
				httputils.ErrorResponse(c, 200, httputils.Error{1005, err.Error()})
				return
			}

			h.CommentRepo.SaveCommentAction(action, u.ID, com, &models.Comment{
				Votes: []models.Vote{*newVote},
			})

			go func() {

				// Загрузим голоса
				com.Votes, err = h.CommentRepo.GetVotesByCommentID(com.ID)

				if err != nil {
					httputils.ErrorResponse(c, 200, httputils.Error{1006, err.Error()})
					return
				}

				dto := comment.CommentDTO{}.DTO(com)

				finalData := comment.CentrifugoPublishData{
					Method: "vote",
					Data: comment.CommentPublishData{
						Likes:     dto.Like,
						CommentID: com.ID,
						Votes:     dto.Votes,
					},
				}
				result, _ := json.Marshal(&finalData)

				h.CentrifugoApi.Publish("comments:room-"+room.Slug, result)

			}()
		}

		publishResult(c)

	}
}

func publishResult(c *gin.Context) {
	c.JSON(200, map[string]interface{}{
		"result": map[string]interface{}{},
	})
}
