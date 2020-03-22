package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/msrshahrukh100/Golang-Clean-Architecture-webapi/controller"
	"github.com/msrshahrukh100/Golang-Clean-Architecture-webapi/router"
)

// var (
// 	repo           repository.Repository = repository.NewFirestoreRepository()
// 	postService    service.Service       = service.NewPostService(repo)
// 	postController controller.Controller = controller.NewPostController(postService)
// )

var (
	postController *controller.PostController = InitializeEvent()
)

func main() {
	const port string = ":8000"
	httpRouter := router.NewMuxRouter()
	httpRouter.GET("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Hello world...")
	})

	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)

	log.Println("Server listening on port ", port)
	httpRouter.SERVE(port)
}
