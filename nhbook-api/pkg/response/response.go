package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Define common response structures
type Metadata struct {
	RequestID string `json:"request_id"`
	Timestamp string `json:"timestamp"`
}

type SuccessResponse struct {
	Metadata Metadata `json:"metadata"`
	Data     any      `json:"data"`
	Success  bool     `json:"success"`
}

type ErrorResponse struct {
	Metadata Metadata `json:"metadata"`
	Message  string   `json:"message"`
	Code     string   `json:"code"`
	Error    any      `json:"error"`
	Success  bool     `json:"success"`
}

var httpStatusMapping = map[string]int{
	"VALIDATION_ERR":        http.StatusBadRequest,
	"UNAUTHORIZED":          http.StatusUnauthorized,
	"FORBIDDEN":             http.StatusForbidden,
	"NOT_FOUND":             http.StatusNotFound,
	"INTERNAL_SERVER_ERROR": http.StatusInternalServerError,
}

// Write response with gin
func WriteSuccessResponse(ctx *gin.Context, statusCode int, data interface{}) {
	ctx.JSON(statusCode, gin.H{
		"success": true,
		"data":    data,
	})
}

func WriteErrorResponse(ctx *gin.Context, statusCode int, message string, err any) {
	if err == nil {
		err = ""
	}
	ctx.JSON(statusCode, gin.H{
		"success": false,
		"message": message,
		"error":   err,
	})
}

// Convert function
func GetStatusCode(code string) int {
	if statusCode, exists := httpStatusMapping[code]; exists {
		return statusCode
	}
	return http.StatusInternalServerError
}
