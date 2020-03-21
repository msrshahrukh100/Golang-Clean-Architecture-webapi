package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/msrshahrukh100/Golang-Clean-Architecture-webapi/controller"
)

func main() {
	router := mux.NewRouter()
	const port string = ":8000"
	router.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Hello world...")
	})

	router.HandleFunc("/posts", controller.GetPosts).Methods("GET")
	router.HandleFunc("/posts", controller.AddPost).Methods("POST")

	log.Println("Server listening on port ", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
