package model

import "time"

// 個人がいいねしているレビューを管理するテーブル
type ReviewFavorite struct {
	ReviewUuid string    `gorm:"type:char(36);primary_key"` // レビューUUID
	UserUuid   string    `gorm:"type:char(36);foreignKey:UserUuid;primary_key"`   // ユーザーUUID
	CreatedAt  time.Time `gorm:"not null"`
	Review     []Review  `gorm:"foreignKey:ReviewUuid;references:ReviewUuid"`                                        // 作成日時によってソートされる
}
