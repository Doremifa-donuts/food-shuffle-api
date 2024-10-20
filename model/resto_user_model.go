package model

type BusyStatus int

const (
	Free   BusyStatus = iota // 暇
	Spare                    // 余裕
	Packed                   // 満席
)

type RestoUser struct {
	RestoUuid     string     `gorm:"type:char(36);primary_key"`                     // データの管理を楽にするためだけのカラム　サロゲートキー
	MailAddress   string     `gorm:"type:varchar(255);not null"`                    // メールアドレス
	Password      string     `gorm:"type:varchar(255);not null"`                    // パスワード
	Resto_name    string     `gorm:"type:varchar(100);not null"`                    // レストラン名
	Address       string     `gorm:"type:varchar(255);not null"`                    // 住所
	Tell          string     `gorm:"type:varchar(20);not null"`                     // 電話番号
	Images        []string   `gorm:"type:json;not null"`                            // 画像のパスをjsonの配列で格納することによって複数保存することが可能になる
	Url           string     `gorm:"type:varchar(255);not null"`                    // WebサイトなどのURL
	Summary       string     `gorm:"type:TEXT;not null"`                            // 店舗概要
	BusinessHours string     `gorm:"type:varchar(50);not null"`                     // 営業時間　なんか文字書く人とかいそうだし、文字列で格納
	JtiToken      string     `gorm:"type:char(36);not null"`                        // ログインした
	BusyStatus    BusyStatus `gorm:"type:enum('Free', 'Spare', 'Packed');not null"` // 0: 空席 1: 余裕あり 2: 満席
}
