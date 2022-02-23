package ws

import (
	"github.com/kokhno-nikolay/lets-go-chat/internal/service"
)

type WSHandler struct {
	services *service.Services
}

func NewWSHandler(services *service.Services) *WSHandler {
	return &WSHandler{
		services: services,
	}
}
