package utils

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
	Metadata   any    `json:"metadata"`
	Options    any    `json:"options"`
}

func WriteResponse(c *gin.Context, statusCode int, message string, metadata any, options any) {
	res := Response{
		StatusCode: statusCode,
		Message:    message,
		Metadata:   metadata,
		Options:    options,
	}

	c.JSON(statusCode, res)
}

func WriteError(c *gin.Context, statusCode int, message string) {
	WriteResponse(c, statusCode, message, nil, nil)
}
