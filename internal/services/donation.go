package services

import (
	"AgriBoost/internal/models/dto"
	entity "AgriBoost/internal/models/entities"
	"AgriBoost/internal/repositories"
)

type DonationServiceItf interface {
	GetDonationById(donation *entity.Donation, donationParam dto.DonationParam) error
	GetDonationByUser(donation *[]entity.Donation, donationParam dto.DonationParam) error
	GetDonationByCampaign(donation *[]entity.Donation, donationParam dto.DonationParam) error
}

type DonationService struct {
	donationRepo repositories.DonationRepoItf
}

func NewDonationService(donationRepo repositories.DonationRepoItf) DonationServiceItf {
	return &DonationService{
		donationRepo: donationRepo,
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
