package errs

import "github.com/gofiber/fiber/v2"

type AppError struct {
	Code    int
	Message string
}

func (e AppError) Error() string {
	return e.Message
}

func ErrNoContent(message string) error {
	return AppError{
		Code:    fiber.StatusNoContent,
		Message: message,
	}
}

func ErrUnexpected(message string) error {
	return AppError{
		Code:    fiber.StatusInternalServerError,
		Message: message,
	}
}
