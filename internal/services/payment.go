package services

import (
	"AgriBoost/internal/models/dto"
	entity "AgriBoost/internal/models/entities"
	"AgriBoost/internal/repositories"
)

type PaymentServiceItf interface {
	GetPaymentById(payment *entity.Payment, paymentParam dto.PaymentParam) error
	GetPaymentByUser(payment *[]entity.Payment, paymentParam dto.PaymentParam) error
	GetAllPayment(payment *[]entity.Payment) error
}

type PaymentService struct {
	paymentRepo repositories.PaymentRepoItf
}

func NewPaymentService(paymentRepo repositories.PaymentRepoItf) PaymentServiceItf {
	return &PaymentService{
		paymentRepo: paymentRepo,
	}
}

func (p *PaymentService) GetPaymentById(payment *entity.Payment, paymentParam dto.PaymentParam) error {
	return p.paymentRepo.GetOne(payment, paymentParam)
}

func (p *PaymentService) GetPaymentByUser(payment *[]entity.Payment, paymentParam dto.PaymentParam) error {
	return p.paymentRepo.Get(payment, paymentParam)
}

func (p *PaymentService) GetAllPayment(payment *[]entity.Payment) error {
	return p.paymentRepo.GetAll(payment)
}
