package validation

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// ValidationRegistry registers custom validation functions.
func ValidationRegistry(v *validator.Validate) {
	v.RegisterValidation("name", validateName)
}

// ValidateNameame checks if the value contains only alphabets and spaces.
func validateName(fl validator.FieldLevel) bool {
	if fl.Field().String() == "" {
		return true
	}

	return regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString(fl.Field().String())
}
