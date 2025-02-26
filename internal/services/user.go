package services

import (
	"AgriBoost/internal/models/dto"
	entity "AgriBoost/internal/models/entities"
	"AgriBoost/internal/repositories"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceItf interface {
	Register(dto.Register) error
}

type UserService struct {
	userRepo repositories.UserRepoItf
}

func NewUserService(userRepo repositories.UserRepoItf) UserServiceItf {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) Register(register dto.Register) error {
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	newUser := entity.User{
		Id:       uuid.New(),
		Name:     register.Name,
		Email:    register.Email,
		Password: string(hasedPassword),
	}

	err = s.userRepo.Create(&newUser)

	return err
}
