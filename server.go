package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"log"
)

func main() {
	router := mux.NewRouter()
	const port string = ":8000"
	router.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request){
		fmt.Fprintln(resp, "Hello world...")
	})

	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/posts", addPost).Methods("POST")

	log.Println("Server listening on port ", port)
	log.Fatalln(http.ListenAndServe(port, router))
}