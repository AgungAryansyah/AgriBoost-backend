package main

import (
	"AgriBoost/internal/handlers"
	"AgriBoost/internal/infra/env"
	"AgriBoost/internal/infra/jwt"
	"AgriBoost/internal/infra/middleware"
	"AgriBoost/internal/infra/midtrans"
	database "AgriBoost/internal/infra/postgres"
	"AgriBoost/internal/repositories"
	"AgriBoost/internal/services"
	"AgriBoost/internal/utils"
	"errors"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(middleware.RateLimiter())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "https://api.sandbox.midtrans.com",
		AllowMethods:     "GET, POST, DELETE",
		AllowHeaders:     "Content-Type, Authorization, X-Requested-With",
		AllowCredentials: true,
	}))

	env := env.NewEnv()
	if env == nil {
		panic(errors.New("tidak ada env"))
	}

	db, err := database.Connect(*env)
	if err != nil {
		panic(err)
	}
	val := validator.New()
	utils.RegisterValidator(val)

	v1 := app.Group("api/v1")
	jwt := jwt.NewJwt(*env)
	middleware := middleware.NewMiddleware(jwt)
	midtrans := midtrans.NewMidtrans(*env)

	userRepository := repositories.NewUserRepo(db)
	userService := services.NewUserService(userRepository, jwt)
	handlers.NewUserHandler(v1, val, userService)

	campaignRepository := repositories.NewCampaignRepo(db)
	campaignService := services.NewCampaignService(campaignRepository)
	handlers.NewCampaignHandler(v1, val, campaignService, middleware)

	donationRepository := repositories.NewDonationRepo(db)
	donationService := services.NewDonationService(donationRepository, campaignRepository)
	handlers.NewDonationHandler(v1, donationService, val, middleware, midtrans)

	quizRepository := repositories.NewQuizRepo(db)
	quizService := services.NewQuizService(quizRepository, userRepository)
	handlers.NewQuizHandler(v1, quizService, val, middleware)

	communityRepository := repositories.NewCommunityRepo(db)
	communityService := services.NewCommunityService(communityRepository, userRepository)
	handlers.NewCommunityHandler(v1, val, communityService, middleware)

	articleRepository := repositories.NewArticleRepo(db)
	articleService := services.NewArticleService(articleRepository)
	handlers.NewArticleHandler(v1, articleService, val, middleware)

	messageRepository := repositories.NewMessageRepo(db)
	messageService := services.NewMessageService(messageRepository)
	handlers.NewMessageHandler(v1, messageService, communityService, userService, val, middleware)

	port := os.Getenv("APP_PORT")
	add := os.Getenv("APP_ADDRESS")
	app.Listen(add + ":" + port)
}
