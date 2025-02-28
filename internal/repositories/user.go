package repositories

import (
	"AgriBoost/internal/models/dto"
	entity "AgriBoost/internal/models/entities"

	"gorm.io/gorm"
)

type UserRepoItf interface {
	Create(user *entity.User) error
	Get(user *entity.User, userParam dto.UserParam) error
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

func (r *UserRepo) Get(user *entity.User, userParam dto.UserParam) error {
	return r.db.First(&user, userParam).Error
}
