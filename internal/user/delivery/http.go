package delivery

import (
	"cloud_payments/internal/models"
	"cloud_payments/internal/user"
	"cloud_payments/internal/user/usecase"
	"cloud_payments/internal/utils"
	"cloud_payments/internal/utils/httputils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const (
	UserGoogle    = "google"
	UserVK        = "vk"
	UserAnonymous = "anonymous"
)

type UserRepo interface {
	GetUsers(limit, offset int) []models.User

	RegisterAnonimous(string, string) (*models.User, error)
	RegisterSocial(data *user.SocialData, typ string) (*models.User, error)

	GetUserByID(id int) (*models.User, error)

	NewUserVisitLog(u *models.User, req *http.Request) error
	CheckUserBan(u *models.User, req *http.Request) ([]models.UserBan, error)
}

type UserHandler struct {
	UserRepo   UserRepo
	JwtManager usecase.JwtManager
}

func (h *UserHandler) GetUsers() func(c *gin.Context) {
	return func(c *gin.Context) {

		c.JSON(200, h.UserRepo.GetUsers(50, 0))
	}
}

func (h *UserHandler) AuthUserHandler() func(c *gin.Context) {

	type ReqAuth struct {
		Token string `json:"token"`
	}

	type UserBans struct {
		BannedForHours float64   `json:"banned_for_hours,omitempty""`
		BannedUntil    time.Time `json:"banned_until"`
		BannedBy       uint      `json:"banned_by"`
		IP             string    `json:"ip,omitempty"`
		UserID         uint      `json:"user_id,omitempty"`
	}

	type RespInfo struct {
		ID        uint       `json:"id"`
		FirstName string     `json:"first_name"`
		LastName  string     `json:"last_name"`
		Image     string     `json:"image"`
		Role      string     `json:"role"`
		UserBans  []UserBans `json:"user_bans"`
	}

	type RespAuth struct {
		Token string   `json:"token"`
		Info  RespInfo `json:"info"`
	}

	return func(c *gin.Context) {

		var req ReqAuth
		c.BindJSON(&req)

		jwtData, err := h.JwtManager.GetJwtData(req.Token)

		u, err := h.UserRepo.GetUserByID(int(jwtData.Info.ID))

		if err != nil {
			httputils.ErrorResponse(c, http.StatusBadRequest, httputils.Error{2000, err.Error()})
			return
		}

		_, token, err := h.JwtManager.SignJwt(u)

		if err != nil {
			fmt.Println("err", err)
		}

		userBans, _ := h.UserRepo.CheckUserBan(u, c.Request)

		bans := make([]UserBans, len(userBans))
		for i := range userBans {
			bans[i] = UserBans{
				BannedForHours: time.Until(userBans[i].BannedUntil).Hours(),
				BannedUntil:    userBans[i].BannedUntil,
				BannedBy:       userBans[i].BannedBy,
				IP:             userBans[i].IP,
				UserID:         userBans[i].UserID,
			}
		}

		err = h.UserRepo.NewUserVisitLog(u, c.Request)

		if err != nil {
			httputils.ErrorResponse(c, http.StatusBadRequest, httputils.Error{2000, err.Error()})
			return
		}

		c.JSON(200, RespAuth{
			Token: token,
			Info: RespInfo{
				ID:        u.ID,
				FirstName: u.FirstName,
				LastName:  u.LastName,
				Image:     u.Image,
				Role:      u.Role,
				UserBans:  bans,
			},
		})

	}

}

func (h *UserHandler) RegisterUserHandler() func(c *gin.Context) {

	type RequestRegister struct {
		Name        string `json:"name"`
		GoogleToken string `json:"google_token"`
		VkToken     string `json:"vk_token"`
		Email       string `json:"email"`
	}

	type ResponseRegister struct {
		Token string `json:"token"`
	}

	return func(c *gin.Context) {

		var req RequestRegister
		c.BindJSON(&req)

		var u *models.User
		var err error

		if req.GoogleToken != "" {

			gData, err := usecase.GoogleData(req.GoogleToken)

			if err != nil {
				httputils.ErrorResponse(c, http.StatusBadRequest, httputils.Error{9000, "invalid token"})
				return
			}

			u, err = h.UserRepo.RegisterSocial(gData, UserGoogle)

			if err != nil {
				fmt.Println("err", err)
			}

		}

		if req.VkToken != "" {

			socialDataVK, err := usecase.VkData(req.VkToken, req.Email)

			if err != nil {
				httputils.ErrorResponse(c, http.StatusBadRequest, httputils.Error{9000, "invalid token"})
				return
			}

			u, err = h.UserRepo.RegisterSocial(socialDataVK, UserVK)

			if err != nil {
				fmt.Println("err", err)
			}

		}

		if req.Name != "" {

			if len(req.Name) < 2 {
				httputils.ErrorResponse(c, http.StatusBadRequest, httputils.Error{9000, "short name"})
				return
			}

			u, err = h.UserRepo.RegisterAnonimous(req.Name, UserAnonymous)

			if err != nil {
				fmt.Println("err", err)
			}

		}

		_, token, err := h.JwtManager.SignJwt(u)

		if err != nil {
			fmt.Println("err", err)
		}

		c.JSON(200, ResponseRegister{
			Token: token,
		})
	}

}

func (h *UserHandler) GetSvgAvatarByName() func(c *gin.Context) {

	return func(c *gin.Context) {
		b := []byte(utils.GetSvgAvatarByText(c.Query("name")))
		c.Data(200, "image/svg+xml", b)
	}
}
