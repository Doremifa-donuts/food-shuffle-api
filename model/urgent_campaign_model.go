package model

import "time"

type UrgentCampaign struct {
	CampaignUuid   string        `gorm:"type:char(36);primary_key;"`              // 割引キャンペーン番号
	RestaurantUuid string        `gorm:"type:char(36);foreignKey:RestaurantUuid"` // レストランUUID
	StartAt		   time.Time	 `gorm:"not null"`								  // 空き時間開始日時
	EndAt		   time.Time	 `gorm:"not null"`								  // 空き時間終了日時
	Description    string        `gorm:"type:text;not null"`                      // 状況などの説明
	DiscountOffer  string        `gorm:"type:varchar(255);not null"`              // 割引情報などの詳細
	CreatedAt      time.Time     `gorm:"not null"`								  // 作成日時
	Reservations   []Reservation `gorm:"foreignKey:CampaignUuid"`
}

// サンプルデータ
var sampleUrgentCampaigns = []UrgentCampaign{}
