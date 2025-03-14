package utils

import (
	"regexp"

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

func ValidateIndonesianPhone(fl validator.FieldLevel) bool {
	phone := fl.Field().String()
	pattern := `^(\+62|62|0)8\d{8,13}$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(phone)
}

func RegisterValidator(v *validator.Validate) {
	v.RegisterValidation("answers_map", ValidateAnswersMap)
	v.RegisterValidation("phone_val", ValidateIndonesianPhone)
}
