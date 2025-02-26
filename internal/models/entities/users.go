package entity

import (
	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Name     string    `json:"name" gorm:"type:varchar(255)"`
	Email    string    `json:"email" gorm:"type:varchar(255);unique"`
	Password string    `json:"password" gorm:"type:varchar(255)"`
	Role     int       `json:"role" gorm:"type:smallint"`
}
