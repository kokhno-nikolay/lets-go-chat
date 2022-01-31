package service

import (
	"context"
	"github.com/kokhno-nikolay/lets-go-chat/internal/models"

	"github.com/kokhno-nikolay/lets-go-chat/internal/repository"
)

type UserInput struct {
	Username string
	Password string
}

type Users interface {
	SignUp(ctx context.Context, inp models.User) (string, error)
	SignIn(ctx context.Context, inp models.User) error
}

type Services struct {
	Users Users
}

type Deps struct {
	Repos *repository.Repositories
}

func NewServices(deps Deps) *Services {
	usersService := NewUsersService(deps.Repos.Users)

	return &Services{
		Users: usersService,
	}
}
