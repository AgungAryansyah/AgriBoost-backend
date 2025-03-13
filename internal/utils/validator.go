package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func ValidateAnswersMap(fl validator.FieldLevel) bool {
	answers, ok := fl.Field().Interface().(map[uuid.UUID]uuid.UUID)
	if !ok {
		return false
	}

	for key, value := range answers {
		if key == uuid.Nil || value == uuid.Nil {
			return false
		}
	}
	return true
}

func RegisterValidator(v *validator.Validate) {
	v.RegisterValidation("answers_map", ValidateAnswersMap)
}
