package model

import "time"

// 個人が保存しているレビューを管理するテーブル
type ReviewArchive struct {
	ReviewUuid string    `gorm:"type:char(36);primary_key"` // レビューのUUID
	UserUuid   string    `gorm:"type:char(36);foreignKey:UserUuid;primary_key"`   // ユーザーのUUID
	CreatedAt  time.Time `gorm:"not null"`
	Review     []Review  `gorm:"foreignKey:ReviewUuid;references:ReviewUuid"`                                        // 登録日によってソートする
}
