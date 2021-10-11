package utils

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type ResponseData struct {
	Data  interface{} `json:"data"`
	Error interface{} `json:"error"`
}

type BaseResponseModel struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message interface{} `json:"message"`
	Code    int         `json:"code"`
	Meta    interface{} `json:"meta,omitempty"`
}

func ResponseSuccess(ctx *fiber.Ctx, data interface{}, message string, code int) error {
	success := false

	if code < http.StatusBadRequest {
		success = true
	}

	result := BaseResponseModel{
		Success: success,
		Data:    data,
		Message: message,
		Code:    code,
	}

	return ctx.Status(code).JSON(result)
}

func ResponseError(ctx *fiber.Ctx, err error, code int) error {

	result := BaseResponseModel{
		Success: false,
		Message: err.Error(),
		Code:    code,
	}

	return ctx.Status(code).JSON(result)
}
