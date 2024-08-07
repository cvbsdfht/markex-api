package handler

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/markex-api/pkg/tools/errs"
	"github.com/markex-api/pkg/tools/utils"
)

type HandlerFn func() (interface{}, error)

func Handler(c *fiber.Ctx, handlerFn HandlerFn) error {
	data, err := handlerFn()
	if err != nil {
		return responseError(c, err)
	}
	return responseSuccess(c, data)
}

func HandlerWithBody(c *fiber.Ctx, request interface{}, handlerFn HandlerFn) error {
	if err := c.BodyParser(request); err != nil {
		return responseError(c, err)
	}
	data, err := handlerFn()
	if err != nil {
		return responseError(c, err)
	}
	return responseSuccess(c, data)
}

func responseError(c *fiber.Ctx, err error) error {
	appError, ok := err.(errs.AppError)

	if ok {
		appErrorResponse := &errs.ApiErrorResponse{
			Status:  appError.Status.Code,
			Code:    appError.Code,
			Message: appError.Message,
			Time:    time.Now(),
			Request: fmt.Sprintf("uri: %v | x-request-id: %v", c.OriginalURL(), utils.InterfaceToString(c.Locals(requestid.ConfigDefault.ContextKey))),
			Detail:  appError.Err.Error(),
		}

		return c.Status(appErrorResponse.Status).JSON(appErrorResponse)
	}

	return c.Status(fiber.StatusBadRequest).JSON(err)
}

func responseSuccess(c *fiber.Ctx, payload interface{}) error {
	return c.Status(fiber.StatusOK).JSON(payload)
}
