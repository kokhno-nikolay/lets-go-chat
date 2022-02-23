package service

import (
	"context"
	"github.com/dgrijalva/jwt-go"

	"github.com/kokhno-nikolay/lets-go-chat/internal/models"
	"github.com/kokhno-nikolay/lets-go-chat/internal/repository"
)

// TODO: move to config
const SECRET_KEY = "secret"

type UserInput struct {
	Username string
	Password string
}

type Users interface {
	SignUp(ctx context.Context, inp models.User) (string, error)
	SignIn(ctx context.Context, inp models.User) (string, error)
}

type ActiveUsers interface {
	Get() int
	Add()
	Remove()
}

type JWT interface {
	GenerateToken(uuid string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type Services struct {
	Users       Users
	ActiveUsers ActiveUsers
	JWT         JWT
}

type Deps struct {
	Repos *repository.Repositories
}

func NewServices(deps Deps) *Services {
	usersService := NewUsersService(deps.Repos.Users)
	activeUsersService := NewActiveUsersService(0)
	jwtService := NewJWTService(deps.Repos.JWT, SECRET_KEY)

	return &Services{
		Users:       usersService,
		ActiveUsers: activeUsersService,
		JWT:         jwtService,
	}
}
