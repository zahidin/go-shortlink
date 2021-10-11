package main

import (
	"log"
	"shortlink/src/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	ioc := Ioc{}
	app := fiber.New()
	ioc.New(app)

	log.Fatal(app.Listen(":" + config.GetConfig("PORT")))
}
