package repository

import (
	"github.com/msrshahrukh100/Golang-Clean-Architecture-webapi/entity"
)

// PostRepository for some use
type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}
