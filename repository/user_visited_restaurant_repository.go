package repository

import (
	"food-shuffle-api/model"

	"gorm.io/gorm"
)

func ExistsUserVisitedRestaurantByUserUuid(db *gorm.DB, userUuid string) error {
	var userVisitedRestaurant model.UserVisitedRestaurant
	err := db.Where("user_uuid = ?", userUuid).First(&userVisitedRestaurant).Error
	if err != nil {
		return err
	}

	return nil
}
