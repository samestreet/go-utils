package validationutils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func TranslateValidationErrors(err error) []string {
	var errors []string

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range validationErrors {
			switch fieldErr.Tag() {
			case "required":
				errors = append(errors, fmt.Sprintf("%s is required", fieldErr.Field()))
			case "excludesall":
				errors = append(errors, fmt.Sprintf("%s contains invalid characters [%s]", fieldErr.Field(), fieldErr.Param()))
			case "gt":
				errors = append(errors, fmt.Sprintf("%s must be greater than %s", fieldErr.Field(), fieldErr.Param()))
			default:
				errors = append(errors, fmt.Sprintf("%s is invalid", fieldErr.Field()))
			}
		}
	}

	return errors
}
