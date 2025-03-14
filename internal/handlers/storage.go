package handlers

import (
	"AgriBoost/internal/infra/middleware"
	"AgriBoost/internal/models/dto"
	"AgriBoost/internal/services"
	"AgriBoost/internal/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type StorageHandler struct {
	storageService services.StorageServiceItf
	validator      *validator.Validate
	middleware     middleware.MiddlewareItf
}

func NewStorageHandler(routerGroup fiber.Router, storageService services.StorageServiceItf, validator *validator.Validate, middleware middleware.MiddlewareItf) {
	StorageHandler := StorageHandler{
		storageService: storageService,
		validator:      validator,
		middleware:     middleware,
	}

	routerGroup = routerGroup.Group("/storage")
	routerGroup.Post("/upload", StorageHandler.UploadHandler)
}

func (s *StorageHandler) UploadHandler(ctx *fiber.Ctx) error {
	file, err := ctx.FormFile("file")
	if err != nil {
		return utils.HttpError(ctx, "can't parse data, wrong request format", err)
	}

	uploadData := dto.FileUpload{File: file}
	if err := s.validator.Struct(uploadData); err != nil {
		errorMessages := ""
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "required":
				errorMessages += "File is required. "
			case "image_val":
				errorMessages += "Invalid file type. Only JPEG, PNG, and GIF are allowed. "
			case "image_size_val":
				errorMessages += "File size exceeds 2MB limit."
			}
		}
		return utils.HttpError(ctx, errorMessages, err)
	}

	url, err := s.storageService.UploadPicture(uploadData.File)
	if err != nil {
		return utils.HttpError(ctx, "can't upload data", err)
	}

	return utils.HttpSuccess(ctx, "success", url)
}
