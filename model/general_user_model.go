package model

type ShareStatus string

const (
	Active   ShareStatus = "Active"
	Silent   ShareStatus = "Silent"
	Disabled ShareStatus = "Disabled"
)

// 一般利用者特有の情報
type GeneralUser struct {
	UserUuid            string                  `gorm:"type:char(36);primary_key;"`
	UserName            string                  `gorm:"type:varchar(50);not null" `                                        // ユーザー名
	ShareStatus         ShareStatus             `gorm:"type:enum('Active', 'Silent', 'Disabled');default:Active;not null"` // 共有ステータス Active: 通知あり Silent: 通知なし Disabled: 無効
	Icon                string                  `gorm:"type:char(36);not null"`
	Reviews             []Review                `gorm:"foreignKey:UserUuid"`
	UserReviewFlags     []UserReviewFlag          `gorm:"foreignKey:UserUuid"`
	Reservations        []Reservation           `gorm:"foreignKey:UserUuid"`
	PopupGroups         []PopupGroupSubmission  `gorm:"foreignKey:UserUuid"`
	UserVisitedRestaurants []UserVisitedRestaurant `gorm:"foreignKey:UserUuid"`
}

// サンプルデータ
var sampleGeneralUsers = []GeneralUser{
	{
		UserUuid:    "a0adb027-0f54-4c1a-9ed3-86041c863344",
		UserName:    "poster",
		ShareStatus: Active,
		Icon:        "fcec1eba-758f-4da8-a9a3-49457fd3b6fe",
	},
	{
		UserUuid:    "91a78381-f472-496b-90e3-2c66a33391d1",
		UserName:    "viewer",
		ShareStatus: Silent,
		Icon:        "4c1e4635-1950-41e8-a561-be995a4a8816",
	},
}
