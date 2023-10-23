package response

import (
	"github.com/gofiber/fiber/v2"
	enum_res "github.com/yeom-c/golang-fiber-api/config/enum/response"
)

type Response struct {
	Code    enum_res.Code `json:"code"`
	Message string        `json:"message,omitempty"`
	Data    interface{}   `json:"data,omitempty"`
}

func NewResponse(code enum_res.Code, message string, data interface{}) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func (r *Response) Send(ctx *fiber.Ctx) error {
	return ctx.Status(codeToStatus(r.Code)).JSON(r)
}

func codeToStatus(code enum_res.Code) int {
	status := fiber.StatusBadRequest
	switch code {
	case enum_res.Fail:
		status = fiber.StatusInternalServerError
	case enum_res.Success:
		status = fiber.StatusOK
	}

	return status
}
