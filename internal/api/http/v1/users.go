package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/kokhno-nikolay/lets-go-chat/internal/config"
	"github.com/kokhno-nikolay/lets-go-chat/internal/models"
)

func (h *Handler) initUsersRoutes(api *gin.RouterGroup) {
	users := api.Group("/user")
	{
		users.POST("/", h.userSignUp)
		users.POST("/login", h.userSignIn)
		users.GET("/active", h.activeUsers)
	}
}

func (h *Handler) userSignUp(c *gin.Context) {
	var inp models.User
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, statusError, "invalid input body", nil)

		return
	}

	uuid, err := h.services.Users.SignUp(c, inp)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, statusError,
			fmt.Sprintf("internal server error. Error: %s", err.Error()),
			nil,
		)

		return
	}

	c.JSON(http.StatusCreated, models.SignUpResponse{UUID: uuid, Username: inp.Username})
}

func (h *Handler) userSignIn(c *gin.Context) {
	var inp models.User
	var cfg = config.GetConfig()

	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, statusError, "invalid input body", nil)

		return
	}

	uuid, err := h.services.Users.SignIn(c, inp)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, statusError,
			fmt.Sprintf("internal server error. Error: %s", err.Error()),
			nil,
		)

		return
	}

	jwsToken := h.services.JWT.GenerateToken(uuid)
	wsURL := fmt.Sprintf("ws://%s:%s/ws?token=%s", cfg.ServerHost, cfg.ServerPort, jwsToken)
	c.JSON(http.StatusOK, models.SignInResponse{URL: wsURL})
}

func (h *Handler) activeUsers(c *gin.Context) {
	c.String(http.StatusOK, strconv.Itoa(h.services.ActiveUsers.Get()))
}
