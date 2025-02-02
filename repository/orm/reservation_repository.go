package orm

import (
	"time"

	"gorm.io/gorm"

	"food-shuffle-api/repository/model"
)

// 予約の登録を行う
func CreateReservation(db *gorm.DB, reservation model.Reservation) error {
	return db.Create(&reservation).Error
}

func GetReservationsByRestaurantUuid(db *gorm.DB, uuid string) ([]model.Reservation, error) {
	// RestoUuidに一致する予約を取得する
	var reservations []model.Reservation

	err := db.Where("restaurant_uuid = ?", uuid).Find(&reservations).Error
	if err != nil {
		return nil, err
	}
	return reservations, nil
}

// ユーザーが予約しているリストを取得する
func ListUpcomingReservationByUserUuid(tx *gorm.DB, userUuid string) ([]model.Reservation, error) {
	// 予約の構造体
	var reservations []model.Reservation
	err := db.Where("user_uuid = ? and reservation_date > ?", userUuid, time.Now()).Find(&reservations).Error
	return reservations, err
}
