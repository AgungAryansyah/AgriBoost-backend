package jwt

import (
	"AgriBoost/internal/infra/env"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

type JWTItf interface {
	GenerateToken(id uuid.UUID) (string, error)
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
	Id uuid.UUID
	jwt.RegisteredClaims
}

func (j *JWT) GenerateToken(id uuid.UUID) (string, error) {
	claim := Claims{
		Id: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(j.expiresAt),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString(j.secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
