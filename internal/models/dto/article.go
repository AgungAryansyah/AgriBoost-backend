package dto

import "github.com/google/uuid"

type ArticleParam struct {
	ArticleId uuid.UUID `json:"article_id" validate:"required,uuid"`
	UserId    uuid.UUID `json:"user_id" validate:"required,uuid"`
}
