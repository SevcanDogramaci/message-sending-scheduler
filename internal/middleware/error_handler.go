package middleware

import (
	"errors"

	"github.com/SevcanDogramaci/message-sending-scheduler/internal/model"
	"github.com/gofiber/fiber/v2"
)

func mapToFiberError(err error) *fiber.Error {
	switch err {
	case model.ErrorInvalidMessageStatus:
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	case model.ErrorMessageNotFound:
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	default:
		return fiber.ErrInternalServerError
	}
}

func InitErrorHandler(c *fiber.Ctx, err error) error {
	err = mapToFiberError(err)
	code := fiber.StatusInternalServerError

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	return c.
		Status(code).
		JSON(fiber.Map{"error": err.Error()})
}
