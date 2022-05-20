package controllers

import (
	"fmt"
	"net/http"

	"github.com/eriicafes/gohttpserver/server"
	"github.com/eriicafes/gohttpserver/services"
)

type counter struct {
	*services.Container
	*server.Router
}

func Counter(container *services.Container) *server.Router {
	c := &counter{
		Container: container,
		Router:    server.NewRouter(""),
	}

	c.Get("/", c.counter, c.Ping)

	return c.Router
}

func (c *counter) Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{\"ping\": \"pong\", \"hits\": %v}", c.CountService.Hits())
}

func (c *counter) counter(w http.ResponseWriter, r *http.Request) {
	c.CountService.Increment()
}
