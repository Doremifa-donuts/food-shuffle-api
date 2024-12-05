package repository

import (
	"food-shuffle-api/model"

	"gorm.io/gorm"
)

// レストランUUIDからレストラン名を取得する
func GetRestaurantNameByRestaurantUuid(db *gorm.DB, uuid string) (string, error) {
	var restoUser model.RestaurantUser
	err := db.Where("restaurant_uuid = ?", uuid).Find(&restoUser).Error
	return restoUser.RestaurantName, err
}

// 混雑状況が満席でないことを確認する
func CheckNotPackedStatusByRestaurantUuid(db *gorm.DB, restaurantUuid string) error {
	return db.Where("restaurant_uuid = ? and busy_status <> ?", restaurantUuid, model.Packed).First(&model.RestaurantUser{}).Error
}

// レストランの詳細情報を取得する
func GetRestaurantDetail(db *gorm.DB, RestaurantUuid string) ([]model.RestaurantUser, error) {
	var restaurantUser []model.RestaurantUser
	err := db.Where("restaurant_uuid = ?", RestaurantUuid).Find(&restaurantUser).Error
	return restaurantUser, err
}

func PutBusyStatus(db *gorm.DB, restaurantUser model.RestaurantUser) (bool, error) {
	result := db.Model(&model.RestaurantUser{}).Where("restaurant_uuid = ?", restaurantUser.RestaurantUuid).Update("busy_status", restaurantUser.BusyStatus)
	return result.RowsAffected == 1, result.Error
}