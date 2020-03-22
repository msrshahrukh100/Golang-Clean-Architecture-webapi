package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/msrshahrukh100/Golang-Clean-Architecture-webapi/controller"
	"github.com/msrshahrukh100/Golang-Clean-Architecture-webapi/repository"
	"github.com/msrshahrukh100/Golang-Clean-Architecture-webapi/router"
	"github.com/msrshahrukh100/Golang-Clean-Architecture-webapi/service"
)

var (
	repo           repository.PostRepository = repository.NewFirestoreRepository()
	postService    service.PostService       = service.NewPostService(repo)
	postController controller.PostController = controller.NewPostController(postService)
)

func main() {
	router := router.NewMuxRouter()
	const port string = ":8000"
	router.GET("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Hello world...")
	})

	router.GET("/posts", postController.GetPosts)
	router.POST("/posts", postController.AddPost)

	log.Println("Server listening on port ", port)
	router.SERVE(port)
}
