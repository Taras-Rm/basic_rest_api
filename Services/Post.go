package services

import (
	"fmt"

	"github.com/Taras-Rm/basic_rest_api/Models"
	repositories "github.com/Taras-Rm/basic_rest_api/Repositories"
)

type PostService interface {
	CreatePost(id uint, post *Models.Post) (*Models.Post, error)
	GetPostsByUserId(id uint) ([]Models.Post, error)
	UpdatePost(id uint, post *Models.Post) error
	DeletePost(id uint) error
}

type postService struct {
	postRepository repositories.PostRepository
	userRepository repositories.UserRepository
}

func NewPostService(postRepo repositories.PostRepository, userRepo repositories.UserRepository) PostService {
	return &postService{
		postRepository: postRepo,
		userRepository: userRepo,
	}
}

func (s *postService) CreatePost(id uint, post *Models.Post) (*Models.Post, error) {
	user, err := s.userRepository.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	res, err := s.postRepository.CreatePost(user, post)
	if err != nil {
		return nil, err
	}
	return res, err
}

func (s *postService) GetPostsByUserId(id uint) ([]Models.Post, error) {
	_, err := s.userRepository.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	res, err := s.postRepository.GetPostsByUserId(id)
	if err != nil {
		return nil, err
	}
	return res, err
}

func (s *postService) UpdatePost(id uint, post *Models.Post) error {
	_, err := s.postRepository.GetPostById(id)
	if err != nil {
		return err
	}
	err = s.postRepository.UpdatePost(id, post)
	if err != nil {
		return err
	}
	return err
}

func (s *postService) DeletePost(id uint) error {
	_, err := s.postRepository.GetPostById(id)
	if err != nil {
		fmt.Println("del neisn post")
		return err
	}
	err = s.postRepository.DeletePost(id)
	if err != nil {
		return err
	}
	return nil
}
