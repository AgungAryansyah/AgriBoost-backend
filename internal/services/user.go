package services

import (
	"AgriBoost/internal/infra/jwt"
	"AgriBoost/internal/models/dto"
	entity "AgriBoost/internal/models/entities"
	"AgriBoost/internal/repositories"
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceItf interface {
	Register(dto.Register) error
	Login(dto.Login) (string, error)
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

func (s *UserService) Login(login dto.Login) (string, error) {
	var user entity.User

	err := s.userRepo.Get(&user, dto.UserParam{Email: login.Email})
	if err != nil {
		return "", errors.New("email atau password salah")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
	if err != nil {
		return "", errors.New("email atau password salah")
	}

	return s.jwt.GenerateToken(user.Id)
}
