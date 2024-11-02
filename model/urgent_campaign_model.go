package model

import "time"

type UrgentCampaign struct {
	CampaignUuid   string        `gorm:"type:char(36);primary_key;"`              // 割引キャンペーン番号
	RestaurantUuid string        `gorm:"type:char(36);foreignKey:RestaurantUuid"` // レストランUUID
	Description    string        `gorm:"type:text;not null"`                      // 状況などの説明
	DiscountOffer  string        `gorm:"type:varchar(255);not null"`              // 割引情報などの詳細
	CreatedAt      time.Time     `gorm:"not null"`
	Reservations   []Reservation `gorm:"foreignKey:CampaignUuid"` // 作成日時
}
