package pkg

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ValidateJsonPayload(validate *validator.Validate, payload any) map[string]string {
	if err := validate.Struct(payload); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return TranslateErrors(validationErrors)
	}
	return nil
}

func ValidateSlicePayload(v *validator.Validate, payload *[]any) []map[string]any {
	validateErrors := []map[string]any{}

	for index, obj := range *payload {
		if err := v.Struct(obj); err != nil {
			fieldErrors := err.(validator.ValidationErrors)
			errorMessages := TranslateErrors(fieldErrors)
			keyMessage := fmt.Sprintf("error at index %v", index)
			validateError := map[string]any{keyMessage: errorMessages}
			validateErrors = append(validateErrors, validateError)
		}
	}

	if len(validateErrors) > 0 {
		return validateErrors
	}

	return nil
}
