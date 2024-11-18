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
