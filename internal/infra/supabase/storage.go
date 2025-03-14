package storage

import (
	"AgriBoost/internal/infra/env"
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
)

type storage struct {
	url       string
	publicUrl string
	token     string
	client    http.Client
}

type StorageItf interface {
	UploadFile(file *multipart.FileHeader) (string, error)
}

func New(env *env.Env) StorageItf {
	url := fmt.Sprintf("%s/storage/v1/object/%s/", env.SUPABASE_PROJECT_URL, env.SUPABASE_BUCKET_NAME)
	publicURL := fmt.Sprintf("%s/storage/v1/object/public/%s/", env.SUPABASE_PROJECT_URL, env.SUPABASE_BUCKET_NAME)
	return storage{
		url:       url,
		publicUrl: publicURL,
		token:     env.SUPABASE_TOKEN,
		client:    http.Client{},
	}
}

func (s storage) UploadFile(file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	fileBytes, err := io.ReadAll(src)
	if err != nil {
		return "", err
	}

	url := s.url + file.Filename

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(fileBytes))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+s.token)
	req.Header.Set("Content-Type", file.Header.Get("Content-Type"))

	resp, err := s.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return "", err
	}

	publicURL := s.publicUrl + file.Filename
	return publicURL, nil
}
