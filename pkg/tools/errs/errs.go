package errs

import "github.com/gofiber/fiber/v2"

type AppError struct {
	Code    int
	Message string
	Err     error
}

func (e AppError) Error() string {
	return e.Message
}

func ErrNoContent(err error) error {
	return AppError{
		Code:    fiber.StatusNoContent,
		Message: "no-content",
		Err:     err,
	}
}

func ErrUnexpected(err error) error {
	return AppError{
		Code:    fiber.StatusInternalServerError,
		Message: "unexpected",
		Err:     err,
	}
}

func ErrNotAcceptable(err error) error {
	return AppError{
		Code:    fiber.StatusNotAcceptable,
		Message: "not-acceptable",
		Err:     err,
	}
}

func ErrValidationFailed(err error) error {
	return AppError{
		Code:    fiber.StatusBadRequest,
		Message: "validation-failed",
		Err:     err,
	}
}