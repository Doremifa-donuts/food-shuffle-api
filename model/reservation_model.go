package model

import "time"

type Reservation struct {
	ReservationUuid   string    `gorm:"type:char(36);primary_key"`
	RestoUuid         string    `gorm:"type:char(36);foreignkey:RestoUuid"`
	UserUuid          string    `gorm:"type:char(36);foreignkey:UserUuid"`
	ReservationDate   time.Time `gorm:"type:date;not null"`
	NumberOfPeople    int       `gorm:"type:integer;not null"`
	CourseNo          *int      `gorm:"type:integer;foreignkey:CourseNo;;"`
	UrgentCampaignNo  *int      `gorm:"type:integer;foreignkey:UrgentCampaignNo;"`
	ReservationStatus bool      `gorm:"type:bool;not null"`
}
