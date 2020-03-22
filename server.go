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
	httpRouter     router.Router             = router.NewMuxRouter()
)

func main() {
	const port string = ":8000"
	httpRouter.GET("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Hello world...")
	})

	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)

	log.Println("Server listening on port ", port)
	httpRouter.SERVE(port)
}
