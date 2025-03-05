package repositories

import (
	"AgriBoost/internal/models/dto"
	entity "AgriBoost/internal/models/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepoItf interface {
	Create(user *entity.User) error
	Get(user *entity.User, userParam dto.UserParam) error
	AddQuizPoint(userParam dto.UserParam, score int) error
	IsUserExist(exist *bool, userId uuid.UUID) error
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
	return r.db.First(user, userParam).Error
}

func (r *UserRepo) AddQuizPoint(userParam dto.UserParam, score int) error {
	var user entity.User
	if err := r.Get(&user, userParam); err != nil {
		return err
	}
	user.QuizPoint += score
	return r.db.Save(&user).Error
}

func (r *UserRepo) IsUserExist(exist *bool, userId uuid.UUID) error {
	return r.db.Model(&entity.User{}).Select("id").Where("id = ?", userId).First(&exist).Error
}
