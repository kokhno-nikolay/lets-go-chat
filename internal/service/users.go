package service

import "context"

type UsersService struct {
}

func NewUsersService() *UsersService {
	return &UsersService{}
}

func (u *UsersService) SignUp(ctx context.Context, input UserInput) error {
	return nil
}

func (u *UsersService) SignIn(ctx context.Context, input UserInput) error {
	return nil
}
