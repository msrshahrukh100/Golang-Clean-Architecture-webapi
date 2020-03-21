package service

import (
	"errors"
	"math/rand"

	"github.com/msrshahrukh100/Golang-Clean-Architecture-webapi/entity"
	"github.com/msrshahrukh100/Golang-Clean-Architecture-webapi/repository"
)

// PostService sdf
type PostService interface {
	Validate(post *entity.Post) error
	FindAll() ([]entity.Post, error)
	Save(post *entity.Post) (*entity.Post, error)
}

type service struct{}

var (
	repo repository.PostRepository = repository.NewFirestoreRepository()
)

// NewPostService sdf
func NewPostService() PostService {
	return &service{}
}

// Validate
func (*service) Validate(post *entity.Post) error {
	if post == nil {
		return errors.New("Post is nil")
	}

	if post.Title == "" {
		return errors.New("Post title is empty")
	}
	return nil
}

// FindAll
func (*service) FindAll() ([]entity.Post, error) {
	return repo.FindAll()
}

// Save
func (*service) Save(post *entity.Post) (*entity.Post, error) {
	post.Id = rand.Int63()
	post, err := repo.Save(post)
	if err != nil {
		return nil, err
	}
	return post, nil
}
