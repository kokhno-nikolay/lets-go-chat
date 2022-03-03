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

func (s *ActiveUsersService) Add(userUUID string) error {
	return s.repo.SwitchToActive(userUUID)
}

func (s *ActiveUsersService) Remove(userUUID string) error {
	return s.repo.SwitchToInactive(userUUID)
}
