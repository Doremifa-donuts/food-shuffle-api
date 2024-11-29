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
		CourseUuid:     "09732eaa-4883-1680-690f-a2958c0f82e7",
		RestaurantUuid: "97961bc1-70c9-43ea-9b4e-18f8bb6574f8",
		CourseName:     "コースA",
		Description:    "コースAの説明",
		Images:         StringArray{"image1.jpg", "image2.jpg"},
		Price:          1000,
		Minimum:        2,
	},
	{
		CourseUuid:     "07a9c86-3939-5ad0-49e8-c2fac4aaac1f",
		RestaurantUuid: "97961bc1-70c9-43ea-9b4e-18f8bb6574f8",
		CourseName:     "コースB",
		Description:    "コースBの説明",
		Images:         StringArray{"image3.jpg", "image4.jpg"},
		Price:          1500,
		Minimum:        3,
	},
	{
		CourseUuid:     "e5d61f9b-6c56-4009-97a2-460bdf7585e8",
		RestaurantUuid: "a80499ae-eb6c-1305-a5cc-e1510c52744a",
		CourseName:     "コースAA",
		Description:    "コースAAの説明",
		Images:         StringArray{"image5.jpg", "image6.jpg"},
		Price:          1800,
		Minimum:        3,
	},
	{
		CourseUuid:     "dc2ebc8e-19c3-47ed-8144-8ec71319b554",
		RestaurantUuid: "a80499ae-eb6c-1305-a5cc-e1510c52744a",
		CourseName:     "コースAB",
		Description:    "コースABの説明",
		Images:         StringArray{"image7.jpg", "image8.jpg"},
		Price:          2500,
		Minimum:        6,
	},
}
