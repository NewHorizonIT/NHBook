package middlewares

import (
	"log"
	"net/http"

	"github.com/NewHorizonIT/nhbook-api/pkg/errs"
	"github.com/gin-gonic/gin"
)

// HandleError handles different error types and returns appropriate HTTP response
func HandleError(c *gin.Context, err error) {
	if err == nil {
		return
	}

	// Type assertion for domain errors
	if domainErr, ok := err.(*errs.DomainError); ok {
		c.JSON(domainErr.StatusCode, gin.H{
			"success": false,
			"code":    domainErr.Code,
			"message": domainErr.Message,
			"details": domainErr.Details,
		})

		// Log for internal tracking
		if domainErr.Err != nil {
			log.Printf("[%s] %s: %v", domainErr.Severity, domainErr.Code, domainErr.Err)
		}
		return
	}

	// Default unknown error
	c.JSON(http.StatusInternalServerError, gin.H{
		"success": false,
		"code":    errs.CodeInternalServer,
		"message": "An unexpected error occurred",
	})
	log.Printf("[ERROR] Unknown error: %v", err)
}
