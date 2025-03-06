package repositories

import (
	"AgriBoost/internal/models/dto"
	entity "AgriBoost/internal/models/entities"
	"time"

	"gorm.io/gorm"
)

type DonationRepoItf interface {
	GetOne(donation *entity.Donation, donationParam dto.DonationParam) error
	Get(donation *[]entity.Donation, donationParam dto.DonationParam) error
	Create(donation *entity.Donation) error
	UpdateDonationStatus(donation *entity.Donation, status string) error
}

type DonationRepo struct {
	db *gorm.DB
}

func NewDonationRepo(db *gorm.DB) DonationRepoItf {
	return &DonationRepo{db}
}

func (d *DonationRepo) Get(donation *[]entity.Donation, donationParam dto.DonationParam) error {
	return d.db.Find(donation, donationParam).Error
}

func (d *DonationRepo) GetOne(donation *entity.Donation, donationParam dto.DonationParam) error {
	return d.db.First(donation, donationParam).Error
}

func (d *DonationRepo) Create(donation *entity.Donation) error {
	return d.db.Create(donation).Error
}

func (d *DonationRepo) UpdateDonationStatus(donation *entity.Donation, status string) error {
	donation.Status = status
	donation.DonationDate = time.Now()

	return d.db.Save(donation).Error
}
