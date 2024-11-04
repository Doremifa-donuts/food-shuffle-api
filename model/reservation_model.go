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
