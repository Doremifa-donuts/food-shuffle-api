package repository

import (
	"food-shuffle-api/model"

	"gorm.io/gorm"
)

// キャンペーンを新規追加する
func CreateUrgentCampaign(db *gorm.DB, urgentCampaign model.UrgentCampaign) error {
	return db.Create(&urgentCampaign).Error
}

// キャンペーンUUIDが一致するキャンペーンの情報を取得する
func GetUrgentCampaign(db *gorm.DB, campaignUuid string) (model.UrgentCampaign, error) {
	var urgentCampaign model.UrgentCampaign
	err := db.Where("campaign_uuid = ?", campaignUuid).First(&urgentCampaign).Error
	return urgentCampaign, err
}
