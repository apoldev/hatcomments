package comment

type RequestProxy struct {
	Client  string `json:"client"`
	User    string `json:"user"`
	Channel string `json:"channel"`
}

type Comment struct {
	ID      int    `json:"id"`
	RoomID  int    `json:"room_id"`
	Level   int    `json:"level"`
	ReplyTo int    `json:"reply_to"`
	Text    string `json:"text"`
	Path    string `json:"path"`
}

type CommentVoteData struct {
	RoomID    int `json:"room_id"`
	ProjectID int `json:"project_id"`

	CommentID uint `json:"comment_id"`
	Vote      int  `json:"vote"`
}

type RequestCommentVote struct {
	Data CommentVoteData `json:"data"`
}

type CommentData struct {
	RoomID      int    `json:"room_id"`
	ProjectID   int    `json:"project_id"`
	Parent      uint   `json:"parent"`
	Input       string `json:"input"`
	Attachments []int  `json:"attachments"`
	CommentID   uint   `json:"comment_id"`
}

type RequestComment struct {
	Data CommentData `json:"data"`
}

type CommentPublishData struct {
	CommentID uint         `json:"comment_id"`
	Likes     int          `json:"likes"`
	Votes     []ResultVote `json:"votes"`
}

type CentrifugoPublishData struct {
	Method string      `json:"method"`
	Data   interface{} `json:"data"`
}

// запрос к центрифуге для удаления
type DeleteCommentData struct {
	CommentID uint `json:"comment_id"`
}
type RequestDeleteComment struct {
	Data DeleteCommentData `json:"data"`
}

// Ответ
type DeleteCommentPublishData struct {
	CommentID   uint `json:"id"`
	DeletedByID uint `json:"deleted_by_id"`
}

// Ответ на измененный комментарий
type EditCommentPublishData struct {
	CommentID   uint               `json:"comment_id"`
	Text        string             `json:"text"`
	Attachments []ResultAttachment `json:"attachments"`
}
