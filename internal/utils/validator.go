package utils

import (
	"mime/multipart"
	"net/http"
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

func ValidateImage(fl validator.FieldLevel) bool {
	fileHeader, ok := fl.Field().Interface().(multipart.FileHeader)
	if !ok {
		return false
	}

	file, err := fileHeader.Open()
	if err != nil {
		return false
	}
	defer file.Close()

	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return false
	}

	mimeType := http.DetectContentType(buffer)
	allowedTypes := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
		"image/gif":  true,
	}

	return allowedTypes[mimeType]
}

func ValidateImageSize(fl validator.FieldLevel) bool {
	fileHeader, ok := fl.Field().Interface().(multipart.FileHeader)
	if !ok {
		return false
	}

	const maxSize = 2 * 1024 * 1024
	return fileHeader.Size <= maxSize
}

func RegisterValidator(v *validator.Validate) {
	v.RegisterValidation("answers_map", ValidateAnswersMap)
	v.RegisterValidation("phone_val", ValidateIndonesianPhone)
	v.RegisterValidation("image_val", ValidateImage)
	v.RegisterValidation("image_size_val", ValidateImageSize)
}
