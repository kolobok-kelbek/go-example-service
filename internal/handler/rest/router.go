package rest

import (
	"fmt"
	"net/http"
)

//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=doc/config.yaml doc/api.yaml

type Method string

const GET Method = "GET"
const POST Method = "POST"
const PUT Method = "PUT"
const PATCH Method = "PATCH"
const DELETE Method = "DELETE"
const HEAD Method = "HEAD"
const CONNECT Method = "CONNECT"
const OPTIONS Method = "OPTIONS"
const TRACE Method = "TRACE"

type Router struct {
	*http.ServeMux
}

func NewRouter() *Router {
	return &Router{
		ServeMux: http.NewServeMux(),
	}
}

func (r *Router) endpoint(method Method, path string, handlerFunc http.HandlerFunc) {
	r.Handle(fmt.Sprintf("%s %s", method, path), handlerFunc)
}

func (r *Router) get(path string, handlerFunc http.HandlerFunc) {
	r.endpoint(GET, path, handlerFunc)
}

func (r *Router) post(path string, handlerFunc http.HandlerFunc) {
	r.endpoint(POST, path, handlerFunc)
}

func (r *Router) put(path string, handlerFunc http.HandlerFunc) {
	r.endpoint(PUT, path, handlerFunc)
}

func (r *Router) patch(path string, handlerFunc http.HandlerFunc) {
	r.endpoint(PATCH, path, handlerFunc)
}

func (r *Router) delete(path string, handlerFunc http.HandlerFunc) {
	r.endpoint(DELETE, path, handlerFunc)
}

func NewRoute(
	login *LoginHandler,
	logout *LogoutHandler,
	registration *RegistrationHandler,
) *http.ServeMux {
	router := NewRouter()

	router.post("/login", login.Handle)
	router.post("/logout", logout.Handle)
	router.post("/registration", registration.Handle)

	return router.ServeMux
}
