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
		ReviewUuid:     "39f93b17-c378-46f4-b55e-0c65642d99b0",
		UserUuid:       "a0adb027-0f54-4c1a-9ed3-86041c863344",
		RestaurantUuid: "97961bc1-70c9-43ea-9b4e-18f8bb6574f8",
		Images:         StringArray{"0193c904-4d2c-7771-836a-be2a763ad8c2.jpg"},
		Comment:        "彩り豊かで品数豊富な和食御膳。お刺身や煮物、焼き魚、炊き込みご飯など、一品一品が丁寧に盛り付けられていました。特にお刺身の新鮮さと、煮物の優しい味付けが印象的です。見た目からも楽しめる美しい構成で、季節の食材を活かした内容に満足しました。ボリュームも適量で、最後のフルーツまで美味しくいただけました。和食好きにはたまらない一品です。",
	},
	{
		ReviewUuid:     "e08505ac-eb06-43ea-a29b-b206367f7b8d",
		UserUuid:       "a0adb027-0f54-4c1a-9ed3-86041c863344",
		RestaurantUuid: "97961bc1-70c9-43ea-9b4e-18f8bb6574f8",
		Images:         StringArray{"0193c908-5884-704f-b377-703909411f03.jpg"},
		Comment:        "ボリューム満点な唐揚げ定食に驚きました！ジューシーな鶏肉とカリッとした衣が絶妙で、一口食べるごとに満足感が広がります。ご飯との相性も抜群で、ついついお箸が進みます。付け合わせのキャベツやお味噌汁もシンプルながら嬉しい一品。量が多くて満腹になること間違いなしで、ガッツリ食べたい時にはピッタリのメニューです！",
	},
	{
		ReviewUuid:     "573fa1e4-1510-4eaf-9f1f-9df903bbd020",
		UserUuid:       "a0adb027-0f54-4c1a-9ed3-86041c863344",
		RestaurantUuid: "97961bc1-70c9-43ea-9b4e-18f8bb6574f8",
		Images:         StringArray{"0193c916-0499-76f3-a8f7-003bc26bab2b.jpg", "0193c915-91c5-7841-9a94-11bbbd16b0e1.jpg", "0193c915-3808-7243-83d5-63e3b94100ce.jpg"},
		Comment:        "新鮮な刺身の盛り合わせは、どれも色鮮やかで、魚の旨味がしっかり感じられました。特にマグロが柔らかくて絶品でした。天ぷらはサクサクで、エビや季節の野菜の甘みが引き立ち、油っぽさもなく軽やかな食感でした。おにぎりは、ふっくらとしたご飯に具材がしっかり混ざり、香りも豊かで食べ応え十分でした。さらに、他にもさまざまな種類のご飯が提供され、どれも風味豊かで満足感がありました。全体的に、丁寧に作られた料理が揃っていて、非常に満足のいく食事でした。",
	},
}
