package dto

import "github.com/google/uuid"

type Register struct {
	Name     string `json:"name" validated:"required"`
	Email    string `json:"email" validated:"required,email"`
	Password string `json:"password" validated:"required,min=6"`
}

type Login struct {
	Email    string `json:"email" validated:"required,email"`
	Password string `json:"password" validated:"required,min=6"`
}

type UserParam struct {
	Id    uuid.UUID
	Email string
}
