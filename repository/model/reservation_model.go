package model

import "time"

type Reservation struct {
	ReservationUuid   string    `gorm:"type:char(36);primary_key"`
	RestaurantUuid    string    `gorm:"type:char(36);foreignKey:RestaurantUuid"`
	UserUuid          string    `gorm:"type:char(36);foreignKey:UserUuid"`
	ReservationDate   time.Time `gorm:"not null"`
	NumberOfPeople    int       `gorm:"type:integer;not null"`
	CourseUuid        *string   `gorm:"type:char(36)"`
	CampaignUuid      *string   `gorm:"type:char(36)"`
	ReservationStatus bool      `gorm:"type:bool;not null"`
}

// サンプルデータ
var sampleReservations = []Reservation{
	{
		ReservationUuid:   "05e804a3-57f6-4bb7-8d58-6d2a6b9d74ba",
		RestaurantUuid:    "97961bc1-70c9-43ea-9b4e-18f8bb6574f8",
		UserUuid:          "a0adb027-0f54-4c1a-9ed3-86041c863344",
		ReservationDate:   time.Now(),
		NumberOfPeople:    1,
		CourseUuid:        nil,
		CampaignUuid:      nil,
		ReservationStatus: true,
	},
	{
		ReservationUuid: "b813c5e4-d76e-42c2-b1c7-d0c15eff884b",
		RestaurantUuid:  "97961bc1-70c9-43ea-9b4e-18f8bb6574f8",
		UserUuid:        "a0adb027-0f54-4c1a-9ed3-86041c863344",
		ReservationDate: time.Now(),
		NumberOfPeople:  1,
		CourseUuid:      nil,
		CampaignUuid:    nil,
	},
}
