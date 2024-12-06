package service

import (
	"errors"
	"fmt"
	"food-shuffle-api/bcrypto"
	"food-shuffle-api/dto"
	logging "food-shuffle-api/log"
	"food-shuffle-api/model"
	"food-shuffle-api/repository"
	"food-shuffle-api/utility/auth"
	"food-shuffle-api/utility/custom_error"
	"net/http"

	"github.com/google/uuid"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geo"
	"gorm.io/gorm"
)

type GeneralUserService struct{}

// 一般ユーザーのアカウントを作成し、トークンを返す
func (service *GeneralUserService) Register(bUser model.User, generalUser model.GeneralUser) (res dto.LoginUser, err error) {
	// トランザクションを開始する
	err = repository.Transaction(func(tx *gorm.DB) error {

		// メールアドレスを元にユーザーが存在するかを確認する
		_, err := repository.GetUserByMailAddress(tx, bUser.MailAddress)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) { // 存在しない以外のエラーがある場合
			logging.LogError("failed to get user", err)
			return err
		}
		if err == nil { // メールアドレスが一致するユーザーが存在する場合
			logging.LogError("users mail address already exists", err)
			return custom_error.NewError(http.StatusConflict, "User already exists")
		}

		// UUIDを生成する
		uuid, err := uuid.NewRandom()
		if err != nil {
			logging.LogError("failed to generate user uuid", err)
			return err
		}

		// それぞれのテーブルにuuidを挿入する
		bUser.UserUuid = uuid.String()
		generalUser.UserUuid = uuid.String()

		// パスワードをハッシュ化する
		hashedPassword, err := bcrypto.GetHashPassword(bUser.Password)
		if err != nil {
			logging.LogError("failed to hash password", err)
			return err
		}

		// ハッシュ化したパスワードに入れ替える
		bUser.Password = hashedPassword

		// ユーザータイプを一般に設定する
		bUser.UserType = model.General

		// 挿入するデータが完成したのでここから挿入していく
		// ユーザーテーブルに一般ユーザーを追加する
		err = repository.CreateUser(tx, bUser)
		if err != nil {
			logging.LogError("failed to create user", err)
			return err
		}

		// 一般ユーザーテーブルに追加情報を追加する
		err = repository.CreateGeneralUser(tx, generalUser)
		if err != nil {
			logging.LogError("failed to create general user", err)
			return err
		}

		// トークンを発行する
		res.JtiToken, err = auth.GenerateToken(tx, &bUser)
		if err != nil {
			logging.LogError("failed to generate token", err)
			return err
		}

		return nil
	})

	// レスポンスを返却する
	return
}

// レストランの詳細情報を取得する
func (service *GeneralUserService) GetRestaurantDetail(uuid string) (res dto.RestaurantDetail, err error) {
	err = repository.Transaction(func(tx *gorm.DB) error {
		//特定のUUIDに一致するレストランの情報を取得
		restaurantDetail, err := repository.GetRestaurantDetail(tx, uuid)
		if err != nil {
			logging.LogError("failed to get restaurant detail", err)
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return custom_error.NewError(http.StatusBadRequest, "restaurant user is not exist by restaurant uuid")
			}
			return err
		}

		//取得したデータをレスポンスの構造体にバインドする
		res = dto.RestaurantDetail{
			RestaurantUuid: restaurantDetail.RestaurantUuid,
			RestaurantName: restaurantDetail.RestaurantName,
			Address:        restaurantDetail.Address,
			Images:         restaurantDetail.Images,
			Url:            restaurantDetail.Url,
			Summary:        restaurantDetail.Summary,
			BusinessHours:  restaurantDetail.BusinessHours,
		}
		return nil
	})
	return
}

// チェックイン処理を行う
func (s *GeneralUserService) PostCheckInRestaurant(userUuid string, restaurantUuid string, latlong dto.CheckInLocation) (err error) {
	// トランザクションを開始する
	err = repository.Transaction(func(tx *gorm.DB) error {
		// 位置情報の構造体をorb.Point型に変換する
		location := orb.Point{latlong.Location.Latitude, latlong.Location.Longitude}

		// チェックインを行う店舗の位置情報を取得
		err := repository.Transaction(func(tx *gorm.DB) error {
			// レストランの詳細を取得
			restaurant, err := repository.GetRestaurantDetail(tx, restaurantUuid)
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return custom_error.NewError(http.StatusBadRequest, "restaurant user is not found by restaurant uuid")
				}
				return err
			}

			// レストランの位置情報を取得
			restaurantLocation := orb.Point{restaurant.Latitude, restaurant.Longitude}
			distance := geo.Distance(location, restaurantLocation)

			fmt.Println("userLocation", location)
			fmt.Println("restaurantLocation", restaurantLocation)
			fmt.Println("distance", distance)
			// CHECK IN DISTANCE	//TODO:変数の切り出し
			var CHECK_IN_DISTANCE float64 = 20
			if distance < CHECK_IN_DISTANCE {
				// チェックインテーブルの構造体
				userVisited := model.UserVisitedRestaurant{
					UserUuid:       userUuid,
					RestaurantUuid: restaurantUuid,
				}

				// 訪れたことがあるかを確認する
				ok, err := repository.ExistsUserVisitedRestaurant(tx, userVisited)
				fmt.Println("行ったことがある", ok)
				fmt.Println("err", err)
				if err != nil {
					return err
				}
				if ok {
					// 最終訪問日の更新を行う
					err = repository.UpdateLastVisitedTime(tx, userVisited)
					if err != nil {
						return err
					}
				} else {
					fmt.Println("ここやってる")
					// 初回訪問なのでレコードの追加を行う
					err = repository.CreateUserVisitedRestaurant(tx, userVisited)
					fmt.Print("err", err)
					if err != nil {
						return err
					}
				}

			} else {
				logging.LogError("this user is too far form the restaurant", nil)
				return custom_error.NewError(http.StatusUnprocessableEntity, "this user is too far from the restaurant")
			}
			return nil
		})
		return err
	})
	return err
}

func (s *GeneralUserService) PutShareStatus(generalUser model.GeneralUser) (err error) {

	err = repository.Transaction(func(tx *gorm.DB) error {
		result, err := repository.PutShareStatus(tx, generalUser)
		if err != nil {
			logging.LogError("failed to put share status", err)
			return err
		}
		if !result {
			logging.LogError("failed to put share status", err)
			return custom_error.NewError(http.StatusBadRequest, "General user not found")
		}

		// トランザクションを終了する
		return nil
	})
	return
}