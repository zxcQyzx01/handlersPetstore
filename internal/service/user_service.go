package service

import (
	"handlersPetstore/internal/domain"
)

type UserService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user *domain.User) error {
	return s.repo.Create(user)
}

func (s *UserService) CreateUsersWithArray(users []domain.User) error {
	return s.repo.CreateWithArray(users)
}

func (s *UserService) UpdateUser(username string, user *domain.User) error {
	return s.repo.Update(username, user)
}

func (s *UserService) DeleteUser(username string) error {
	return s.repo.Delete(username)
}

func (s *UserService) GetUserByUsername(username string) (*domain.User, error) {
	return s.repo.GetByUsername(username)
}

func (s *UserService) Login(username, password string) (string, error) {
	return s.repo.Login(username, password)
}

func (s *UserService) Logout() error {
	return s.repo.Logout()
}
