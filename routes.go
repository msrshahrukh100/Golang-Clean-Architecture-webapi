package main

import (
	"encoding/json"
	"net/http"
	"github.com/msrshahrukh100/Golang-Clean-Architecture-webapi/entity"
	"github.com/msrshahrukh100/Golang-Clean-Architecture-webapi/repository"
	"math/rand"
)

var (
	repo repository.PostRepository = repository.NewPostRepository()
)



func getPosts(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")

	posts, err := repo.FindAll()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error getting the posts"}`))
		return
	}

	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(posts)

}


func addPost(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")

	var post entity.Post
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error marshalling posts"}`))
		return
	}

	post.Id = rand.Int63()
	repo.Save(&post)
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(post)

}