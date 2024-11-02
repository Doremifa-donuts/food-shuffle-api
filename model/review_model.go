package model

import "time"

// レビューテーブル
type Review struct {
	ReviewUuid              string                    `gorm:"type:char(36);primary_key;"`              // レビューのUUID
	UserUuid                string                    `gorm:"type:char(36);foreignKey:UserUuid"`       // レビューを投稿したユーザーのUUID
	RestaurantUuid          string                    `gorm:"type:char(36);foreignKey:RestaurantUuid"` // レビューを投稿したレストランのUUID
	Images                  StringArray               `gorm:"type:json;not null"`                      // レビューに関連する画像のパスをJSONで保存する
	CreatedAt               time.Time                 `gorm:"not null"`                                // レビューを投稿した日時
	Comment                 string                    `gorm:"type:text;not null"`                      // レビューのコメント
	ReviewArchives          []ReviewArchive           `gorm:"foreignKey:ReviewUuid"`
	ReviewFavorites         []ReviewFavorite          `gorm:"foreignKey:ReviewUuid"`
	ReviewReceives          []ReviewReceive           `gorm:"foreignKey:ReviewUuid"`
	PopupGroupSharedReviews []PopupGroupSharedReviews `gorm:"foreignKey:ReviewUuid"`
}
