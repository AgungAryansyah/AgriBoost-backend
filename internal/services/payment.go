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

func (d *PaymentService) GetPaymentById(payment *entity.Payment, paymentParam dto.PaymentParam) error {
	return d.paymentRepo.GetOne(payment, paymentParam)
}

func (d *PaymentService) GetPaymentByUser(payment *[]entity.Payment, paymentParam dto.PaymentParam) error {
	return d.paymentRepo.Get(payment, paymentParam)
}

func (d *PaymentService) GetAllPayment(payment *[]entity.Payment) error {
	return d.paymentRepo.GetAll(payment)
}
