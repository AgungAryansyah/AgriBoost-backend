package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	DB_HOST              string
	DB_PORT              string
	DB_USER              string
	DB_PASSWORD          string
	DB_NAME              string
	DB_SSLMODE           string
	JWT_EXPIRED          string
	JWT_SECRET           string
	MIDTRANS_SERVER_KEY  string
	SUPABASE_PROJECT_URL string
	SUPABASE_TOKEN       string
	SUPABASE_BUCKET_NAME string
}

func NewEnv() *Env {
	err := godotenv.Load()

	if err != nil {
		fmt.Print("no env")
		return nil
	}

	return &Env{
		DB_HOST:              os.Getenv("DB_HOST"),
		DB_PORT:              os.Getenv("DB_PORT"),
		DB_USER:              os.Getenv("DB_USER"),
		DB_PASSWORD:          os.Getenv("DB_PASSWORD"),
		DB_NAME:              os.Getenv("DB_NAME"),
		DB_SSLMODE:           os.Getenv("DB_SSLMODE"),
		JWT_EXPIRED:          os.Getenv("JWT_EXPIRED"),
		JWT_SECRET:           os.Getenv("JWT_SECRET"),
		MIDTRANS_SERVER_KEY:  os.Getenv("MIDTRANS_SERVER_KEY"),
		SUPABASE_PROJECT_URL: os.Getenv("SUPABASE_PROJECT_URL"),
		SUPABASE_TOKEN:       os.Getenv("SUPABASE_TOKEN"),
		SUPABASE_BUCKET_NAME: os.Getenv("SUPABASE_BUCKET_NAME"),
	}
}
