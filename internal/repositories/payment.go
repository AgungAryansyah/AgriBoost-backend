package repositories

import (
	"AgriBoost/internal/models/dto"
	entity "AgriBoost/internal/models/entities"

	"gorm.io/gorm"
)

type PaymentRepoItf interface {
	GetOne(payment *entity.Payment, paymentParam dto.PaymentParam) error
	Get(payment *[]entity.Payment, paymentParam dto.PaymentParam) error
	GetAll(payment *[]entity.Payment) error
}

type PaymentRepo struct {
	db *gorm.DB
}

func NewPaymentRepo(db *gorm.DB) PaymentRepoItf {
	return &PaymentRepo{db}
}

func (p *PaymentRepo) GetOne(payment *entity.Payment, paymentParam dto.PaymentParam) error {
	return p.db.First(payment, paymentParam).Error
}

func (p *PaymentRepo) Get(payment *[]entity.Payment, paymentParam dto.PaymentParam) error {
	return p.db.First(payment, paymentParam).Error
}

func (p *PaymentRepo) GetAll(payment *[]entity.Payment) error {
	return p.db.First(payment).Error
}
