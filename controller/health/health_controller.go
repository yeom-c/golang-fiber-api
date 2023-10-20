package health

import "github.com/gofiber/fiber/v2"

type HealthController struct{}

func InitHealthController() *HealthController {
	healthController := NewHealthController()

	return healthController
}

func NewHealthController() *HealthController {
	return &HealthController{}
}

func (c *HealthController) Check(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "OK",
	})
}
