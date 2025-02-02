package orm

import (
	"food-shuffle-api/repository/model"

	"gorm.io/gorm"
)

func GetCourses(db *gorm.DB, restaurantUuid string) ([]model.Course, error) {
	var courses []model.Course
	err := db.Where("restaurant_uuid = ?", restaurantUuid).Find(&courses).Error
	return courses, err
}

func GetSpecificCourse(tx *gorm.DB, courseUuid string) (model.Course, error) {
	var course model.Course
	err := db.Where("course_uuid = ?", courseUuid).First(&course).Error
	return course, err
}
