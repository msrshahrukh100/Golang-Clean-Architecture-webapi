package controller

import (
	"encoding/json"
	"net/http"

	"github.com/msrshahrukh100/Golang-Clean-Architecture-webapi/errors"

	"github.com/msrshahrukh100/Golang-Clean-Architecture-webapi/entity"
	"github.com/msrshahrukh100/Golang-Clean-Architecture-webapi/service"
)

type postController struct{}

// PostController ...
type PostController interface {
	GetPosts(resp http.ResponseWriter, req *http.Request)
	AddPost(resp http.ResponseWriter, req *http.Request)
}

// NewPostController ..
func NewPostController() PostController {
	return &postController{}
}

var (
	postService service.PostService = service.NewPostService()
)

// GetPosts ..
func (*postController) GetPosts(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")

	posts, err := postService.FindAll()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error getting all posts"})
		return
	}

	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(posts)

}

// AddPost ..
func (*postController) AddPost(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")

	var post entity.Post
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error marshalling posts"})
		return
	}
	if err := postService.Validate(&post); err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: err.Error()})

		return
	}
	postService.Save(&post)
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(post)

}
