package services

import (
	"Astral-Back-End/internal/repositories"
)

type UserService struct {
	UserRepository              *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{
		UserRepository:              userRepo,
	}
}

func (us *UserService) CreateUser(user *repositories.User) error {
	err := us.UserRepository.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (us *UserService) GetUserByEmail(email string) (*repositories.User, error) {
	return us.UserRepository.GetByEmail(email)
}
