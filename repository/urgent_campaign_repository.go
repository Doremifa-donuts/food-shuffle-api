package repository

import (
	"food-shuffle-api/model"

	"gorm.io/gorm"
)

func GetDescriptionByCampaignUuid(db *gorm.DB, uuid string) (string, error) {
	var campaign model.UrgentCampaign
	err := db.Where("campaign_uuid = ?", uuid).First(&campaign).Error
	if err != nil {
		return "", err
	}
	return campaign.Description, nil
}

func GetDiscountOfferByCampaignUuid(db *gorm.DB, uuid string) (string, error) {
	var campaign model.UrgentCampaign
	err := db.Where("campaign_uuid = ?", uuid).First(&campaign).Error
	if err != nil {
		return "", err
	}
	return campaign.DiscountOffer, nil
}
