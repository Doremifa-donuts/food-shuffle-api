package model

import "time"

// 個人がいいねしているレビューを管理するテーブル
type ReviewLike struct {
	ReviewUuid string    `gorm:"type:char(36);foreignKey:ReviewUuid;primary_key"` // レビューUUID
	UserUuid   string    `gorm:"type:char(36);foreignKey:UserUuid;primary_key"`   // ユーザーUUID
	CreatedAt  time.Time `gorm:"not null"`                                        // 作成日時によってソートされる
}

// サンプルデータ
var sampleReviewLikes = []ReviewLike{
	{
		ReviewUuid: "573fa1e4-1510-4eaf-9f1f-9df903bbd020",
		UserUuid:   "91a78381-f472-496b-90e3-2c66a33391d1",
	},
}
