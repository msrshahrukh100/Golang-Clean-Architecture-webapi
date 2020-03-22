package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type muxRouter struct{}

// NewMuxRouter sdf
func NewMuxRouter() Router {
	return &muxRouter{}
}

var (
	newMuxrouter = mux.NewRouter()
)

func (*muxRouter) GET(path string, f func(resp http.ResponseWriter, req *http.Request)) {
	newMuxrouter.HandleFunc(path, f).Methods("GET")
}

func (*muxRouter) POST(path string, f func(resp http.ResponseWriter, req *http.Request)) {
	newMuxrouter.HandleFunc(path, f).Methods("POST")
}

func (*muxRouter) SERVE(port string) {
	log.Fatalln(http.ListenAndServe(port, newMuxrouter))
}
