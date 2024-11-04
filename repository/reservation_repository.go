package repository

import (
	"food-shuffle-api/model"

	"gorm.io/gorm"
)

func GetReservationsByRestaurantUuid(db *gorm.DB, uuid string) ([]model.Reservation, error) {
	// RestoUuidに一致する予約を取得する
	var reservations []model.Reservation

	err := db.Where("resto_uuid = ?", uuid).Find(&reservations).Error
	if err != nil {
		return nil, err
	}
	return reservations, nil
}
