package v1

import (
	"log"

	"github.com/gin-gonic/gin"
)

const (
	statusSuccess = "success"
	statusError   = "error"
)

type response struct {
	Status  string      `json:"status,omitempty"`
	Message string      `json:"message,omitempty"`
	Body    interface{} `json:"body,omitempty"`
}

func newResponse(c *gin.Context, statusCode int, status, message string, body interface{}) {
	log.Println(message)
	c.AbortWithStatusJSON(statusCode, response{status, message, body})
}
