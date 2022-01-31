package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	v1 "github.com/kokhno-nikolay/lets-go-chat/internal/api/http/v1"
	"github.com/kokhno-nikolay/lets-go-chat/internal/config"
	"github.com/kokhno-nikolay/lets-go-chat/internal/service"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Init(cfg *config.Config) *gin.Engine {
	router := gin.Default()

	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	h.initAPI(router)

	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	handlerV1 := v1.NewHandler(h.services)
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}
