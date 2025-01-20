package model

import "time"

// レビューテーブル
type Review struct {
	ReviewUuid              string                    `gorm:"type:char(36);primary_key"`               // レビューのUUID
	UserUuid                string                    `gorm:"type:char(36);foreignKey:UserUuid"`       // レビューを投稿したユーザーのUUID
	RestaurantUuid          string                    `gorm:"type:char(36);foreignKey:RestaurantUuid"` // レビューを投稿したレストランのUUID
	Images                  StringArray               `gorm:"type:json;not null"`                      // レビューに関連する画像のパスをJSONで保存する
	CreatedAt               time.Time                 `gorm:"not null"`                                // レビューを投稿した日時
	Comment                 string                    `gorm:"type:text;not null"`                      // レビューのコメント
	UserReviewFlags         []UserReviewFlag          `gorm:"foreignKey:ReviewUuid"`
	PopupGroupSharedReviews []PopupGroupSharedReviews `gorm:"foreignKey:ReviewUuid"`
}

// サンプルデータ
var sampleReviews = []Review{

	{
		ReviewUuid:     "0194657f-ce88-7106-b597-956627ca0c3c",
		UserUuid:       "a0adb027-0f54-4c1a-9ed3-86041c863344",
		RestaurantUuid: "97961bc1-70c9-43ea-9b4e-18f8bb6574f8",
		Images:         StringArray{"0194657f-ce82-7d3e-a62b-0802edc909eb.jpg", "0194657f-ce85-76ac-89d4-223cd3d25f96.jpg"},
		Comment:        "落ち着いた雰囲気の中で、本格的なビストロ料理を楽しめる素敵なお店です。素材の味を生かした一皿一皿が丁寧に作られており、味わい深く心に残ります。ワインとのペアリングも絶妙で、特別な時間を過ごせる空間でした。",
	},
	{
		ReviewUuid:     "0194656d-bb86-725e-bac0-5e8009a8eb05",
		UserUuid:       "a0adb027-0f54-4c1a-9ed3-86041c863344",
		RestaurantUuid: "a80499ae-eb6c-1305-a5cc-e1510c52744a",
		Images:         StringArray{"0194656f-16d0-734b-89e7-5e236e1460e2.jpg"},
		Comment:        "小さな店内ながら、香ばしい焼き鳥の香りが広がり、居心地の良さを感じます。一串一串が丁寧に焼かれており、外はカリッと中はジューシー。素材の良さと職人技が光る、通いたくなるお店です。",
	},
	{
		ReviewUuid:     "0194656b-cd71-7775-b489-a8d37ca623a7",
		UserUuid:       "a0adb027-0f54-4c1a-9ed3-86041c863344",
		RestaurantUuid: "0bf97fc8-019e-421b-85f5-84818aab19d8",
		Images:         StringArray{"0194656b-cd6e-7295-b158-baf162839eff.jpg"},
		Comment:        "おにぎりの握り具合が絶妙で、口に含むとふわっとほどけるような食感が楽しめます。具材もたっぷりと詰められており、一つ一つに丁寧さとこだわりを感じられる逸品でした。ぜひまた訪れたいお店です。",
	},
	{
		ReviewUuid:     "01946573-c51b-7d62-93d4-e0442a59d3e0",
		UserUuid:       "a0adb027-0f54-4c1a-9ed3-86041c863344",
		RestaurantUuid: "d61aed9f-68b0-4efd-af77-98d7e061526d",
		Images:         StringArray{"01946573-c518-7960-b93e-56b805775912.jpg"},
		Comment:        "旬のフルーツがこれでもかと盛り付けられた「ボンボン」は、見た目も味も贅沢そのもの。一口ごとに異なるフルーツの甘みや酸味が広がり、最後まで飽きることなく楽しめます。美しいビジュアルと洗練された味わいに感動しました！",
	},
	{
		ReviewUuid:     "019465a5-33d0-76f1-99a0-527ffdff6251",
		UserUuid:       "a0adb027-0f54-4c1a-9ed3-86041c863344",
		RestaurantUuid: "6d7c3625-a1fa-4d63-8600-39f538dcac87",
		Images:         StringArray{"0194659b-8282-7b48-a049-697f83fcee57.png"},
		Comment:        "しゃぶしゃぶが美味しい",
		CreatedAt:      time.Date(2024, time.December, 14, 0, 0, 0, 0, time.Local),
	},
	{
		ReviewUuid:     "019465a5-8670-76be-a2e0-45855c448be2",
		UserUuid:       "a0adb027-0f54-4c1a-9ed3-86041c863344",
		RestaurantUuid: "5923b6b8-a4d6-4419-acf1-b1410480b0b5",
		Images:         StringArray{"0194659b-d0b9-70ff-85fb-cbf2271b5ec8.png"},
		Comment:        "スープが濃厚で、麺ももちもち。大満足。",
		CreatedAt:      time.Date(2024, time.December, 2, 0, 0, 0, 0, time.Local),
	},
}
