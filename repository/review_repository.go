package repository

import (
	"food-shuffle-api/model"
	"gorm.io/gorm"
)

// レビューを作成する
func CreateReview(db *gorm.DB, review *model.Review) error {
	return db.Create(review).Error
}

// レビューUUIDからレビューを複数取得する
func ListReviewsByReviewUuids(db *gorm.DB, reviewUuids []string) ([]model.Review, error) {
	var reviews []model.Review
	err := db.Where("review_uuid IN (?)", reviewUuids).Find(&reviews).Error
	return reviews, err
}