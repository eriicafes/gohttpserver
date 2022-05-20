package main

import (
	"github.com/eriicafes/gohttpserver/controllers"
	"github.com/eriicafes/gohttpserver/server"
	"github.com/eriicafes/gohttpserver/services"
)

func main() {
	c := services.CreateContainer(
		services.NewCountService,
	)

	s := server.CreateServer(c)

	s.RegisterController(
		controllers.Counter,
	)

	s.ListenAndServe(":8080")
}
