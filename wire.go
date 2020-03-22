//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/msrshahrukh100/Golang-Clean-Architecture-webapi/controller"
	"github.com/msrshahrukh100/Golang-Clean-Architecture-webapi/repository"
	"github.com/msrshahrukh100/Golang-Clean-Architecture-webapi/service"
)

// InitializeEvent ..
func InitializeEvent() *controller.PostController {
	wire.Build(
		wire.NewSet(
			wire.Bind(new(controller.Controller), new(*controller.PostController)),
			controller.NewPostController,
			// wire.Bind(new(service.Service), new(*service.PostService)),
			service.NewPostService,
			// wire.Bind(new(repository.Repository), new(*repository.PostRepository)),
			repository.NewFirestoreRepository,
		))

	return &controller.PostController{}
}
