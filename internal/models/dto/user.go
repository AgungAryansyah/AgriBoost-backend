package dto

import "github.com/google/uuid"

type Register struct {
	Phone    string `json:"phone" validate:"required,phone_val"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type UserParam struct {
	Id    uuid.UUID `json:"id" validate:"required,uuid"`
	Email string    `json:"email" validate:"required,email"`
}
