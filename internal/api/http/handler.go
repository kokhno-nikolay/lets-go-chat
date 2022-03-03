package http

import (
	"github.com/kokhno-nikolay/lets-go-chat/internal/api/http/ws"
	"net/http"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"

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
		cors.Default(),
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

		claims, ok := h.services.JWT.ExtractClaims(token)
		if !ok {
			c.String(http.StatusInternalServerError, "jwt claims not found")
		}

		ws.Chat(c.Writer, c.Request, claims["uuid"].(string))
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
