package main

import (
	"log"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/yeom-c/golang-fiber-api/route"
)

func main() {
	app := fiber.New(
		fiber.Config{
			JSONEncoder: sonic.Marshal,
			JSONDecoder: sonic.Unmarshal,
		},
	)

	route.SetRoute(app)

	log.Fatal(app.Listen(":5000"))
}
