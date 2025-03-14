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
	IsUserExist(user *entity.User, userId uuid.UUID) error
	IsUserExistName(userName *string, userId uuid.UUID) error
	AddDonationPoint(userParam dto.UserParam, score int) error
}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepoItf {
	return &UserRepo{db}
}

func (u *UserRepo) Create(user *entity.User) error {
	return u.db.Create(user).Error
}

func (u *UserRepo) Get(user *entity.User, userParam dto.UserParam) error {
	return u.db.First(user, userParam).Error
}

func (u *UserRepo) AddQuizPoint(userParam dto.UserParam, score int) error {
	var user entity.User
	if err := u.Get(&user, userParam); err != nil {
		return err
	}
	user.QuizPoint += score
	return u.db.Save(&user).Error
}

func (u *UserRepo) AddDonationPoint(userParam dto.UserParam, score int) error {
	var user entity.User
	if err := u.Get(&user, userParam); err != nil {
		return err
	}
	user.DonationPoint += score
	return u.db.Save(&user).Error
}

func (u *UserRepo) IsUserExist(user *entity.User, userId uuid.UUID) error {
	return u.db.Model(&entity.User{}).Select("id").Where("id = ?", userId).First(&user).Error
}

func (u *UserRepo) IsUserExistName(userName *string, userId uuid.UUID) error {
	var user entity.User
	return u.db.Model(&entity.User{}).Select("name").Where("id = ?", userId).First(&user).Error
}
