package pkg

import (
	"github.com/go-playground/validator/v10"
)

func TranslateErrors(errors validator.ValidationErrors) map[string]string {
	var errorMessage = make(map[string]string)
	for _, err := range errors {
		errorMessage[err.StructNamespace()] = err.ActualTag()
	}
	return errorMessage
}

func GetSimpleErrorMessage(msg string) map[string]string {
	return map[string]string{"error": msg}
}
