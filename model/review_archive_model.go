package model

import "time"

// 個人が保存しているレビューを管理するテーブル
type ReviewArchive struct {
	ReviewUuid string    `gorm:"type:char(36);primary_key;foreignkey:ReviewUuid"` // レビューのUUID
	UserUuid   string    `gorm:"type:char(36);primary_key;foreignkey:UserUuid"`   // ユーザーのUUID
	CreatedAt  time.Time `gorm:"not null"`                                        // 登録日によってソートする
}
