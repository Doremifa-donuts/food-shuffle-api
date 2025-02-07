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
	{
		ReviewUuid:   "019465a5-8670-76be-a2e0-45855c448be2",
		UserUuid:     "91a78381-f472-496b-90e3-2c66a33391d1",
		ReviewStatus: Unclassified,
	},
	{
		ReviewUuid:   "0194de87-aed4-7972-be26-13d18d58c9b2", // らーめん大王
		UserUuid:     "91a78381-f472-496b-90e3-2c66a33391d1",
		ReviewStatus: Unclassified,
	},
	// {
	// 	ReviewUuid:   "0194656b-cd71-7775-b489-a8d37ca623a7", // おにぎり
	// 	UserUuid:     "91a78381-f472-496b-90e3-2c66a33391d1",
	// 	ReviewStatus: Unclassified,
	// },
	{
		ReviewUuid:   "0194de89-3d75-7200-81db-dca8a2969c1e", // ビストロ
		UserUuid:     "91a78381-f472-496b-90e3-2c66a33391d1",
		ReviewStatus: Unclassified,
	},
	{
		ReviewUuid:   "0194de76-c2f2-71f6-9c4c-0fd397047142", // らーめん大王
		UserUuid:     "91a78381-f472-496b-90e3-2c66a33391d1",
		ReviewStatus: Unclassified,
	},
	{
		ReviewUuid:   "0194de76-fc3c-72ed-93a5-5f4ed033d91f", // おにぎり
		UserUuid:     "91a78381-f472-496b-90e3-2c66a33391d1",
		ReviewStatus: Unclassified,
	},
	{
		ReviewUuid:   "0194de77-2bbe-7d16-8214-7a89cc41f597", // ビストロ
		UserUuid:     "91a78381-f472-496b-90e3-2c66a33391d1",
		ReviewStatus: Unclassified,
	},
	{
		ReviewUuid:   "0194de77-5c1d-7877-8ebb-436e8a16c1a8", // らーめん大王
		UserUuid:     "91a78381-f472-496b-90e3-2c66a33391d1",
		ReviewStatus: Unclassified,
	},
	{
		ReviewUuid:   "0194de77-f924-7698-91aa-6e4146b0cdd6", // おにぎり
		UserUuid:     "91a78381-f472-496b-90e3-2c66a33391d1",
		ReviewStatus: Unclassified,
	},
	{
		ReviewUuid:   "0194de78-2b3b-793d-8006-86aabff0306c", // ビストロ
		UserUuid:     "91a78381-f472-496b-90e3-2c66a33391d1",
		ReviewStatus: Unclassified,
	},
	{
		ReviewUuid:   "0194de78-671c-7b6b-a84e-0e6e9055f9a7", // らーめん大王
		UserUuid:     "91a78381-f472-496b-90e3-2c66a33391d1",
		ReviewStatus: Unclassified,
	},
	// {
	// 	ReviewUuid:   "0194de78-671c-7b6b-a84e-0e6e9055f9a7", // らーめん大王
	// 	UserUuid:     "91a78381-f472-496b-90e3-2c66a33391d1",
	// 	ReviewStatus: Unclassified,
	// },
	{
		ReviewUuid:   "0194de78-8e8c-7949-b2b0-8ce79a2b11cb", // らーめん大王
		UserUuid:     "91a78381-f472-496b-90e3-2c66a33391d1",
		ReviewStatus: Unclassified,
	},
	{
		ReviewUuid:   "0194de78-c7ca-776e-8ec3-1ec8c08fb961", // おにぎり
		UserUuid:     "91a78381-f472-496b-90e3-2c66a33391d1",
		ReviewStatus: Unclassified,
	},
	{
		ReviewUuid:   "0194de78-ed4c-76b5-9d00-84e60a5fdd21", // ビストロ
		UserUuid:     "91a78381-f472-496b-90e3-2c66a33391d1",
		ReviewStatus: Unclassified,
	},
}
