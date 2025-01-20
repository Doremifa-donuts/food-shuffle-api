package model

import "time"

type ReviewStatus string

const (
	Unclassified  ReviewStatus = "Unclassified"  // 未分類
	NotInterested ReviewStatus = "NotInterested" // 興味なし
	Interested    ReviewStatus = "Interested"    // 興味あり
	Iiked         ReviewStatus = "Iiked"         // いいね
)

type UserReviewFlag struct {
	ReviewUuid   string       `gorm:"type:char(36);foreignKey:ReviewUuid;primary_key"`
	UserUuid     string       `gorm:"type:char(36);foreignKey:UserUuid;primary_key"`
	CreatedAt    time.Time    `gorm:"not null"`
	ReviewStatus ReviewStatus `gorm:"type:enum('Unclassified', 'NotInterested', 'Interested', 'Iiked');default:Unclassified;not null"`
}

// サンプルデータ
var sampleUserReviewFlag = []UserReviewFlag{
	{
		ReviewUuid:   "0194657f-ce88-7106-b597-956627ca0c3c",
		UserUuid:     "91a78381-f472-496b-90e3-2c66a33391d1",
		ReviewStatus: Unclassified,
	},
	{
		ReviewUuid:   "0194656d-bb86-725e-bac0-5e8009a8eb05",
		UserUuid:     "91a78381-f472-496b-90e3-2c66a33391d1",
		ReviewStatus: Unclassified,
	},
	{
		ReviewUuid:   "0194656b-cd71-7775-b489-a8d37ca623a7",
		UserUuid:     "91a78381-f472-496b-90e3-2c66a33391d1",
		ReviewStatus: Unclassified,
	},
	{
		ReviewUuid:   "01946573-c51b-7d62-93d4-e0442a59d3e0",
		UserUuid:     "91a78381-f472-496b-90e3-2c66a33391d1",
		ReviewStatus: Unclassified,
	},
	{
		ReviewUuid:   "019465a5-33d0-76f1-99a0-527ffdff6251",
		UserUuid:     "91a78381-f472-496b-90e3-2c66a33391d1",
		ReviewStatus: Unclassified,
	},
}
