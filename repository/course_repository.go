package repository

import (
	"food-shuffle-api/model"

	"gorm.io/gorm"
)

func GetCourseNameByCourseUuid(db *gorm.DB, courseUuid string) (string, error) {
	var course model.Course
	err := db.Where("course_uuid = ?", courseUuid).First(&course).Error
	if err != nil {
		return "", err
	}
	return course.CourseName, nil
}
