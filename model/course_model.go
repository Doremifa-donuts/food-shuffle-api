package model

type Course struct {
	CourseUuid     string        `gorm:"type:char(36);primary_key;"`                       // コース番号　複合主キー
	RestaurantUuid string        `gorm:"type:char(36);foreignKey:RestaurantUuid;not null"` // お店のUUID
	CourseName     string        `gorm:"type:varchar(50);not null"`                        // コース名
	Description    string        `gorm:"type:text;not null"`                               // コースの説明
	CourseImages         StringArray   `gorm:"type:json;not null"`                               // 画像
	Price          int           `gorm:"type:integer;not null"`                            // 金額
	Minimum        int           `gorm:"type:integer;not null"`                            // 最小人数
	Reservations   []Reservation `gorm:"foreignKey:CourseUuid"`
}

// サンプルデータ
var sampleCourses = []Course{}
