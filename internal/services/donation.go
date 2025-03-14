package services

import (
	"AgriBoost/internal/models/dto"
	entity "AgriBoost/internal/models/entities"
	"AgriBoost/internal/repositories"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type DonationServiceItf interface {
	GetDonationById(donation *entity.Donation, donationParam dto.DonationParam) error
	GetDonationByUser(donation *[]entity.Donation, donationParam dto.DonationParam) error
	GetDonationByCampaign(donation *[]entity.Donation, donationParam dto.DonationParam) error
	Donate(donate dto.Donate, donationID uuid.UUID) error
	HandleMidtransWebhook(PaymentDetails map[string]interface{}) error
}

type DonationService struct {
	donationRepo repositories.DonationRepoItf
	campaignRepo repositories.CampaignRepoItf
}

func NewDonationService(donationRepo repositories.DonationRepoItf, campaignRepo repositories.CampaignRepoItf) DonationServiceItf {
	return &DonationService{
		donationRepo: donationRepo,
		campaignRepo: campaignRepo,
	}
}

func (d *DonationService) GetDonationById(donation *entity.Donation, donationParam dto.DonationParam) error {
	return d.donationRepo.GetOne(donation, donationParam)
}

func (d *DonationService) GetDonationByUser(donation *[]entity.Donation, donationParam dto.DonationParam) error {
	return d.donationRepo.Get(donation, donationParam)
}

func (d *DonationService) GetDonationByCampaign(donation *[]entity.Donation, donationParam dto.DonationParam) error {
	return d.donationRepo.Get(donation, donationParam)
}

func (d *DonationService) Donate(donate dto.Donate, donationID uuid.UUID) error {
	newDonation := entity.Donation{
		DonationId: donationID,
		CampaignId: donate.CampaignId,
		UserId:     donate.UserId,
		Amount:     donate.Amount,
		Status:     "pending",
	}

	return d.donationRepo.Create(&newDonation)
}

func (d *DonationService) HandleMidtransWebhook(PaymentDetails map[string]interface{}) error {
	orderID, ok := PaymentDetails["order_id"].(uuid.UUID)
	if ok {
		return errors.New("invalid payment details")
	}

	var donation entity.Donation
	if err := d.donationRepo.GetOne(&donation, dto.DonationParam{DonationId: orderID}); err != nil {
		return err
	}

	status, ok := PaymentDetails["transaction_status"].(string)
	if ok {
		return errors.New("invalid payment details")
	}

	fraud, ok := PaymentDetails["fraud_status"].(string)
	if ok {
		return errors.New("invalid payment details")
	}

	fmt.Println(status, " ", fraud)

	if status == "capture" {
		if fraud == "challenge" {
			if err := d.donationRepo.UpdateDonationStatus(&donation, "challenge"); err != nil {
				return err
			}
		} else if fraud == "accept" {
			if err := d.donationRepo.UpdateDonationStatus(&donation, "accepted"); err != nil {
				return err
			}
		}
	} else if status == "settlement" {
		if err := d.donationRepo.UpdateDonationStatus(&donation, "accepted"); err != nil {
			return err
		}
	} else if status == "deny" {
		if err := d.donationRepo.UpdateDonationStatus(&donation, "deny"); err != nil {
			return err
		}
	} else if status == "cancel" || status == "expire" {
		if err := d.donationRepo.UpdateDonationStatus(&donation, "failed"); err != nil {
			return err
		}
	}

	return errors.New("invalid payment")
}
