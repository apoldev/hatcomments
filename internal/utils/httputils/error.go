package httputils

import "github.com/gin-gonic/gin"

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ErrorProxyResponse struct {
	Error Error `json:"error"`
}

func ErrorResponse(c *gin.Context, code int, error Error) {
	c.JSON(code, ErrorProxyResponse{
		Error: error,
	})
}
