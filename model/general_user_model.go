package model

type ShareStatus string

const (
	Active   ShareStatus = "Active"
	Silent   ShareStatus = "Silent"
	Disabled ShareStatus = "Disabled"
)

// 一般利用者特有の情報
type GeneralUser struct {
	UserUuid        string                 `gorm:"type:char(36);primary_key;"`
	UserName        string                 `gorm:"type:varchar(50);not null" `                                        // ユーザー名
	ShareStatus     ShareStatus            `gorm:"type:enum('Active', 'Silent', 'Disabled');default:Active;not null"` // 共有ステータス Active: 通知あり Silent: 通知なし Disabled: 無効
	Reviews         []Review               `gorm:"foreignKey:UserUuid"`
	ReviewReceives  []ReviewReceive        `gorm:"foreignKey:UserUuid"`
	ReviewArchives  []ReviewArchive        `gorm:"foreignKey:UserUuid"`
	ReviewFavorites []ReviewFavorite       `gorm:"foreignKey:UserUuid"`
	Reservations    []Reservation          `gorm:"foreignKey:UserUuid"`
	PopupGroups     []PopupGroupSubmission `gorm:"foreignKey:UserUuid"`
	// User            User              `gorm:"foreignKey:UserUuid"`
	User 			[]User 				   `gorm:"foreignKey:UserUuid;"`
}

// サンプルデータ
var SampleGeneralUsers = []GeneralUser{
	{
		UserUuid:    "a0adb027-0f54-4c1a-9ed3-86041c863344",
		UserName:    "test_user",
		ShareStatus: Active,
	},
}
