package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/yeom-c/golang-fiber-api/route"
)

func main() {
	app := fiber.New()

	route.SetRoute(app)

	log.Fatal(app.Listen(":5000"))
}
