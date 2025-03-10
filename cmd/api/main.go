package main

import (
	"AgriBoost/internal/handlers"
	"AgriBoost/internal/infra/env"
	"AgriBoost/internal/infra/jwt"
	"AgriBoost/internal/infra/middleware"
	database "AgriBoost/internal/infra/postgres"
	"AgriBoost/internal/repositories"
	"AgriBoost/internal/services"
	"errors"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	env := env.NewEnv()
	if env == nil {
		panic(errors.New("tidak ada env"))
	}

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

	articleRepository := repositories.NewArticleRepo(db)
	articleService := services.NewArticleService(articleRepository)
	handlers.NewArticleHandler(v1, articleService, val, middleware)

	port := os.Getenv("APP_PORT")
	add := os.Getenv("APP_ADDRESS")
	app.Listen(add + ":" + port)
}
