package main

import (
	"shortlink/src/controllers"
	"shortlink/src/repositories"
	"shortlink/src/services"

	"github.com/gofiber/fiber/v2"
	"github.com/sarulabs/di"
)

type Ioc struct{}

func (i *Ioc) New(app *fiber.App) {
	oc := i.NewIOC()

	r := NewRouter(oc)
	r.Routes(app)

}

func (i *Ioc) NewIOC() di.Container {

	builder, _ := di.NewBuilder()

	builder.Add(di.Def{
		Name: "controller",
		Build: func(ctn di.Container) (interface{}, error) {
			return controllers.NewController(ctn), nil
		},
	})
	builder.Add(di.Def{
		Name: "service",
		Build: func(ctn di.Container) (interface{}, error) {
			return services.NewService(ctn), nil
		},
	})

	builder.Add(di.Def{
		Name: "repository",
		Build: func(ctn di.Container) (interface{}, error) {
			return repositories.NewRepository(ctn), nil
		},
	})

	return builder.Build()
}
