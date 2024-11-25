package repository

import (
	"food-shuffle-api/model"

	"gorm.io/gorm"
)

func CreateUrgentCampaign(db *gorm.DB, urgentCampaign model.UrgentCampaign) error {
	return db.Create(&urgentCampaign).Error
}