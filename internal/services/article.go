package services

import (
	"AgriBoost/internal/models/dto"
	entity "AgriBoost/internal/models/entities"
	"AgriBoost/internal/repositories"
)

type ArticleServiceItf interface {
	GetAll(article *[]entity.Article) error
	Get(article *entity.Article, articleParam *dto.ArticleParam) error
}

type ArticleService struct {
	articleRepo repositories.ArticleRepoItf
}

func NewArticleService(articleRepo repositories.ArticleRepoItf) ArticleServiceItf {
	return &ArticleService{
		articleRepo: articleRepo,
	}
}

func (a *ArticleService) GetAll(article *[]entity.Article) error {
	return a.articleRepo.GetAll(article)
}

func (a *ArticleService) Get(article *entity.Article, articleParam *dto.ArticleParam) error {
	return a.articleRepo.Get(article, articleParam)
}
