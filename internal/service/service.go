package service

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/kokhno-nikolay/lets-go-chat/internal/config"

	"github.com/kokhno-nikolay/lets-go-chat/internal/models"
	"github.com/kokhno-nikolay/lets-go-chat/internal/repository"
)

type UserInput struct {
	Username string
	Password string
}

type Users interface {
	SignUp(ctx context.Context, inp models.User) (string, error)
	SignIn(ctx context.Context, inp models.User) (models.User, error)
}

type ActiveUsers interface {
	Get() (int, error)
	Add(userUUID string) error
	Remove(userUUID string) error
}

type JWT interface {
	GenerateToken(uuid, username string) string
	ValidateToken(token string) (*jwt.Token, error)
	ExtractClaims(tokenStr string) (jwt.MapClaims, bool)
}

type Messages interface {
	GetAllMessages() error
	CreateNewMessage() error
}

type Services struct {
	Users       Users
	ActiveUsers ActiveUsers
	JWT         JWT
	Messages    Messages
}

type Deps struct {
	Repos *repository.Repositories
}

func NewServices(deps Deps) *Services {
	usersService := NewUsersService(deps.Repos.Users)
	activeUsersService := NewActiveUsersService(deps.Repos.ActiveUsers)
	jwtService := NewJWTService(deps.Repos.JWT, config.GetConfig().SecretKey)
	messagesService := NewMessagesService(deps.Repos.Messages)

	return &Services{
		Users:       usersService,
		ActiveUsers: activeUsersService,
		JWT:         jwtService,
		Messages:    messagesService,
	}
}
