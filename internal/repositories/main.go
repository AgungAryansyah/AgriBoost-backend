package repositories

import (
	entity "AgriBoost/internal/models/entities"

	"gorm.io/gorm"
)

type UserRepoItf interface {
	Create(user *entity.User) error
}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepoItf {
	return &UserRepo{db}
}

func (r *UserRepo) Create(user *entity.User) error {
	return r.db.Create(user).Error
}
