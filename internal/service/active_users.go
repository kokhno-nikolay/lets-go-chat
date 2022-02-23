package service

type ActiveUsersService struct {
	quantity int
}

func NewActiveUsersService(quantity int) *ActiveUsersService {
	return &ActiveUsersService{
		quantity: quantity,
	}
}

func (s *ActiveUsersService) Get() int {
	return s.quantity
}

func (s *ActiveUsersService) Add() {
	s.quantity++
}

func (s *ActiveUsersService) Remove() {
	s.quantity--
}
