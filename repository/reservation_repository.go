package repository

import (
	"food-shuffle-api/model"

	"gorm.io/gorm"
)

func GetReservationsByRestaurantUuid(db *gorm.DB, uuid string) ([]model.Reservation, error) {
	// RestoUuidに一致する予約を取得する
	var reservations []model.Reservation

	err := db.Where("restaurant_uuid = ?", uuid).Find(&reservations).Error
	if err != nil {
		return nil, err
	}
	return reservations, nil
}

func GetReservationByReservationUuid(db *gorm.DB, uuid string, reservation_uuid string) (model.Reservation, error) {
	var reservation model.Reservation
	err := db.Where("restaurant_uuid = ? AND reservation_uuid = ?", uuid, reservation_uuid).First(&reservation).Error
	if err != nil {
		return reservation, err
	}
	return reservation, nil
}

// 予約が通ってるかどうかを承認する
func UpdateReservationStatus(db *gorm.DB, uuid string, reservation_uuid string, status bool) error {
	return db.Model(&model.Reservation{}).Where("reservation_uuid = ?", reservation_uuid).Update("reservation_status", status).Error
}
