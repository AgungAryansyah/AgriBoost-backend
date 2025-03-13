package storage

import (
	"AgriBoost/internal/infra/env"
	"mime/multipart"

	supabasestorageuploader "github.com/adityarizkyramadhan/supabase-storage-uploader"
)

type storage struct {
	client *supabasestorageuploader.Client
}

type StorageItf interface {
	UploadFile(file *multipart.FileHeader) (string, error)
	DeleteFile(link string) error
}

func New(env env.Env) StorageItf {
	supClient := supabasestorageuploader.New(
		env.SUPABASE_PROJECT_URL,
		env.SUPABASE_TOKEN,
		env.SUPABASE_BUCKET_NAME,
	)
	return storage{
		client: supClient,
	}
}

func (s storage) UploadFile(file *multipart.FileHeader) (string, error) {
	return s.client.Upload(file)
}

func (s storage) DeleteFile(link string) error {
	return s.client.Delete(link)
}
