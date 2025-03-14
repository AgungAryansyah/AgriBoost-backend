package services

import (
	storage "AgriBoost/internal/infra/supabase"
	"mime/multipart"
)

type StorageServiceItf interface {
	UploadPicture(file *multipart.FileHeader) (string, error)
}

type StorageService struct {
	storage storage.StorageItf
}

func NewStorageService(storage storage.StorageItf) StorageServiceItf {
	return &StorageService{
		storage: storage,
	}
}

func (s *StorageService) UploadPicture(file *multipart.FileHeader) (string, error) {
	return s.storage.UploadFile(file)
}
