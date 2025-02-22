package repository

import (
	"handlersPetstore/internal/domain"
)

type userRepository struct{}

func NewUserRepository() domain.UserRepository {
	return &userRepository{}
}

func (r *userRepository) Create(user *domain.User) error {
	return nil
}

func (r *userRepository) CreateWithArray(users []domain.User) error {
	return nil
}

func (r *userRepository) Update(username string, user *domain.User) error {
	return nil
}

func (r *userRepository) Delete(username string) error {
	return nil
}

func (r *userRepository) GetByUsername(username string) (*domain.User, error) {
	return nil, nil
}

func (r *userRepository) Login(username, password string) (string, error) {
	return "", nil
}

func (r *userRepository) Logout() error {
	return nil
}
