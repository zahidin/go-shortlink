package main

import (
	"shortlink/src/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/sarulabs/di"
)

type Router struct {
	controller *controllers.Controllers
}

func NewRouter(ioc di.Container) *Router {
	return &Router{
		controller: ioc.Get("controller").(*controllers.Controllers),
	}
}

func (r *Router) Routes(f *fiber.App) {
	v1 := f.Group("/api/v1")

	v1.Get("/shortlink/:id", r.controller.Shortlink.GetLink)
	v1.Post("/shortlink", r.controller.Shortlink.AddShortLink)
}
