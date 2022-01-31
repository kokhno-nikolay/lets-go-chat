package service

import "context"

type UserInput struct {
	Username string
	Password string
}

type Users interface {
	SignUp(ctx context.Context, input UserInput) error
	SignIn(ctx context.Context, input UserInput) error
}

type Services struct {
	Users Users
}

func NewServices() *Services {
	usersService := NewUsersService()

	return &Services{
		Users: usersService,
	}
}
