package model

type ShareStatus string

const (
	Active ShareStatus =  "Active"
	Silent ShareStatus = "Silent"
	Disabled      ShareStatus = "Disabled"
)

type User struct {
	UserUuid    string       `gorm:"type:char(36);not null;primary_key;"` // ユーザーのUUID
	Username    string       `gorm:"type:varchar(50);not null" `           // ユーザー名
	MailAddress string       `gorm:"type:varchar(255);not null"`          // メールアドレス
	Password    string       `gorm:"type:varchar(255);not null"`          // パスワード
	Tell        string       `gorm:"type:varchar(20);not null"`           // 電話番号
	JtiToken    string       `gorm:"type:varchar(255);not null"`          // JTIトークン
	ShareStatus ShareStatus  `gorm:"type:enum('Active', 'Silent', 'Disabled');default:'Silent';not null"` // 共有ステータス
}

// サンプルデータ
var SampleUsers = []User{
	{
		UserUuid:    "0192aa4a-2e3f-7000-a78d-4830ada1b887",
		Username:    "test_user",
		MailAddress: "test@test.com",
		Password:    "$2a$10$UkrQfUmAsPJ35cw5TVzJeOuoLySOWpMHN/b2zN561eixU0abBSCpq",	// test
		Tell:        "08012341234",
		JtiToken:    "test_jti_token",
		ShareStatus: Active,
	},
}
