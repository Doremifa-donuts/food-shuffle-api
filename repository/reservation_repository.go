package repository

import (
	"gorm.io/gorm"

	"food-shuffle-api/model"
)

// 予約の登録を行う
func CreateReservation(db *gorm.DB, reservation model.Reservation) error {
	return db.Create(&reservation).Error
}

