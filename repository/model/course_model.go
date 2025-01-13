package model

type Course struct {
	CourseUuid     string        `gorm:"type:char(36);primary_key;"`                       // コース番号　複合主キー
	RestaurantUuid string        `gorm:"type:char(36);foreignKey:RestaurantUuid;not null"` // お店のUUID
	CourseName     string        `gorm:"type:varchar(50);not null"`                        // コース名
	Description    string        `gorm:"type:text;not null"`                               // コースの説明
	Images         StringArray   `gorm:"type:json;not null"`                               // 画像
	Price          int           `gorm:"type:integer;not null"`                            // 金額
	Minimum        int           `gorm:"type:integer;not null"`                            // 最小人数
	Reservations   []Reservation `gorm:"foreignKey:CourseUuid"`
}

// サンプルデータ
var sampleCourses = []Course{
	{
		CourseUuid:     "a6814a18-5426-482d-ad6a-ab5b4218619e",
		RestaurantUuid: "97961bc1-70c9-43ea-9b4e-18f8bb6574f8",
		CourseName:     "定番人気コース",
		Description:    "気軽に楽しめる定番メニューです。豚のしょうが焼きがメイン。",
		Images:         StringArray{"0193c8c9-1535-7338-bcd5-7c5f93f3844f.jpg"},
		Price:          1200,
		Minimum:        1,
	},
	{
		CourseUuid:     "4b665e72-873e-4318-a46a-93749e1e5302",
		RestaurantUuid: "97961bc1-70c9-43ea-9b4e-18f8bb6574f8",
		CourseName:     "お肉好きコース",
		Description:    "ボリュームたっぷり！唐揚げ定食がメインの満足コース。",
		Images:         StringArray{"0193c8c8-6153-7721-b868-120431b0e332.jpg", "0193c8c9-8139-7687-804e-e6dfc5318a2d.jpg"},
		Price:          1500,
		Minimum:        1,
	},
	{
		CourseUuid:     "0a4d87ca-623c-4641-b29b-3cd475f5f0cd",
		RestaurantUuid: "97961bc1-70c9-43ea-9b4e-18f8bb6574f8",
		CourseName:     "ヘルシー野菜コース",
		Description:    "野菜たっぷりの健康志向コース。チキン南蛮がメイン。サラダもおかわり自由です。",
		Images:         StringArray{"0193c8ce-fe1e-777e-98ec-6e41544aca55.jpg", "0193c8ce-6d67-73a2-ab84-4b2e742b38ff.jpg"},
		Price:          1300,
		Minimum:        1,
	},
	{
		CourseUuid:     "839ab143-17ed-42a0-b286-60b3885b4ac9",
		RestaurantUuid: "97961bc1-70c9-43ea-9b4e-18f8bb6574f8",
		CourseName:     "シェアコース",
		Description:    "グループで楽しめる盛り合わせコースです。",
		Images:         StringArray{"0193c8d1-6f11-71b0-a9a3-ecde128659f9.jpg", "0193c8d1-05d6-7b3f-b452-5ff38e1c586e.jpg"},
		Price:          2000,
		Minimum:        2,
	},
}
