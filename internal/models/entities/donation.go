package entity

import (
	"time"

	"github.com/google/uuid"
)

type Donation struct {
	DonationId     uuid.UUID `json:"donation_id" gorm:"type:uuid;primaryKey"`
	CampaignId     uuid.UUID `json:"campaign_id" gorm:"type:uuid;not null"`
	UserId         uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	Amount         int       `json:"amount" gorm:"type:integer;check:amount>0"`
	Status         string    `json:"status" gorm:"type:varchar(20);default:'pending'"`
	PaymentDetails string    `json:"payment_details" gorm:"type:jsonb"`
	DonationDate   time.Time `json:"donation_date" gorm:"type:timestamp;default:current_timestamp"`
}
