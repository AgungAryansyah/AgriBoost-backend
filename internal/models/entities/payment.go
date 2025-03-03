package entity

import (
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	PaymentId     uuid.UUID `json:"payment_id" gorm:"type:uuid;primaryKey"`
	DonationId    uuid.UUID `json:"donation_id" gorm:"type:uuid;not null"`
	TransactionId string    `json:"transaction_id" gorm:"type:varchar(50);unique;not null"`
	PaymentType   string    `json:"payment_type" gorm:"type:varchar(50)"`
	Status        string    `json:"status" gorm:"type:varchar(20);default:'pending'"`
	RawResponse   string    `json:"raw_response" gorm:"type:jsonb"`
	CreatedAt     time.Time `json:"created_at" gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"type:timestamp;default:current_timestamp"`
	Donation      Donation  `gorm:"foreignKey:DonationId"`
}
