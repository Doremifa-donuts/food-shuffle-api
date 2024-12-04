package model

type BusyStatus string

const (
	Free   BusyStatus = "Free"   // 暇
	Spare  BusyStatus = "Spare"  // 余裕
	Packed BusyStatus = "Packed" // 満席
)

type RestaurantUser struct {
	RestaurantUuid         string                  `gorm:"type:char(36);primary_key"`  // データの管理を楽にするためだけのカラム　サロゲートキー
	RestaurantName         string                  `gorm:"type:varchar(100);not null"` // レストラン名
	Address                string                  `gorm:"type:varchar(255);not null"` // 住所
	Images                 StringArray             `gorm:"type:json;not null"`         // 画像のパスをjsonの配列で格納することによって複数保存することが可能になる
	Url                    string                  `gorm:"type:varchar(255);not null"` // WebサイトなどのURL
	Summary                string                  `gorm:"type:TEXT;not null"`         // 店舗概要
	BusinessHours          string                  `gorm:"type:varchar(50);not null"`  // 営業時間　なんか文字書く人とかいそうだし、文字列で格納
	BusyStatus             BusyStatus              `gorm:"type:enum('Free', 'Spare', 'Packed');default:Free;not null"`
	Reservations           []Reservation           `gorm:"foreignKey:RestaurantUuid"` // Free: 空席 Spare: 余裕あり Packed: 満席
	Reviews                []Review                `gorm:"foreignKey:RestaurantUuid"`
	Courses                []Course                `gorm:"foreignKey:RestaurantUuid"`
	UrgentCampaigns        []UrgentCampaign        `gorm:"foreignKey:RestaurantUuid"`
	UserVisitedRestaurants []UserVisitedRestaurant `gorm:"foreignKey:RestaurantUuid"`
}

var sampleRestaurantUsers = []RestaurantUser{
	{
		RestaurantUuid: "97961bc1-70c9-43ea-9b4e-18f8bb6574f8",
		RestaurantName: "sample_resto_name",
		Address:        "東京都千代田区千代田１−１",
		Images:         StringArray{"sample_image1", "sample_image2"},
		Url:            "http://google.co.jp",
		Summary:        "sample_summary",
		BusinessHours:  "sample_business_hours",
		BusyStatus:     Free,
	},
	{
		RestaurantUuid: "a80499ae-eb6c-1305-a5cc-e1510c52744a",
		RestaurantName: "sample_res_name",
		Address:        "sample_add",
		Images:         StringArray{"sample_image3", "sample_image4", "sample_image5"},
		Url:            "http://google.co.jp",
		Summary:        "sample_summary",
		BusinessHours:  "sample_business_hours",
		BusyStatus:     Free,
	},
}
