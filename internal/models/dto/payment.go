package dto

import "github.com/google/uuid"

type PaymentParam struct {
	Id     uuid.UUID `json:"id" validated:"uuid"`
	UserId uuid.UUID `json:"user_id" validated:"uuid"`
}
