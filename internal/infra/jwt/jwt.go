package jwt

import (
	"AgriBoost/internal/infra/env"
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

type JWTItf interface {
	GenerateToken(id uuid.UUID, isAdmin bool) (string, error)
	ValidateToken(token string) (uuid.UUID, bool, error)
}

type JWT struct {
	secretKey string
	expiresAt time.Time
}

func NewJwt(env env.Env) JWTItf {
	err := godotenv.Load()
	if err != nil {
		return nil
	}
	exp, err := strconv.Atoi(env.JWT_EXPIRED)
	if err != nil {
		return nil
	}
	secret := env.JWT_SECRET
	return &JWT{
		secretKey: secret,
		expiresAt: time.Now().Add(time.Duration(exp) * time.Hour),
	}
}

type Claims struct {
	Id      uuid.UUID
	IsAdmin bool
	jwt.RegisteredClaims
}

func (j *JWT) GenerateToken(id uuid.UUID, isAdmin bool) (string, error) {
	claim := Claims{
		Id:      id,
		IsAdmin: isAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(j.expiresAt),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (j *JWT) ValidateToken(tokenString string) (uuid.UUID, bool, error) {
	var claim Claims

	token, err := jwt.ParseWithClaims(tokenString, &claim, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})

	if err != nil {
		return uuid.Nil, false, err
	}

	if !token.Valid {
		return uuid.Nil, false, errors.New("token invalid")
	}

	return claim.Id, claim.IsAdmin, nil
}
