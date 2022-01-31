package service

import (
	"context"

	"github.com/kokhno-nikolay/lets-go-chat/internal/models"
	"github.com/kokhno-nikolay/lets-go-chat/internal/repository"
)

type UsersService struct {
	repo repository.Users
}

func NewUsersService(repo repository.Users) *UsersService {
	return &UsersService{
		repo: repo,
	}
}

func (u *UsersService) SignUp(ctx context.Context, user models.User) (string, error) {
	uuid, err := u.repo.Create(user)
	if err != nil {
		return "", err
	}

	return uuid, nil
}

func (u *UsersService) SignIn(ctx context.Context, user models.User) (string, error) {
	uuid, err := u.repo.GetUUID(user.Username)
	return uuid, err
}
