package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kokhno-nikolay/lets-go-chat/internal/models"
)

func (h *Handler) initUsersRoutes(api *gin.RouterGroup) {
	users := api.Group("/user")
	{
		users.POST("/", h.userSignUp)
		users.POST("/login", h.userSignIn)
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

	newResponse(c, http.StatusCreated, statusSuccess, "",
		models.SignUpResponse{UUID: uuid, Username: inp.Username},
	)
}

func (h *Handler) userSignIn(c *gin.Context) {
	var inp models.User
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

	newResponse(c, http.StatusCreated, statusSuccess, "",
		models.SignInResponse{UUID: uuid, URL: "http://random-url.com"},
	)
}
