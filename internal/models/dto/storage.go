package dto

import "mime/multipart"

type FileUpload struct {
	File *multipart.FileHeader `form:"file" validate:"required,image_val,image_size_val=2"`
}
