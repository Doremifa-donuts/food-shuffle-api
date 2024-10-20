package model

type User struct {
	UserUuid    string `gorm:"type:char(36);not null;primary_key;"` // ユーザーのUUID
	Username    string `gorm:"type:varchar(50);not null"`           // ユーザー名
	MailAddress string `gorm:"type:varchar(255);not null"`          // メールアドレス
	Password    string `gorm:"type:varchar(255);not null"`          // パスワード
	Tell        string `gorm:"type:varchar(20);not null"`           // 電話番号
	JtiToken    string `gorm:"type:varchar(255);not null"`          // JTIトークン
	ShareStatus int    `gorm:"type:varchar(255);not null"`          // "1: 通知を受け取る / 2: 通知を受け取らない / 3: 共有を行わない"
}
