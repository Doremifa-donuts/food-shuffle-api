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
		RestaurantName: "路地裏ビストロ816",
		Address:        "大阪府大阪市北区堂山町8-16",
		Latitude:       34.703527,
		Longitude:      135.502889,
		Images:         StringArray{"01946460-cda0-71f1-b383-03bfad01549b.jpg", "01946460-faa0-78ec-9ef1-e9b01ab56a0f.jpg", "01946461-16ab-7037-bd00-9ff16522cfae.jpg"},
		Url:            "https://www.instagram.com/bistro_816/",
		Summary:        "【梅田駅5分】ワインに合う肉料理や四季折々の一品を気軽に楽しめる、お洒落な隠れ家ビストロ",
		BusinessHours:  "月・火・水・木・金・土・日・祝日・祝前日・祝後日\n17:00 - 00:00",
		BusyStatus:     Free,
	},
	{
		RestaurantUuid: "a80499ae-eb6c-1305-a5cc-e1510c52744a",
		RestaurantName: "炭焼 高田屋",
		Address:        "大阪府大阪市北区神山町9-14 オリカ神山テナントビル 1F",
		Latitude:       34.703625819928256,
		Longitude:      135.5047858612578,
		Images:         StringArray{"01946457-92bf-726a-9bcf-d5afde519f39.jpg", "01946457-611d-7db2-9a76-96b0d8e5a258.jpg", "01946457-b3a2-7f96-a9b6-6e969628631b.jpg"},
		Url:            "https://kds7100.gorp.jp/",
		Summary:        "朝引き地鶏を1本1本丁寧に串打ちした炭火焼き鳥が自慢の【炭焼 高田屋】。新鮮な鶏のお造りからこだわりの創作料理やジャンクフードまで、お酒のアテにぴったりなアラカルトも多彩にご用意しています！生ビールをはじめ、ハイボール・サワー・焼酎・カクテルなどアルコール類も充実の品揃え◎広々とした店内には、各種お集まりにぴったりな掘りごたつタイプの個室も完備しています。老若男女問わずお気軽にご来店ください！",
		BusinessHours:  "18:00 - 08:00",
		BusyStatus:     Free,
	},
	{
		RestaurantUuid: "0bf97fc8-019e-421b-85f5-84818aab19d8",
		RestaurantName: "おにぎりごりちゃん 中崎町本店",
		Address:        "大阪府大阪市北区中崎1丁目5-20 TKビル1階",
		Latitude:       34.70712271496019,
		Longitude:      135.50587938159558,
		Images:         StringArray{"01946556-19e3-7f3b-a748-a176e4bc08c1.jpg"},
		Url:            "https://www.instagram.com/onigiri_gorichan_nakazaki/",
		Summary:        "大阪で大人気！行列2時間待ちの【おにぎりごりちゃん】が北陸に初上陸！注文が入ってから職人がふわっと握る！中まで具がたっぷり！ご馳走おにぎり！",
		BusinessHours:  "11:00〜20:00(お米が無くなり次第終了)",
		BusyStatus:     Free,
	},
	{
		RestaurantUuid: "d61aed9f-68b0-4efd-af77-98d7e061526d",
		RestaurantName: "Cafe de paris 大丸心斎橋店",
		Address:        "大阪府大阪市北区中崎1丁目5-20 TKビル1階",
		Latitude:       34.67294898135667,
		Longitude:      135.50081393721177,
		Images:         StringArray{"01946561-c0de-7a05-ada9-0568506f7d58.jpg", "01946562-1fba-7fe9-a422-77a88c7fd8f3.jpg"},
		Url:            "https://www.daimaru.co.jp/shinsaibashi/restaurant/cafedeparis.html",
		Summary:        "2009年に韓国で創業、「カフェ ド パリ」が大阪初登場。旬のフルーツをふんだんに使用して美しく盛り付けたフルーツパフェ「ボンボン」は、圧倒的な人気を誇る看板メニューです。他にも、季節のメニューや日本限定のメニューなど、フルーツたっぷりのご褒美スイーツをぜひお楽しみください。",
		BusinessHours:  "大丸心斎橋店の営業時間に準じます。",
		BusyStatus:     Free,
	},
	{
		RestaurantUuid: "6d7c3625-a1fa-4d63-8600-39f538dcac87",
		RestaurantName: "かごの屋 三国本町店",
		Address:        "大阪府大阪市淀川区三国本町２-13-9",
		Latitude:       34.73341050778812,
		Longitude:      135.48577446398468,
		Images:         StringArray{"0194b096-bf9f-7cb0-8923-ef2681765f56.jpg"},
		Url:            "https://kagonoya.food-kr.com/",
		Summary:        "家庭ではなかなかできないひと手間かけた和食を楽しんで頂けるよう、多彩な和の献立をご用意しています。旬の素材にこだわった季節のメニューを通じて、和食の魅力を伝えたい。「最高のごちそうさま」をお届けすること、それがかごの屋の理念です。",
		BusinessHours:  "月～金: 11:00～15:30 （料理L.O. 15:00 ドリンクL.O. 15:00）",
		BusyStatus:     Free,
	},
	{
		RestaurantUuid: "5923b6b8-a4d6-4419-acf1-b1410480b0b5",
		RestaurantName: "ラーメン大王 西中島店",
		Address:        "大阪府大阪市淀川区西中島7丁目1-13",
		Latitude:       34.72994680917593,
		Longitude:      135.49723484400593,
		Images:         StringArray{"0194b095-37d3-7e5e-aaca-7c174cc2b1f8.jpg"},
		Url:            "https://nishinakajima.ramennoodleclub.com/daiou/",
		Summary:        "本場仕込みの台湾ラーメンだけでなく、台湾中華の食べ放題やコスパ最強の飲み放題が楽しめる隠れ家的なお店。メニューの数もかなり多く何度も訪れたくなる台湾ラーメン大王。優しい味の焼き飯が絶品。",
		BusinessHours:  "11:30～15:00",
		BusyStatus:     Free,
	},
}
