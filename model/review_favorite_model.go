package model

import "time"

// 個人がいいねしているレビューを管理するテーブル
type ReviewFavorite struct {
	ReviewUuid string    `gorm:"type:char(36);foreignKey:ReviewUuid;primary_key"` // レビューUUID
	UserUuid   string    `gorm:"type:char(36);foreignKey:UserUuid;primary_key"`   // ユーザーUUID
	CreatedAt  time.Time `gorm:"not null"`                                        // 作成日時によってソートされる
}
