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
	Latitude               float64                 `gorm:"type:float;not null"`        // 緯度
	Longitude              float64                 `gorm:"type:float;not null"`        // 経度
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
		RestaurantName: "街のご飯やさん",
		Address:        "東京都千代田区千代田１−１",
		Latitude:       35.685175,
		Longitude:      139.7528,
		Images:         StringArray{"0193c8df-a496-7d04-a897-0c4c374da9fa.jpg", "0193c8df-e6b0-7f0d-b4c6-d6fae17eaa01.jpg", "0193c8e1-a0de-7be0-88ae-04a53bb47d16.jpg"},
		Url:            "http://google.co.jp",
		Summary:        "「街のご飯やさん」は、誰もが気軽に立ち寄れるアットホームな定食屋です。定番の家庭料理からボリューム満点のメニューまで、リーズナブルな価格でご提供しています。おひとりさまでもグループでも大歓迎！温かい雰囲気で、皆さまのお腹も心も満たします。",
		BusinessHours:  "10:00 ~ 21:00 水曜定休日",
		BusyStatus:     Free,
	},
	{
		RestaurantUuid: "a80499ae-eb6c-1305-a5cc-e1510c52744a",
		RestaurantName: "よくある普通のカフェ",
		Address:        "sample_add",
		Latitude:       35.685175,
		Longitude:      139.7528,
		Images:         StringArray{"0193c8e6-eb48-7e5f-ac42-2e17f5aeb260.jpg", "0193c8e7-445b-770a-a91e-42c5f5a098e0.jpg", "0193c8e7-849d-73db-bc1f-09fb2a8625c2.jpg"},
		Url:            "http://google.co.jp",
		Summary:        "気軽に立ち寄れる、ほっと一息つけるカフェです。こだわりのコーヒーと手作りスイーツ、軽食をご用意してお待ちしています。おひとりでも、お友達とのおしゃべりにもぴったりな、心地よい空間です。",
		BusinessHours:  "11:00 ~ 20:00 不定休",
		BusyStatus:     Free,
	},
}
