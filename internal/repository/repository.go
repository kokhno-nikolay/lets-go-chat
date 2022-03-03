package repository

import (
	"database/sql"

	"github.com/kokhno-nikolay/lets-go-chat/internal/models"
	"github.com/kokhno-nikolay/lets-go-chat/internal/repository/postgres"
)

type Users interface {
	GetUUID(username string) (string, error)
	Create(user models.User) (string, error)
}

type ActiveUsers interface {
	GetAllActive() ([]models.User, error)
	SwitchToActive(userUUID string) error
	SwitchToInactive(userUUID string) error
}

type JWT interface {
}

type Messages interface {
	GetAll() ([]models.Message, error)
	Create(message models.Message) error
}

type Repositories struct {
	Users       Users
	ActiveUsers ActiveUsers
	JWT         JWT
	Messages    Messages
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		Users:       postgres.NewUserRepo(db),
		ActiveUsers: postgres.NewUserRepo(db),
		JWT:         postgres.NewJWTRepo(db),
		Messages:    postgres.NewMessageRepo(db),
	}
}
