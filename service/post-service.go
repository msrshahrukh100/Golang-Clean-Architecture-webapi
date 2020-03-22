package service

import (
	"errors"
	"math/rand"

	"github.com/msrshahrukh100/Golang-Clean-Architecture-webapi/entity"
	"github.com/msrshahrukh100/Golang-Clean-Architecture-webapi/repository"
)

// Service sdf
type Service interface {
	Validate(post *entity.Post) error
	FindAll() ([]entity.Post, error)
	Save(post *entity.Post) (*entity.Post, error)
}

// PostService ..
type PostService struct{}

var (
	repo repository.Repository
)

// NewPostService sdf
func NewPostService(r repository.Repository) Service {
	repo = r
	return &PostService{}
}

// Validate ..
func (*PostService) Validate(post *entity.Post) error {
	if post == nil {
		return errors.New("Post is nil")
	}

	if post.Title == "" {
		return errors.New("Post title is empty")
	}
	return nil
}

// FindAll ..
func (*PostService) FindAll() ([]entity.Post, error) {
	return repo.FindAll()
}

// Save ..
func (*PostService) Save(post *entity.Post) (*entity.Post, error) {
	post.Id = rand.Int63()
	post, err := repo.Save(post)
	if err != nil {
		return nil, err
	}
	return post, nil
}
