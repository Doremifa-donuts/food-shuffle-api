package service

import (
	"food-shuffle-api/model"
	"food-shuffle-api/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReviewService struct{}

func (s *ReviewService) PostReview(review model.Review) (string, error) {

	// ユーザーのレビューをDBに保存する
	// レビューUUIDを生成する
	reviewUuid, err := uuid.NewV7()
	if err != nil {
		return "", err
	}
	review.ReviewUuid = reviewUuid.String()

	// トランザクションを開始する
	err = repository.Transaction(func(tx *gorm.DB) error {

		// ユーザーのレビューをDBに保存する
		err := repository.CreateReview(tx, &review)
		if err != nil {
			return err
		}
		return nil
	})

	return review.ReviewUuid, err

}
