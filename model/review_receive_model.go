package model

import (
	"time"
)

// いいねしているレビューを管理するテーブル
type ReviewReceive struct {
	ReviewUuid string    `gorm:"type:char(36);foreignKey:ReviewUuid;primary_key"`
	UserUuid   string    `gorm:"type:char(36);foreignKey:UserUuid;primary_key"`
	CreatedAt  time.Time `gorm:"not null"`
}

// サンプルデータ
var sampleReviewReceives = []ReviewReceive{
	{
		ReviewUuid: "39f93b17-c378-46f4-b55e-0c65642d99b0",
		UserUuid:   "91a78381-f472-496b-90e3-2c66a33391d1",
	},
}
