package model

import "time"

// レビューテーブル
type Review struct {
	ReviewUuid string    `gorm:"type:char(36);primary_key"`          // レビューのUUID
	UserUuid   string    `gorm:"type:char(36);foreignkey:UserUuid"`  // レビューを投稿したユーザーのUUID
	RestoUuid  string    `gorm:"type:char(36);foreignkey:RestoUuid"` // レビューを投稿したレストランのUUID
	Images     []string  `gorm:"type:json;not null"`                 // レビューに関連する画像のパスをJSONで保存する
	CreatedAt  time.Time `gorm:"not null"`                           // レビューを投稿した日時
	Comment    string    `gorm:"type:text;not null"`                 // レビューのコメント
}
