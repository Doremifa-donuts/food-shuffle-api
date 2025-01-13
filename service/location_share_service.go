package service

import (
	"errors"
	logging "food-shuffle-api/log"
	"food-shuffle-api/repository/model"
	"food-shuffle-api/repository/orm"
	"food-shuffle-api/repository/redis"
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
	redis.SetGeoLocation(userUuid, latitude, longitude)

	// レビューを受け取ったユーザーのリスト
	var receiveUserUuids []string
	// 共有するレビューの情報
	var shareSettingReview model.ShareSettingReview

	// 共有するレビューに設定しているリストを取得する
	err = orm.Transaction(func(tx *gorm.DB) error {
		// 共有するレビューがあるかを確認する
		shareSettingReview, err = orm.GetShareSettingReviewByUserUuid(tx, userUuid)
		if err != nil {
			return err
		}
		if shareSettingReview.FirstReview == nil {
			logging.LogError("No reviews available for sharing", nil)
			//HACK: カスタムエラーにステータスコード入れるな
			return custom_error.NewError(0, "No reviews available for sharing")
		}

		// レビューの共有と追加の処理を行う
		// レビューのいいね数を取得する
		shareRadius, err := orm.CountReviewLikesByReviewUuid(tx, *shareSettingReview.FirstReviewUuid)
		if err != nil {
			logging.LogError("can not get review like count", err)
			return err
		}
		shareRadius += 10 // 最小共有範囲を10mに設定する //TODO: 切り出し

		// 共有範囲にいる人のリストを取得する
		withinUserUuids, err := redis.GetUserUuidsByReviewShareRadius(userUuid, shareRadius)
		if err != nil {
			logging.LogError("colud not get share target user uuid", err)
			return err
		}

		// 共有対象がいない場合は早期リターン
		if len(withinUserUuids) <= 0 {
			return custom_error.NewError(0, "There is no one within the sharing range of the review.")
		}

		// レビューを所持していない人のみに絞り込む
		receiveUserUuids, err = orm.ListExcludeUserUuidByReviewUuid(tx, *shareSettingReview.FirstReviewUuid, withinUserUuids)
		if err != nil {
			logging.LogError("could not get revieve target user", err)
			return err
		}

		// 共有対象がいない場合は早期リターン
		if len(receiveUserUuids) <= 0 {
			return custom_error.NewError(0, "There is no one don't have review.")
		}

		// 受け取ったレビューリストに追加する
		err = orm.CreateUserReviewFlag(tx, *shareSettingReview.FirstReviewUuid, receiveUserUuids)
		if err != nil {
			logging.LogError("colud not create review flag", err)
			return err
		}

		// ここまでで共有の処理は完了
		return nil
	})

	if err == nil {

		// 通知について処理を行う
		err = orm.Transaction(func(tx *gorm.DB) error {
			// 店が混雑しているかを取得
			err = orm.CheckNotPackedStatusByRestaurantUuid(tx, shareSettingReview.FirstReview.RestaurantUuid)
			if err != nil {
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					logging.LogError("colud not check restaurant is packed", err)
					//HACK: カスタムエラーにステータスコード入れるな
					return custom_error.NewError(0, "There is no one within the sharing range of the review.")
				}
				return err
			}

			// 通知モードをオンにしている人のみ通知リストに格納
			notifyUserUuids, err = orm.ListFilterActiveStatusByUserUuids(tx, receiveUserUuids)
			if err != nil {
				logging.LogError("colud not filter user uuid", err)
				return err
			}
			// トランザクションを終了する
			return nil
		})
	}
	return
}
