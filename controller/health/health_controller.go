package health

import (
	"github.com/gofiber/fiber/v2"
	enum_res "github.com/yeom-c/golang-fiber-api/config/enum/response"
	response "github.com/yeom-c/golang-fiber-api/util/reponse"
)

type HealthController struct{}

func InitHealthController() *HealthController {
	healthController := NewHealthController()

	return healthController
}

func NewHealthController() *HealthController {
	return &HealthController{}
}

func (c *HealthController) Check(ctx *fiber.Ctx) error {
	return response.NewResponse(enum_res.Success, "OK", nil).Send(ctx)
}
