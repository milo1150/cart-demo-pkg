package pkg

import (
	"github.com/go-playground/validator"
	"github.com/samber/lo"
)

func TranslateErrors(errors validator.ValidationErrors) map[string]string {
	var errorMessage = make(map[string]string)
	for _, err := range errors {
		key := lo.SnakeCase(err.Field())
		errorMessage[key] = err.Tag()
	}
	return errorMessage
}

func GetSimpleErrorMessage(msg string) map[string]string {
	return map[string]string{"error": msg}
}
