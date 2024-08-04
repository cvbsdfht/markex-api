package errs

import "github.com/gofiber/fiber/v2"

type AppError struct {
	Status  *fiber.Error
	Code    int
	Message string
	Err     error
}

func (e AppError) Error() string {
	return e.Message
}

func ErrNotFound(err error) error {
	return AppError{
		Status:  fiber.ErrNotFound,
		Code:    fiber.StatusNotFound,
		Message: "content not found",
		Err:     err,
	}
}

func ErrBadRequest(err error) error {
	return AppError{
		Status:  fiber.ErrBadRequest,
		Code:    fiber.StatusBadRequest,
		Message: "bad request",
		Err:     err,
	}
}

func ErrNotAcceptable(err error) error {
	return AppError{
		Status:  fiber.ErrNotAcceptable,
		Code:    fiber.StatusNotAcceptable,
		Message: "not acceptable",
		Err:     err,
	}
}

func ErrValidationFailed(err error) error {
	return AppError{
		Status:  fiber.ErrBadRequest,
		Code:    CodeValidationFailed,
		Message: "validation failed",
		Err:     err,
	}
}

const (
	CodeValidationFailed = 1401
)
