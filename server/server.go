package server

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/eriicafes/gohttpserver/services"
)

type Server struct {
	Container *services.Container
	routes    []Route
}

func (s *Server) RegisterRoute(routes ...Route) {
	s.routes = append(s.routes, routes...)
}

func (s *Server) RegisterController(controllers ...Controller) {
	for _, controller := range controllers {
		router := controller(s.Container)
		s.RegisterRoute(router.routes...)
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method+":", r.RequestURI)

	for _, route := range s.routes {
		if route.Path == r.RequestURI && route.Method == r.Method {
			route.Handler(w, r)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "{\"error\": \"not found\"}")
}

func (s *Server) ListenAndServe(port string) {
	log.Println("Server listening on port", strings.Replace(port, ":", "", 1))

	for _, route := range s.routes {
		log.Println(route.Method, route.Path)
	}

	log.Fatal(http.ListenAndServe(port, s))
}

func CreateServer(c *services.Container) *Server {
	return &Server{
		Container: c,
	}
}
