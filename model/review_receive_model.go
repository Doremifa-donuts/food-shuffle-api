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