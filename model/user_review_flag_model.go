package model

import "time"

type ReviewStatus string

const (
	Unclassified  ReviewStatus = "Unclassified"
	NotInterested ReviewStatus = "NotInterested"
	Interested    ReviewStatus = "Interested"
	Iiked         ReviewStatus = "Iiked"
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
		ReviewUuid:   "e08505ac-eb06-43ea-a29b-b206367f7b8d",
		UserUuid:     "91a78381-f472-496b-90e3-2c66a33391d1",
		ReviewStatus: Interested,
	},
	{
		ReviewUuid:   "573fa1e4-1510-4eaf-9f1f-9df903bbd020",
		UserUuid:     "91a78381-f472-496b-90e3-2c66a33391d1",
		ReviewStatus: Iiked,
	},
	{
		ReviewUuid:   "39f93b17-c378-46f4-b55e-0c65642d99b0",
		UserUuid:     "91a78381-f472-496b-90e3-2c66a33391d1",
		ReviewStatus: Unclassified,
	},
}
