package repositories

import (
	"AgriBoost/internal/models/dto"
	entity "AgriBoost/internal/models/entities"

	"gorm.io/gorm"
)

type ArticleRepoItf interface {
	GetAll(article *[]entity.Article) error
	Get(article *entity.Article, articleParam *dto.ArticleParam) error
}

type ArticleRepo struct {
	db *gorm.DB
}

func NewArticleRepo(db *gorm.DB) ArticleRepoItf {
	return &ArticleRepo{db}
}

func (a *ArticleRepo) GetAll(article *[]entity.Article) error {
	return a.db.Find(article).Error
}

func (a *ArticleRepo) Get(article *entity.Article, articleParam *dto.ArticleParam) error {
	return a.db.First(article, articleParam).Error
}
