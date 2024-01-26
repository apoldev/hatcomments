package comment

import (
	"cloud_payments/internal/models"
	"net/url"
	"os"
	"strings"
	"time"
)

type ResultUser struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Image     string `json:"image"`
	Role      string `json:"role"`
	Icon      string `json:"icon"`
	Status    string `json:"status"`

	SocialType string `json:"social_type"`
	SocialID   string `json:"social_id"`
}

type ResultAttachment struct {
	ID      uint   `json:"id"`
	Hash    string `json:"hash"`
	Type    string `json:"type"`
	Alt     string `json:"alt"`
	Name    string `json:"name"`
	Preview string `json:"preview"`
	Format  string `json:"format"`
}

type ResultVote struct {
	ID     uint `json:"id"`
	Vote   int  `json:"vote"`
	UserID uint `json:"user_id"`
}

type ResultComments struct {
	ID          uint         `json:"id"`
	RoomID      uint         `json:"room_id"`
	Level       uint         `json:"level"`
	ReplyTo     uint         `json:"reply_to"`
	Text        string       `json:"text"`
	Path        string       `json:"path"`
	CreatedAt   time.Time    `json:"created_at"`
	Deleted     bool         `json:"deleted"`
	DeletedByID *uint        `json:"deleted_by_id"`
	Like        int          `json:"like"`
	Votes       []ResultVote `json:"votes"`

	User        ResultUser         `json:"user"`
	Attachments []ResultAttachment `json:"attachments"`
	ReplyToUser *ResultUser        `json:"reply_to_user"`
}

type CommentDTO struct {
}

func (d CommentDTO) DTO(c *models.Comment) *ResultComments {

	if c == nil {
		return nil
	}

	level := uint(0)
	replyTo := uint(0)

	if c.Level != nil {
		level = *c.Level
	}

	if c.ReplyTo != nil {
		replyTo = *c.ReplyTo
	}

	// подсчет лайков
	c.Like = 0
	if len(c.Votes) > 0 {
		for v := range c.Votes {
			c.Like += c.Votes[v].Vote
		}
	}

	attachments := []ResultAttachment{}

	for i := range c.Attachments {

		f := strings.Split(c.Attachments[i].Name, ".")
		if len(f) == 0 {
			continue
		}

		format := f[len(f)-1]

		attachments = append(attachments, ResultAttachment{
			ID:      c.Attachments[i].ID,
			Hash:    c.Attachments[i].Hash,
			Type:    c.Attachments[i].Type,
			Alt:     c.Attachments[i].Alt,
			Name:    c.Attachments[i].Name,
			Preview: c.Attachments[i].Preview,
			Format:  format,
		})
	}

	votes := []ResultVote{}

	for i := range c.Votes {

		votes = append(votes, ResultVote{
			ID:     c.Votes[i].ID,
			Vote:   c.Votes[i].Vote,
			UserID: c.Votes[i].UserID,
		})
	}

	image := c.User.Image

	if image == "" {
		name := strings.TrimSpace(c.User.FirstName + " " + c.User.LastName)
		image = os.Getenv("APP_URL") + "/user/avatar?name=" + url.PathEscape(name)
	}

	var resultUser *ResultUser = nil
	if c.Parent != nil {
		p := d.DTO(c.Parent)
		resultUser = &ResultUser{
			ID:        p.User.ID,
			FirstName: p.User.FirstName,
			LastName:  p.User.LastName,
			Image:     p.User.Image,
		}
	}

	var deletedByID *uint = nil
	if c.DeletedByID != 0 {
		deletedByID = &c.DeletedByID
	}

	return &ResultComments{
		ID:          c.ID,
		RoomID:      c.RoomID,
		Level:       level,
		ReplyTo:     replyTo,
		Text:        c.Text,
		Path:        c.Path,
		CreatedAt:   c.CreatedAt,
		Deleted:     c.DeletedAt.Valid,
		DeletedByID: deletedByID,
		Like:        c.Like,
		Votes:       votes,
		ReplyToUser: resultUser,

		User: ResultUser{
			ID:         c.User.ID,
			FirstName:  c.User.FirstName,
			LastName:   c.User.LastName,
			Image:      image,
			Role:       c.User.Role,
			Icon:       c.User.Icon,
			Status:     c.User.Status,
			SocialID:   c.User.SocialID,
			SocialType: c.User.Type,
		},

		Attachments: attachments,
	}

}
