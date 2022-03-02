package service

import "github.com/kokhno-nikolay/lets-go-chat/internal/repository"

type ActiveUsersService struct {
	repo repository.ActiveUsers
}

func NewActiveUsersService(repo repository.ActiveUsers) *ActiveUsersService {
	return &ActiveUsersService{
		repo: repo,
	}
}

func (s *ActiveUsersService) Get() (int, error) {
	activeUsers, err := s.repo.GetAllActive()
	return len(activeUsers), err
}

func (s *ActiveUsersService) Add(userId int) error {
	return s.repo.SwitchToActive(userId)
}

func (s *ActiveUsersService) Remove(userId int) error {
	return s.repo.SwitchToInactive(userId)
}
