package service

import (
	"errors"
	"food-shuffle-api/dto"
	logging "food-shuffle-api/log"
	"food-shuffle-api/model"
	"food-shuffle-api/repository"
	"food-shuffle-api/utility/custom_error"
	"food-shuffle-api/utility/img"
	"food-shuffle-api/utility/prefix"
	"mime/multipart"
	"net/http"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReviewService struct{}

// すれ違いで受け取ったレビューを取得する
func (s *ReviewService) GetReceivedReviewsByUser(uuid string) ([]dto.GetReviewsByUser, error) {
	// ステータスを Unclassified に限定した構造体に変換する
	reviewFlag := model.UserReviewFlag{
		UserUuid:     uuid,
		ReviewStatus: model.Unclassified,
	}
	// 受け取ったレビューを取得する
	return s.getReviewsByUser(reviewFlag)
}

// 興味ありに登録したレビューを取得する
func (s *ReviewService) GetInterestedReviewsByUser(uuid string) ([]dto.GetReviewsByUser, error) {
	// ステータスを Interested に限定した構造体に変換する
	reviewFlag := model.UserReviewFlag{
		UserUuid:     uuid,
		ReviewStatus: model.Interested,
	}
	// 興味あり状態のレビューを取得する
	return s.getReviewsByUser(reviewFlag)
}

// ユーザーがいいねしたレビューを取得する
func (s *ReviewService) GetLikedReviewsByUser(uuid string) ([]dto.GetReviewsByUser, error) {
	// ステータスをLikedに限定した構造体に変換する
	reviewFlag := model.UserReviewFlag{
		UserUuid:     uuid,
		ReviewStatus: model.Iiked,
	}
	// いいねしたレビューを格納するテーブルに対して、レビューを取得する
	return s.getReviewsByUser(reviewFlag)
}

// ユーザーのレビューを取得する
func (s *ReviewService) getReviewsByUser(reviewFlag model.UserReviewFlag) (res []dto.GetReviewsByUser, err error) {
	// トランザクションを開始する
	err = repository.Transaction(func(tx *gorm.DB) error {
		// ユーザーが取得するレビューのUUID
		var reviewUuids []string

		// ユーザーのレビューを取得する
		reviewUuids, err := repository.ListReviewUuidsByUserUuidAndReviewStatus(tx, reviewFlag)
		if err != nil {
			logging.LogError("failed to get reviews", err)
			return err
		}

		// レビューの構造体
		var reviews []model.Review

		// レビューの内容を取得する
		reviews, err = repository.ListReviewsByReviewUuids(tx, reviewUuids)
		if err != nil {
			logging.LogError("failed to get reviews", err)
			return err
		}

		// それぞれに不足している項目を取得する
		for _, review := range reviews {
			// レストラン名を取得する
			restaurantName, err := repository.GetRestaurantNameByRestaurantUuid(tx, review.RestaurantUuid)
			if err != nil {
				logging.LogError("failed to get restaurant name", err)
				return err
			}

			// いいね数をカウントする
			likes, err := repository.CountReviewLikesByReviewUuid(tx, review.ReviewUuid)
			if err != nil {
				return err
			}

			// Imagesにプレフィックスを追加する
			for i, image := range review.Images {
				review.Images[i] = prefix.ImagePrefixReview + image
			}

			// ユーザーのアイコンを取得する
			icon, err := repository.GetIconByUserUuid(tx, review.UserUuid)
			if err != nil {
				logging.LogError("failed to get icon", err)
				return err
			}

			// アイコンにプレフィックスを追加する
			icon = prefix.ImagePrefixUserIcon + icon

			// レビューをレスポンスに追加する
			res = append(res, dto.GetReviewsByUser{
				RestaurantUuid: review.RestaurantUuid,
				RestaurantName: restaurantName,
				ReviewUuid:     review.ReviewUuid,
				Comment:        review.Comment,
				PostedAt:       review.CreatedAt,
				Images:         review.Images,
				Icon:           icon,
				Good:           int(likes),
			})
		}

		// エラーが出なかった場合はnilでトランザクションを終了する
		return nil
	})

	// レスポンスを返却する
	return
}

// レビューのステータスを更新する
func (s *ReviewService) UpdateReviewStatus(bReviewFlag model.UserReviewFlag) (err error) {
	// トランザクションを開始する
	err = repository.Transaction(func(tx *gorm.DB) error {
		// レビューのステータスを更新する
		result, err := repository.UpdateReviewStatus(tx, bReviewFlag)
		if err != nil {
			logging.LogError("failed to update review status", err)
			return err
		}
		if !result {
			logging.LogError("failed to update review status", err)
			return custom_error.NewError(http.StatusBadRequest, "Review not found")
		}

		// トランザクションを終了する
		return nil
	})

	// エラーが出なかった場合はnilを返却する
	return
}

// ユーザーがレビューを投稿する
func (s *ReviewService) PostReview(bReview model.Review, images []*multipart.FileHeader) (res dto.PostReview, err error) {
	// トランザクションを開始する
	err = repository.Transaction(func(tx *gorm.DB) error {

		// その店舗に訪れたことがあるかを確認する
		// チェックインを記録する構造体
		userVisited := model.UserVisitedRestaurant{
			UserUuid:       bReview.UserUuid,
			RestaurantUuid: bReview.RestaurantUuid,
		}

		ok, err := repository.ExistsUserVisitedRestaurant(tx, userVisited)
		if err != nil {
			logging.LogError("failed query exists user visited restaurant table", err)
			return err
		}
		if !ok {
			// エラーログを書き込む
			logging.LogError("Your user has not visited the restaurant.", err)
			return custom_error.NewError(http.StatusForbidden, "Your user has not visited the restaurant.")
		}

		// 画像を保存する
		dirPath := "public/images/reviews"
		imagesPath, err := img.SaveImages(dirPath, images)
		if err != nil {
			return err
		}
		bReview.Images = imagesPath

		// TODO: トランザクション失敗した場合に保存した画像を破棄する

		// レビューUUIDを生成する
		reviewUuid, err := uuid.NewV7()
		if err != nil {
			logging.LogError("failed to generate review uuid", err)
			return err
		}

		// レビューUUIDを格納する
		bReview.ReviewUuid = reviewUuid.String()
		res.ReviewUuid = reviewUuid.String()

		// ユーザーのレビューをDBに保存する
		err = repository.CreateReview(tx, &bReview)
		if err != nil {
			logging.LogError("failed to create review", err)
			return err
		}

		// トランザクションを終了する
		return nil
	})

	// レスポンスを返却する
	return
}

func (s *ReviewService) SetShareReview(bShareSettingReview model.ShareSettingReview) (res dto.ShareSettingReview, err error) {
	// トランザクションを開始する
	err = repository.Transaction(func(tx *gorm.DB) error {
		// レビューUUIDが設定されているかを確認する
		if bShareSettingReview.FirstReviewUuid == nil {
			logging.LogError("do not set review uuid", nil)
			return custom_error.NewError(http.StatusBadRequest, "do not set review uuid")
		}
		// レビューが本人のものであるかを確認する
		err = repository.ExistsReviewByUserUuidAndReviewUuid(tx, bShareSettingReview.UserUuid, *bShareSettingReview.FirstReviewUuid)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				logging.LogError("review uuid is not own", err)
				return custom_error.NewError(http.StatusBadRequest, "review uuid is not this users")
			}
		}
		// すでに設定されたものがあるかを確認する
		_, err := repository.GetShareSettingReviewByUserUuid(tx, bShareSettingReview.UserUuid)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {

		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			// レコードがない場合は新規作成
			err = repository.CreateShareSettingReview(tx, bShareSettingReview)
			if err != nil {
				return err
			}
		} else {
			// レコードが存在する場合は更新
			err = repository.UpdateShareSettingReview(tx, bShareSettingReview)
			if err != nil {
				return err
			}
		}

		// 作成、更新したデータを取得する
		setReview, err := repository.GetShareSettingReviewByUserUuid(tx, bShareSettingReview.UserUuid)
		if err != nil {
			return err
		}

		// レスポンス生成
		if setReview.FirstReview != nil {
			res.FirstShareReviewUuid = *setReview.FirstReviewUuid
		}
		return nil
	})

	return
}

// 自身の書いたレビューの詳細を取得する
func (service *ReviewService) GetReviewDetail(RestaurantUuid string, userUuid string) (res dto.ReviewDetail, err error) {
	err = repository.Transaction(func(tx *gorm.DB) error {

		//特定のUUIDに一致するレストランの情報を取得
		reviewDetail, err := repository.GetReviewDetail(tx, RestaurantUuid, userUuid)
		if err != nil {
			logging.LogError("failed to get review detail", err)
			return err
		}

		// imagesにprefixをつける
		var prefixedImages []string
		for _, image := range reviewDetail.Images {
			prefixedImages = append(prefixedImages, prefix.ImagePrefixReview+image)
		}

		//取得したデータを格納する
		res = dto.ReviewDetail{
			ReviewUuid:     reviewDetail.ReviewUuid,
			UserUuid:       reviewDetail.UserUuid,
			RestaurantUuid: reviewDetail.RestaurantUuid,
			Images:         prefixedImages,
			CreatedAt:      reviewDetail.CreatedAt,
			Comment:        reviewDetail.Comment,
		}
		return nil
	})
	return
}
