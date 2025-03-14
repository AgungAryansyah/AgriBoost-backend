package services

import (
	"AgriBoost/internal/infra/jwt"
	"AgriBoost/internal/models/dto"
	entity "AgriBoost/internal/models/entities"
	"AgriBoost/internal/repositories"
	"AgriBoost/internal/utils"
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceItf interface {
	Register(register dto.Register, id *uuid.UUID) error
	Login(login dto.Login, id *uuid.UUID) (string, error)
	IsUserExistName(userName string, userId uuid.UUID) error
	EditProfile(edit *dto.EditProfile) error
}

type UserService struct {
	userRepo repositories.UserRepoItf
	jwt      jwt.JWTItf
}

func NewUserService(userRepo repositories.UserRepoItf, jwt jwt.JWTItf) UserServiceItf {
	return &UserService{
		userRepo: userRepo,
		jwt:      jwt,
	}
}

func (u *UserService) Register(register dto.Register, id *uuid.UUID) error {
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	newUser := entity.User{
		Id:       uuid.New(),
		Name:     utils.GetUsername(register.Email),
		Phone:    register.Phone,
		Email:    register.Email,
		Password: string(hasedPassword),
	}

	err = u.userRepo.Create(&newUser)

	*id = newUser.Id
	return err
}

func (u *UserService) Login(login dto.Login, id *uuid.UUID) (string, error) {
	var user entity.User

	err := u.userRepo.Get(&user, dto.UserParam{Email: login.Email})
	if err != nil {
		return "", errors.New("email atau password salah")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
	if err != nil {
		return "", errors.New("email atau password salah")
	}

	*id = user.Id

	return u.jwt.GenerateToken(user.Id, user.IsAdmin)
}

func (u *UserService) IsUserExistName(userName string, userId uuid.UUID) error {
	return u.userRepo.IsUserExistName(&userName, userId)
}

func (u *UserService) EditProfile(edit *dto.EditProfile) error {
	return u.userRepo.EditProfile(edit)
}
