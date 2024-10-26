package model

type ShareStatus string

const (
	Active ShareStatus = "Active"
	Silent ShareStatus = "Silent"
	Disabled ShareStatus = "Disabled"
)

// 一般利用者特有の情報
type GeneralUser struct {
	UserUuid    string `gorm:"type:char(36);not null;references:UserUuid;references:users;primary_key;"`
	Username    string       `gorm:"type:varchar(50);not null" `           // ユーザー名
	Tell        string       `gorm:"type:varchar(20);not null"`           // 電話番号
	ShareStatus ShareStatus  `gorm:"type:enum('Active', 'Silent', 'Disabled');not null"` // 共有ステータス Active: 通知あり Silent: 通知なし Disabled: 無効
}

// サンプルデータ
var SampleGeneralUsers = []GeneralUser{
	{
		UserUuid:    "a0adb027-0f54-4c1a-9ed3-86041c863344",
		Username:    "test_user",
		Tell:        "08012341234",
		ShareStatus: Active,
	},
}
