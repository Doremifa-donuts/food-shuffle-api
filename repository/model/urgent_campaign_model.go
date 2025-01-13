package model

import "time"

type UrgentCampaign struct {
	CampaignUuid   string        `gorm:"type:char(36);primary_key;"`              // 割引キャンペーン番号
	RestaurantUuid string        `gorm:"type:char(36);foreignKey:RestaurantUuid"` // レストランUUID
	StartAt        time.Time     `gorm:"not null"`                                // 空き時間開始日時
	EndAt          time.Time     `gorm:"not null"`                                // 空き時間終了日時
	Description    string        `gorm:"type:text;not null"`                      // 状況などの説明
	DiscountOffer  string        `gorm:"type:varchar(255);not null"`              // 割引情報などの詳細
	CreatedAt      time.Time     `gorm:"not null"`                                // 作成日時
	Reservations   []Reservation `gorm:"foreignKey:CampaignUuid"`
}

// サンプルデータ
var sampleUrgentCampaigns = []UrgentCampaign{
	{
		CampaignUuid:   "0193a8ee-6972-7a4e-bc20-71de6517b565",
		RestaurantUuid: "97961bc1-70c9-43ea-9b4e-18f8bb6574f8",
		StartAt:        time.Now().Add(3 * time.Hour),
		EndAt:          time.Now().Add(5 * time.Hour),
		Description:    "10名の予約が突然キャンセルとなり、大変困っています。",
		DiscountOffer:  "お食事合計金額から500円割引きいたします",
	},
}
