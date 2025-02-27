package pkg

import "github.com/go-playground/validator"

func ValidateJsonPayload(validate *validator.Validate, payload any) map[string]string {
	if err := validate.Struct(payload); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return TranslateErrors(validationErrors)
	}
	return nil
}
