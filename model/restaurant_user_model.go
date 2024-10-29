package model

type BusyStatus string

const (
	Free   BusyStatus = "Free" // 暇
	Spare                BusyStatus = "Spare"   // 余裕
	Packed               BusyStatus = "Packed"    // 満席
)

type RestaurantUser struct {
	UserUuid     string     `gorm:"type:char(36);references:UserUuid;references:users;primary_key"`                     // データの管理を楽にするためだけのカラム　サロゲートキー
	Resto_name    string     `gorm:"type:varchar(100);not null"`                    // レストラン名
	Address       string     `gorm:"type:varchar(255);not null"`                    // 住所
	Images        []string   `gorm:"type:json;not null"`                            // 画像のパスをjsonの配列で格納することによって複数保存することが可能になる
	Url           string     `gorm:"type:varchar(255);not null"`                    // WebサイトなどのURL
	Summary       string     `gorm:"type:TEXT;not null"`                            // 店舗概要
	BusinessHours string     `gorm:"type:varchar(50);not null"`                     // 営業時間　なんか文字書く人とかいそうだし、文字列で格納
	BusyStatus    BusyStatus `gorm:"type:enum('Free', 'Spare', 'Packed');not null"` // Free: 空席 Spare: 余裕あり Packed: 満席
}

var SampleRestaurantUsers = RestaurantUser{
	UserUuid:      "97961bc1-70c9-43ea-9b4e-18f8bb6574f8",
	Resto_name:    "sample_resto_name",
	Address:       "sample_address",
	Images:        []string{"sample_image"},
	Url:           "sample_url",
	Summary:       "sample_summary",
	BusinessHours: "sample_business_hours",
	BusyStatus:    Free,
}