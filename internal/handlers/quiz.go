package handlers

import (
	"AgriBoost/internal/infra/middleware"
	"AgriBoost/internal/models/dto"
	entity "AgriBoost/internal/models/entities"
	"AgriBoost/internal/services"
	"AgriBoost/internal/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type QuizHandler struct {
	quizService services.QuizServiceItf
	validator   *validator.Validate
	middleware  middleware.MiddlewareItf
}

func NewQuizHandler(routerGroup fiber.Router, quizService services.QuizServiceItf, validator *validator.Validate, middleware middleware.MiddlewareItf) {
	QuizHandler := QuizHandler{
		quizService: quizService,
		validator:   validator,
		middleware:  middleware,
	}
	routerGroup = routerGroup.Group("/quiz")

	routerGroup.Get("/quizzes", middleware.Authentication, QuizHandler.GetAllQuizzes)
	routerGroup.Post("/quizz", middleware.Authentication, QuizHandler.GetQuizWithQuestionAndOption)
	routerGroup.Post("/attempt", middleware.Authentication, QuizHandler.CreateAttempt)
}

func (q *QuizHandler) GetAllQuizzes(ctx *fiber.Ctx) error {
	var quiz []entity.Quiz
	if err := q.quizService.GetAllQuizzes(&quiz); err != nil {
		return utils.HttpError(ctx, "failed to get data from the database", err)
	}

	return utils.HttpSuccess(ctx, "success", quiz)
}

func (q *QuizHandler) GetQuizWithQuestionAndOption(ctx *fiber.Ctx) error {
	var param dto.QuizParam
	if err := ctx.BodyParser(&param); err != nil {
		return utils.HttpError(ctx, "can't parse data, wrong JSON request format", err)
	}

	if err := q.validator.Struct(param); err != nil {
		return utils.HttpError(ctx, "invalid data", err)
	}

	var quiz dto.QuizDto
	if err := q.quizService.GetQuizWithQuestionAndOption(&quiz, param); err != nil {
		return utils.HttpError(ctx, "failed to get data from the database", err)
	}

	return utils.HttpSuccess(ctx, "success", quiz)
}

func (q *QuizHandler) CreateAttempt(ctx *fiber.Ctx) error {
	var answers dto.UserAnswersDto
	if err := ctx.BodyParser(&answers); err != nil {
		return utils.HttpError(ctx, "can't parse data, wrong JSON request format", err)
	}

	if err := q.validator.Struct(answers); err != nil {
		return utils.HttpError(ctx, "invalid data", err)
	}

	if err := q.quizService.CreateAttempt(answers); err != nil {
		return utils.HttpError(ctx, "can't process answers", err)
	}

	return utils.HttpSuccess(ctx, "answers submitted successfully", nil)
}
