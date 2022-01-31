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
	Status  string `json:"status"`
	Message string `json:"message"`
}

func newResponse(c *gin.Context, statusCode int, status, message string) {
	log.Println(message)
	c.AbortWithStatusJSON(statusCode, response{status, message})
}

//register:
////	uuid
////	username
////
////login:
////	url
