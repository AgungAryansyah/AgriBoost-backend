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

type ArticleHandler struct {
	articleService services.ArticleServiceItf
	validator      *validator.Validate
	middleware     middleware.MiddlewareItf
}

func NewArticleHandler(routerGroup fiber.Router, articleService services.ArticleServiceItf, validator *validator.Validate, middleware middleware.MiddlewareItf) {
	ArticleHandler := ArticleHandler{
		articleService: articleService,
		validator:      validator,
		middleware:     middleware,
	}

	routerGroup = routerGroup.Group("/article")
	routerGroup.Get("/all", ArticleHandler.GetAllArticles)
	routerGroup.Get("/one", ArticleHandler.GetArticle)
}

func (a *ArticleHandler) GetAllArticles(ctx *fiber.Ctx) error {
	var articles []entity.Article
	if err := a.articleService.GetAll(&articles); err != nil {
		return utils.HttpError(ctx, "failed to get data from the database", err)
	}

	return utils.HttpSuccess(ctx, "success", articles)
}

func (a *ArticleHandler) GetArticle(ctx *fiber.Ctx) error {
	var param dto.ArticleParam
	if err := ctx.BodyParser(&param); err != nil {
		return utils.HttpError(ctx, "can't parse data, wrong JSON request format", err)
	}

	var article entity.Article
	if err := a.articleService.Get(&article, &param); err != nil {
		return utils.HttpError(ctx, "failed to get data from the database", err)
	}

	return utils.HttpSuccess(ctx, "success", article)
}
