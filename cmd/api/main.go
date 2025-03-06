package main

import (
	"AgriBoost/internal/handlers"
	database "AgriBoost/internal/infra/database"
	"AgriBoost/internal/infra/env"
	"AgriBoost/internal/infra/jwt"
	"AgriBoost/internal/infra/middleware"
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
	jwt := jwt.NewJwt(*env)
	middleware := middleware.NewMiddleware(jwt)

	userRepository := repositories.NewUserRepo(db)
	userService := services.NewUserService(userRepository, jwt)
	handlers.NewUserHandler(v1, val, userService)

	campaignRepository := repositories.NewCampaignRepo(db)
	campaignService := services.NewCampaignService(campaignRepository)
	handlers.NewCampaignHandler(v1, val, campaignService, middleware)

	donationRepository := repositories.NewDonationRepo(db)
	donationService := services.NewDonationService(donationRepository, campaignRepository)
	handlers.NewDonationHandler(v1, donationService, val, middleware)

	quizRepository := repositories.NewQuizRepo(db)
	quizService := services.NewQuizService(quizRepository, userRepository)
	handlers.NewQuizHandler(v1, quizService, val, middleware)

	communityRepository := repositories.NewCommunityRepo(db)
	communityService := services.NewCommunityService(communityRepository, userRepository)
	handlers.NewCommunityHandler(v1, val, communityService, middleware)

	app.Listen(":8081")
}
