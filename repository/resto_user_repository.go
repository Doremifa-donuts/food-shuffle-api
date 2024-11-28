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

// レストランの詳細情報を取得する
func GetRestaurantDetail(db *gorm.DB, RestaurantUuid string) ([]model.RestaurantUser, error) {
	var restaurantUser []model.RestaurantUser
	err := db.Where("restaurant_uuid = ?", RestaurantUuid).Find(&restaurantUser).Error
	return restaurantUser, err
}