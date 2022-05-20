package server

import (
	"net/http"
)

type Router struct {
	Base   string
	routes []Route
}

func (r *Router) Route(method string, path string, handlers ...http.HandlerFunc) {
	route := Route{
		Method: method,
		Path:   r.Base + path,
		Handler: func(w http.ResponseWriter, r *http.Request) {
			for _, handler := range handlers {
				handler(w, r)
			}
		},
	}

	r.routes = append(r.routes, route)
}

func (r *Router) Get(path string, handler ...http.HandlerFunc) {
	r.Route(http.MethodGet, path, handler...)
}
func (r *Router) Post(path string, handler ...http.HandlerFunc) {
	r.Route(http.MethodPost, path, handler...)
}
func (r *Router) Put(path string, handler ...http.HandlerFunc) {
	r.Route(http.MethodPut, path, handler...)
}
func (r *Router) Patch(path string, handler ...http.HandlerFunc) {
	r.Route(http.MethodPatch, path, handler...)
}
func (r *Router) Delete(path string, handler ...http.HandlerFunc) {
	r.Route(http.MethodDelete, path, handler...)
}

func NewRouter(base string) *Router {
	return &Router{
		Base: base,
	}
}
