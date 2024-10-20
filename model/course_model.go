package model

type Course struct {
	CourseUuid  string   `gorm:"type:char(36);primary_key"`                // お店のUUID
	CourseNo    int      `gorm:"type:integer;auto_increment;primary_key;"` // コース番号　複合主キー
	CourseName  string   `gorm:"type:varchar(50);not null"`                // コース名
	Description string   `gorm:"type:text;not null"`                       // コースの説明
	Images      []string `gorm:"type:json;not null"`                       // 画像
	Price       int      `gorm:"type:varchar(7);not null"`                 // 金額
	Minimum     int      `gorm:"type:integer;not null"`                    // 最小人数
}
