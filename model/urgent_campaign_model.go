package model

import "time"

type UrgentCampaign struct {
	RestoUuid        string    `gorm:"type:char(36);foreignkey:RestoUuid;primary_key"` // レストランUUID
	UrgentCampaignNo int       `gorm:"type:integer;auto_increment;primary_key;"`       // 割引キャンペーン番号
	Description      string    `gorm:"type:text;not null"`                             // 状況などの説明
	DiscountOffer    string    `gorm:"type:varchar(255);not null"`                     // 割引情報などの詳細
	CreatedAt        time.Time `gorm:"not null"`                                       // 作成日時
}
