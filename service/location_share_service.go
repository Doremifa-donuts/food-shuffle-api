package service

import (
	"errors"
	"fmt"
	logging "food-shuffle-api/log"
	"food-shuffle-api/model"
	"food-shuffle-api/redisconn"
	"food-shuffle-api/repository"
	"food-shuffle-api/utility/custom_error"
	"regexp"
	"strconv"

	"gorm.io/gorm"
)

type LocationShareService struct{}

func (s *LocationShareService) NotifyReviewByLocationMessage(userUuid string, latlong []byte) (notifyUserUuids []string, err error) {
	reg := "\r\n|\n"
	// WSからデータを取得した際に決められた位置情報の形に整形する
	latLong := regexp.MustCompile(reg).Split(string(latlong), -1)
	latitude, _ := strconv.ParseFloat(latLong[0], 64)
	longitude, err := strconv.ParseFloat(latLong[1], 64)
	if err != nil {
		logging.LogError("can not parse float latlong", err)
		return
	}

	// 自分のデータをredisに登録する
	redisconn.SetGeoLocation(userUuid, latitude, longitude)

	// レビューを受け取ったユーザーのリスト
	var receiveUserUuids []string
	// 共有するレビューの情報
	var shareSettingReview model.ShareSettingReview

	// 共有するレビューに設定しているリストを取得する
	err = repository.Transaction(func(tx *gorm.DB) error {
		// 共有するレビューがあるかを確認する
		shareSettingReview, err = repository.GetShareSettingReviewByUserUuid(tx, userUuid)
		if err != nil {
			return err
		}
		if shareSettingReview.FirstReview == nil {
			logging.LogError("No reviews available for sharing", nil)
			//HACK: カスタムエラーにステータスコード入れるな
			return custom_error.NewError(0, "No reviews available for sharing")
		}
		fmt.Println("シェアしたいレビュー:", *shareSettingReview.FirstReviewUuid)

		// レビューの共有と追加の処理を行う
		// レビューのいいね数を取得する
		shareRadius, err := repository.CountReviewLikesByReviewUuid(tx, *shareSettingReview.FirstReviewUuid)
		if err != nil {
			logging.LogError("can not get review like count", err)
			return err
		}
		shareRadius += 10 // 最小共有範囲を10mに設定する //TODO: 切り出し
		fmt.Println("共有範囲:", shareRadius)

		// 共有範囲にいる人のリストを取得する
		withinUserUuids, err := redisconn.GetUserUuidsByReviewShareRadius(userUuid, float64(shareRadius+10))
		if err != nil {
			logging.LogError("colud not get share target user uuid", err)
			return err
		}
		fmt.Println("共有範囲内にいる人たち:", withinUserUuids)

		// 共有対象がいない場合は早期リターン
		if len(withinUserUuids) <= 0 {
			return custom_error.NewError(0, "There is no one within the sharing range of the review.")
		}

		// レビューを所持していない人のみに絞り込む
		receiveUserUuids, err = repository.ListExcludeUserUuidByReviewUuid(tx, *shareSettingReview.FirstReviewUuid, withinUserUuids)
		if err != nil {
			logging.LogError("could not get revieve target user", err)
			return err
		}
		fmt.Println("レビューを所持していない人のみ:", receiveUserUuids)
		// 共有対象がいない場合は早期リターン
		if len(receiveUserUuids) <= 0 {
			return custom_error.NewError(0, "There is no one don't have review.")
		}

		// 受け取ったレビューリストに追加する
		err = repository.CreateUserReviewFlag(tx, *shareSettingReview.FirstReviewUuid, receiveUserUuids)
		if err != nil {
			logging.LogError("colud not create review flag", err)
			return err
		}

		// ここまでで共有の処理は完了
		return nil
	})

	// 通知について処理を行う
	repository.Transaction(func(tx *gorm.DB) error {
		// 店が混雑しているかを取得
		err = repository.CheckNotPackedStatusByRestaurantUuid(tx, shareSettingReview.FirstReview.RestaurantUuid)
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				logging.LogError("colud not check restaurant is packed", err)
				//HACK: カスタムエラーにステータスコード入れるな
				return custom_error.NewError(0, "There is no one within the sharing range of the review.")
			}
			return err
		}

		// 通知モードをオンにしている人のみ通知リストに格納
		notifyUserUuids, err = repository.ListFilterActiveStatusByUserUuids(tx, receiveUserUuids)
		if err != nil {
			logging.LogError("colud not filter user uuid", err)
			return err
		}
		// トランザクションを終了する
		return nil
	})

	return
}
