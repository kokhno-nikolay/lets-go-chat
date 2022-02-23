package http

import (
	"github.com/kokhno-nikolay/lets-go-chat/internal/api/http/ws"
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

	router.HandleMethodNotAllowed = true

	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	ws := ws.NewWSHandler(h.services)
	router.GET("/ws", func(c *gin.Context) {
		token, ok := c.GetQuery("token")
		if !ok {
			c.String(http.StatusUnauthorized, "missing auth token")
			return
		}

		if _, err := h.services.JWT.ValidateToken(token); err != nil {
			c.String(http.StatusForbidden, "bad auth token")
			return
		}
		ws.Chat(c.Writer, c.Request)
	})

	h.initAPI(router)

	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	handlerV1 := v1.NewHandler(h.services)
	api := router.Group("/")
	{
		handlerV1.Init(api)
	}
}
