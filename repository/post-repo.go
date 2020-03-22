package repository

import (
	"github.com/msrshahrukh100/Golang-Clean-Architecture-webapi/entity"
)

// Repository for some use
type Repository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}
