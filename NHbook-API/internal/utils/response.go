package utils

import (
	"github.com/gin-gonic/gin"
)

type ResponseSuccess struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
	Metadata   any    `json:"metadata"`
	Options    any    `json:"options"`
}

type ResponseError struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
}

func WriteResponse(c *gin.Context, statusCode int, message string, metadata any, options any) {
	res := ResponseSuccess{
		Message:    message,
		StatusCode: statusCode,
		Metadata:   metadata,
		Options:    options,
	}

	c.JSON(statusCode, res)
}

func WriteError(c *gin.Context, statusCode int, message string) {
	res := ResponseError{
		StatusCode: statusCode,
		Message:    message,
	}

	c.JSON(statusCode, res)
}
