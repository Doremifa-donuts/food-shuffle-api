package model

type ShareStatus string

const (
	Active   ShareStatus = "Active"
	Silent   ShareStatus = "Silent"
	Disabled ShareStatus = "Disabled"
)

// 一般利用者特有の情報
type GeneralUser struct {
	UserUuid               string                  `gorm:"type:char(36);primary_key;"`
	UserName               string                  `gorm:"type:varchar(50);not null" `                                        // ユーザー名
	ShareStatus            ShareStatus             `gorm:"type:enum('Active', 'Silent', 'Disabled');default:Active;not null"` // 共有ステータス Active: 通知あり Silent: 通知なし Disabled: 無効
	Icon                   string                  `gorm:"type:varchar(45);not null"`
	Reviews                []Review                `gorm:"foreignKey:UserUuid"`
	UserReviewFlags        []UserReviewFlag        `gorm:"foreignKey:UserUuid"`
	Reservations           []Reservation           `gorm:"foreignKey:UserUuid"`
	PopupGroups            []PopupGroupSubmission  `gorm:"foreignKey:UserUuid"`
	UserVisitedRestaurants []UserVisitedRestaurant `gorm:"foreignKey:UserUuid"`
	ShareSettingReview     ShareSettingReview      `gorm:"foreignKey:UserUuid"`
}

// サンプルデータ
var sampleGeneralUsers = []GeneralUser{
	{
		UserUuid:    "a0adb027-0f54-4c1a-9ed3-86041c863344",
		UserName:    "poster",
		ShareStatus: Active,
		Icon:        "0193c880-bae4-7f4e-b6f2-9582e1f0dac1.png",
	},
	{
		UserUuid:    "91a78381-f472-496b-90e3-2c66a33391d1",
		UserName:    "viewer",
		ShareStatus: Active,
		Icon:        "0193c880-e065-7e8b-9e0c-9f333cb92ceb.png",
	},
}
