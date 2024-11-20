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
	UserReviewFlags         []UserReviewFlag          `gorm:"foreignKey:ReviewUuid"`
	PopupGroupSharedReviews []PopupGroupSharedReviews `gorm:"foreignKey:ReviewUuid"`
}

// サンプルデータ
var sampleReviews = []Review{
	{
		ReviewUuid:     "39f93b17-c378-46f4-b55e-0c65642d99b0",
		UserUuid:       "a0adb027-0f54-4c1a-9ed3-86041c863344",
		RestaurantUuid: "97961bc1-70c9-43ea-9b4e-18f8bb6574f8",
		Images:         StringArray{"92319be6-48e1-485a-9eaf-8c7b6af0f789", "27922620-d0b6-4fbe-b76f-35521b1cb851", "a5c7632f-a3e0-410a-9d73-75cc6646ad42"},
		Comment:        "This review was received by viewer. posted by poster",
	},
	{
		ReviewUuid:     "e08505ac-eb06-43ea-a29b-b206367f7b8d",
		UserUuid:       "a0adb027-0f54-4c1a-9ed3-86041c863344",
		RestaurantUuid: "97961bc1-70c9-43ea-9b4e-18f8bb6574f8",
		Images:         StringArray{"92319be6-48e1-485a-9eaf-8c7b6af0f789", "27922620-d0b6-4fbe-b76f-35521b1cb851", "a5c7632f-a3e0-410a-9d73-75cc6646ad42"},
		Comment:        "This review was archived by viewer. posted by poster",
	},
	{
		ReviewUuid:     "573fa1e4-1510-4eaf-9f1f-9df903bbd020",
		UserUuid:       "a0adb027-0f54-4c1a-9ed3-86041c863344",
		RestaurantUuid: "97961bc1-70c9-43ea-9b4e-18f8bb6574f8",
		Images:         StringArray{"92319be6-48e1-485a-9eaf-8c7b6af0f789", "27922620-d0b6-4fbe-b76f-35521b1cb851", "a5c7632f-a3e0-410a-9d73-75cc6646ad42"},
		Comment:        "This review was liked by viewer. posted by poster",
	},
}
