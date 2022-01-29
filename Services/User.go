package services

import (
	"github.com/Taras-Rm/basic_rest_api/Models"
	repositories "github.com/Taras-Rm/basic_rest_api/Repositories"
)

type UserService interface {
	GetUsers() ([]Models.User, error)
	CreateUser(user *Models.User) (*Models.User, error)
	GetUserByID(id uint) (*Models.User, error)
	UpdateUser(id uint, user *Models.User) error
	DeleteUser(id uint) error
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepository: userRepo}
}

func (s *userService) GetUsers() ([]Models.User, error) {
	res, err := s.userRepository.GetUsers()
	if err != nil {
		return nil, err
	}
	return res, err
}

func (s *userService) CreateUser(user *Models.User) (*Models.User, error) {
	res, err := s.userRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return res, err
}

func (s *userService) GetUserByID(id uint) (*Models.User, error) {
	res, err := s.userRepository.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return res, err
}

func (s *userService) UpdateUser(id uint, user *Models.User) error {
	_, err := s.GetUserByID(id)
	if err != nil {
		return err
	}
	err = s.userRepository.UpdateUser(id, user)
	if err != nil {
		return err
	}
	return nil
}

func (s *userService) DeleteUser(id uint) error {
	_, err := s.userRepository.GetUserByID(id)
	if err != nil {
		return err
	}
	err = s.userRepository.DeleteUserPosts(id)
	if err != nil {
		return err
	}
	err = s.userRepository.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}
