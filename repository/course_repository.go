package repository

import (
	"food-shuffle-api/model"

	"gorm.io/gorm"
)

func GetCourses(db *gorm.DB, restaurantUuid string) ([]model.Course, error) {
	var courses []model.Course
	err := db.Where("restaurant_uuid = ?", restaurantUuid).Find(&courses).Error
	return courses, err
}
