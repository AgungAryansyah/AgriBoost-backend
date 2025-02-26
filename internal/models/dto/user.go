package dto

type Register struct {
	Name     string `json:"name" validated:"required"`
	Email    string `json:"email" validated:"required,email"`
	Password string `json:"password" validated:"required,min=6"`
}
