package main

import (
	"AgriBoost/internal/handlers"
	database "AgriBoost/internal/infra/database"
	"AgriBoost/internal/infra/env"
	"AgriBoost/internal/infra/jwt"
	"AgriBoost/internal/repositories"
	"AgriBoost/internal/services"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	env := env.NewEnv()
	db, err := database.Connect(*env)
	if err != nil {
		panic(err)
	}
	val := validator.New()
	v1 := app.Group("api/v1")

	userRepository := repositories.NewUserRepo(db)
	userService := services.NewUserService(userRepository, jwt.NewJwt(*env))
	handlers.NewUserHandler(v1, val, userService)

	app.Listen(":8081")
}
