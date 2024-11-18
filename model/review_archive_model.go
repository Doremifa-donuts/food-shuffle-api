package model

import "time"

// 個人が保存しているレビューを管理するテーブル
type ReviewArchive struct {
	ReviewUuid string    `gorm:"type:char(36);foreignKey:ReviewUuid;primary_key"` // レビューのUUID
	UserUuid   string    `gorm:"type:char(36);foreignKey:UserUuid;primary_key"`   // ユーザーのUUID
	CreatedAt  time.Time `gorm:"not null"`                                        // 登録日によってソートする
}

// サンプルデータ
var sampleReviewArchives = []ReviewArchive{
	{
		ReviewUuid: "e08505ac-eb06-43ea-a29b-b206367f7b8d",
		UserUuid: "91a78381-f472-496b-90e3-2c66a33391d1",
	},
}
