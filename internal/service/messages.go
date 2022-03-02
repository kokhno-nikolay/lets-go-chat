package service

import "github.com/kokhno-nikolay/lets-go-chat/internal/repository"

type MessagesService struct {
	repo repository.Messages
}

func NewMessagesService(repo repository.Messages) *MessagesService {
	return &MessagesService{
		repo: repo,
	}
}

func (s *MessagesService) GetAllMessages() error {
	return nil
}

func (s *MessagesService) CreateNewMessage() error {
	return nil
}
