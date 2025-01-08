package repository

import (
	"food-shuffle-api/model"

	"gorm.io/gorm"
)

func CreateReview(db *gorm.DB, review *model.Review) error {
	return db.Create(review).Error
}

func GetReviewsByRestaurantUuid(db *gorm.DB, uuid string) ([]model.Review, error) {
	var reviews []model.Review
	err := db.Where("restaurant_uuid = ?", uuid).Find(&reviews).Error
	if err != nil {
		return nil, err
	}
	return reviews, nil
}
