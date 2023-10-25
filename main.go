package main

import (
	"fmt"
	"log"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/yeom-c/golang-fiber-api/config/env"
	"github.com/yeom-c/golang-fiber-api/route"
)

func init() {
	env.InitEnv("./config/env")
}

func main() {
	app := fiber.New(
		fiber.Config{
			JSONEncoder: json.Marshal,
			JSONDecoder: json.Unmarshal,
		},
	)

	route.SetRoute(app)

	log.Fatal(app.Listen(fmt.Sprintf(":%d", env.Env.ServerPort)))
}
