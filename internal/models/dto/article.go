package dto

import "github.com/google/uuid"

type ArticleParam struct {
	ArticleId uuid.UUID `json:"article_id" validated:"required,uuid"`
	UserId    uuid.UUID `json:"user_id" validated:"required,uuid"`
}
