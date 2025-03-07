package database

import (
	"AgriBoost/internal/infra/env"
	entity "AgriBoost/internal/models/entities"
	"errors"
	"fmt"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(env env.Env) (*gorm.DB, error) {
	err := godotenv.Load()

	if err != nil {
		return nil, errors.New(".env file deos not exist")
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		env.DB_HOST,
		env.DB_PORT,
		env.DB_USER,
		env.DB_PASSWORD,
		env.DB_NAME, env.DB_SSLMODE,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.New("can not connect to databse")
	}

	migrate := []interface{}{
		&entity.User{},
		&entity.Quiz{},
		&entity.Question{},
		&entity.QuestionOption{},
		&entity.QuizAttempt{},
		&entity.Donation{},
		&entity.Community{},
		&entity.CommunityMember{},
		&entity.Campaign{},
		&entity.Article{},
	}
	db.AutoMigrate(migrate...)

	return db, nil
}
