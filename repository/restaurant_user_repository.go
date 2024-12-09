package repository

import (
	"food-shuffle-api/model"

	"gorm.io/gorm"
)

// ユーザーを取得する（busy_status確認用）
func GetRestaurantUserByUuid(db *gorm.DB, uuid string) (model.RestaurantUser, error) {
	var restoUser model.RestaurantUser
	err := db.Where("restaurant_uuid = ?", uuid).First(&restoUser).Error
	if err != nil {
		return restoUser, err
	}
	return restoUser, nil
}

// 混雑状況を更新する
func UpdateBusyStatus(db *gorm.DB, uuid string, busyStatusInput model.BusyStatus) error {
	return db.Model(&model.RestaurantUser{}).Where("restaurant_uuid = ?", uuid).Update("busy_status", busyStatusInput).Error
}
