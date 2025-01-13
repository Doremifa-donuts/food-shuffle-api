package orm

import (
	"food-shuffle-api/repository/model"
	"time"

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

// 開始時間付近のお助けブースト情報を取得する
func ListUrgentCampaignByStartAt(db *gorm.DB, startAt time.Time) ([]model.UrgentCampaign, error) {
	var urgentCampaigns []model.UrgentCampaign
	// TODO: 通知対象の時間を厳密にする
	err := db.Where("start_at between ? and ?", startAt.Add(-5*time.Minute), startAt).Find(&urgentCampaigns).Error
	return urgentCampaigns, err
}
