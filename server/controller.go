package server

import "github.com/eriicafes/gohttpserver/services"

type Controller func(c *services.Container) *Router
