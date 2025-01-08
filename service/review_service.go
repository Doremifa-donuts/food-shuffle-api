package service

import (
	"food-shuffle-api/model"
	"food-shuffle-api/repository"
	"time"

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

type ReviewListResponse struct {
	UserName  string
	Images    []string
	CreatedAt time.Time
	Comment   string
}

func (s *ReviewService) GetReviewsByRestaurant(uuid string) ([]ReviewListResponse, error) {
	// 返り値を宣言する
	var ReviewListResponses []ReviewListResponse

	// トランザクションを開始する
	err := repository.Transaction(func(tx *gorm.DB) error {
		// レビュー一覧を取得する
		reviews, err := repository.GetReviewsByRestaurantUuid(tx, uuid)
		if err != nil {
			return err
		}

		for _, review := range reviews {
			// レビュー情報を格納する
			ReviewListResponse := ReviewListResponse{
				Images:    review.Images,
				CreatedAt: review.CreatedAt,
				Comment:   review.Comment,
			}

			// ユーザー情報を格納する
			user, err := repository.GetGeneralUserByUserUuid(tx, review.UserUuid)
			if err != nil {
				return err
			}
			ReviewListResponse.UserName = user.UserName

			ReviewListResponses = append(ReviewListResponses, ReviewListResponse)
		}

		return nil
	})

	// 返り値を返す
	return ReviewListResponses, err
}
