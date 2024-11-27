package repository

import (
	"food-shuffle-api/model"

	"gorm.io/gorm"
)

func CreateUrgentCampaign(db *gorm.DB, urgentCampaign model.UrgentCampaign) error {
	return db.Create(&urgentCampaign).Error
}

func GetUrgentCampaign(db *gorm.DB, campaignUuid string) ([]model.UrgentCampaign, error) {
	var urgentCampaign []model.UrgentCampaign
	err := db.Where("campaign_uuid = ?", campaignUuid).Find(&urgentCampaign).Error
	return urgentCampaign, err
}