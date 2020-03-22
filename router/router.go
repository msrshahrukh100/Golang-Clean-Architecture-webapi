package router

import "net/http"

// Router sfd
type Router interface {
	GET(path string, f func(resp http.ResponseWriter, req *http.Request))
	POST(path string, f func(resp http.ResponseWriter, req *http.Request))
	SERVE(port string)
}
