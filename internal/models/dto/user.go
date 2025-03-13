package dto

import "github.com/google/uuid"

type Register struct {
	Name     string `json:"name" validated:"required,min=5"`
	Email    string `json:"email" validated:"required,email"`
	Password string `json:"password" validated:"required,min=6"`
}

type Login struct {
	Email    string `json:"email" validated:"required,email"`
	Password string `json:"password" validated:"required,min=6"`
}

type UserParam struct {
	Id    uuid.UUID `json:"id" validated:"required,uuid"`
	Email string    `json:"email" validated:"required,email"`
}
