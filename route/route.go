package route

import (
	"github.com/gofiber/fiber/v2"
	cont_health "github.com/yeom-c/golang-fiber-api/controller/health"
)

func SetRoute(fiber *fiber.App) {
	healthController := cont_health.InitHealthController()
	fiber.Get("/health", healthController.Check)
}
