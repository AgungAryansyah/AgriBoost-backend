package dto

import "github.com/google/uuid"

type Register struct {
	Name     string `json:"name" validate:"required,min=5"`
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
