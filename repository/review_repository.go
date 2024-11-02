package repository

import (
	"food-shuffle-api/model"
	"gorm.io/gorm"
)


func CreateReview(db *gorm.DB, review *model.Review) error {
	return db.Create(review).Error
}