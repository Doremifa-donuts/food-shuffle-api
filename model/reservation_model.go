package model

import "time"

type Reservation struct {
	ReservationUuid   string    `gorm:"type:char(36);primary_key"`
	RestaurantUuid    string    `gorm:"type:char(36);foreignKey:RestaurantUuid"`
	UserUuid          string    `gorm:"type:char(36);foreignKey:UserUuid"`
	ReservationDate   time.Time `gorm:"type:date;not null"`
	NumberOfPeople    int       `gorm:"type:integer;not null"`
	CourseUuid        *string   `gorm:"type:char(36)"`
	CampaignUuid      *string   `gorm:"type:char(36)"`
	ReservationStatus bool      `gorm:"type:bool;not null"`
}

// サンプルデータ
var sampleReservation = []Reservation{
	(
	ReservationUuid:   "05e804a3-57f6-4bb7-8d58-6d2a6b9d74ba",
	RestoUuid:         "97961bc1-70c9-43ea-9b4e-18f8bb6574f8",
	UserUuid:          "a0adb027-0f54-4c1a-9ed3-86041c863344",
	ReservationDate:   time.Now(),
	NumberOfPeople:    1,
	CourseNo:          nil,
	UrgentCampaignNo:  nil,
	ReservationStatus: true,
	),
	(
	ReservationUuid:   "2",
	RestoUuid:         "97961bc1-70c9-43ea-9b4e-18f8bb6574f8",
	UserUuid:          "a0adb027-0f54-4c1a-9ed3-86041c863344",
	ReservationDate:   time.Now(),
	NumberOfPeople:    1,
	CourseNo:          nil,
	UrgentCampaignNo:  nil,
	)
}
