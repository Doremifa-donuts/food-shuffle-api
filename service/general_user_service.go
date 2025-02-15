package service

import (
	crypto_rand "crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"food-shuffle-api/bcrypto"
	"food-shuffle-api/dto"
	logging "food-shuffle-api/log"
	"food-shuffle-api/repository/fireauth"
	"food-shuffle-api/repository/model"
	"food-shuffle-api/repository/orm"
	"food-shuffle-api/repository/redis"
	"food-shuffle-api/utility/auth"
	"food-shuffle-api/utility/custom_error"
	"food-shuffle-api/utility/parameters"
	"food-shuffle-api/utility/prefix"
	"io"
	"math/rand/v2"
	"net/http"

	"github.com/google/uuid"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geo"
	"gorm.io/gorm"
)

type GeneralUserService struct{}

// 一般ユーザーのアカウント作成のために入力されたデータが正しいものかを判定する
func (s *GeneralUserService) PreRegister(req dto.PreRegisterRequest) (res dto.PreRegisterResponse, err error) {
	// 送信されたデータに未入力の項目がないかを確認する
	if req.MailAddress == "" || req.UserName == "" || req.Password == "" || req.ConfirmPassword == "" || req.Tell == "" {
		return res, custom_error.NewError(http.StatusBadRequest, "Some fields are incomplete. ")
	}

	// パスワードと入力の確認が一致しているかを確認する
	if req.Password != req.ConfirmPassword {
		return res, custom_error.NewError(http.StatusBadRequest, "password is not match")
	}
	// パスワードを暗号化する
	// パスワードをハッシュ化する
	hashedPassword, err := bcrypto.GetHashPassword(req.Password)
	if err != nil {
		logging.LogError("failed to hash password", err)
		return res, err
	}

	// データベースに保存しているものと重複していないことを確認する
	err = orm.Transaction(func(tx *gorm.DB) error {
		// メールアドレスが登録済みでないことを確認する
		result, err := orm.ExistsUserByMailAddress(tx, req.MailAddress)
		if err != nil {
			return err
		}
		if result {
			return custom_error.NewError(http.StatusConflict, "This email address is already in use.")
		}
		// 電話番号が登録済みでないことを確認する
		result, err = orm.ExistsUserByTell(tx, req.Tell)
		if err != nil {
			return err
		}
		if result {
			return custom_error.NewError(http.StatusConflict, "This tell is already in use.")
		}
		return nil
	})
	if err != nil {
		return res, err
	}

	// 期限を1時間にしてredisに仮データを保管する
	randBytes := make([]byte, 64)
	_, err = io.ReadFull(crypto_rand.Reader, randBytes)
	if err != nil {
		return res, err
	}

	// バイト列のデータをBase64でエンコーディングし、文字列を生成
	res.Key = base64.RawURLEncoding.WithPadding(base64.NoPadding).EncodeToString(randBytes)

	// redisに仮データを登録
	userInfo := redis.UserInfo{
		MailAddress: req.MailAddress,
		UserName:    req.UserName,
		Password:    hashedPassword,
		Tell:        req.Tell,
	}
	// jsonに変換
	value, err := json.Marshal(userInfo)
	if err != nil {
		return res, err
	}
	// Redisに仮登録データを保存
	err = redis.CachePreRegistrationUser(res.Key, value)

	return
}

// 一般ユーザーのアカウントを作成し、トークンを返す
func (service *GeneralUserService) Register(req dto.RegisterRequest) (res dto.LoginUser, err error) {
	// idTokenがfirebaseによって生成されたものかを確認する
	token, err := fireauth.VerifyIDToken(req.Token)
	if err != nil {
		return res, custom_error.NewError(http.StatusUnauthorized, "id token can not verification.")
	}
	// トークンが電話番号によって得られたものをか確認する
	if token.Firebase.SignInProvider != "phone" {
		return res, custom_error.NewError(http.StatusUnauthorized, "sign in provider is not phone.")
	}

	// 仮登録キーから保存されているユーザーデータを取得する
	result, err := redis.GetPreRegistrationUser(req.PreRegisterKey)
	if err != nil {
		return res, err
	}
	
	// 仮登録キーに一致するユーザーデータが見つからなかったとき
	if result == nil {
		return res, custom_error.NewError(http.StatusBadRequest, "pre-register-key is not available.")
	}
	// jsonになっている仮登録キーを復元する
	var userInfo redis.UserInfo
	err = json.Unmarshal(result, &userInfo)
	if err != nil {
		return res, err
	}

	// 仮登録と電話番号認証で使われた電話番号が同一であることを確認する
	if token.Claims["phone_number"] != userInfo.Tell {
		return res, custom_error.NewError(http.StatusUnauthorized, "token phone number not equals pre-register number.")
	}

	// トランザクションを開始する
	err = orm.Transaction(func(tx *gorm.DB) error {
		// メールアドレスが登録済みでないことを確認する
		result, err := orm.ExistsUserByMailAddress(tx, userInfo.MailAddress)
		if err != nil {
			return err
		}
		if result {
			return custom_error.NewError(http.StatusConflict, "This email address is already in use.")
		}
		// 電話番号が登録済みでないことを確認する
		result, err = orm.ExistsUserByTell(tx, userInfo.Tell)
		if err != nil {
			return err
		}
		if result {
			return custom_error.NewError(http.StatusConflict, "This tell is already in use.")
		}

		// UUIDを生成する
		uuid, err := uuid.NewRandom()
		if err != nil {
			logging.LogError("failed to generate user uuid", err)
			return err
		}

		// それぞれのテーブルにuuidを挿入する
		userUuid := uuid.String()
		user := model.User{
			UserUuid:    userUuid,
			MailAddress: userInfo.MailAddress,
			Password:    userInfo.Password,
			Tell:        userInfo.Tell,
			UserType:    model.General,
		}
		// HACK: ユーザーアイコンは初期はランダム設定にしている
		userIcons := []string{"0193c880-bae4-7f4e-b6f2-9582e1f0dac1.png", "0193c880-e065-7e8b-9e0c-9f333cb92ceb.png", "0193c880-fbc1-7fcc-a7e6-a95b0547368a.png"}

		genUser := model.GeneralUser{
			UserUuid: userUuid,
			UserName: userInfo.UserName,
			Icon:     userIcons[rand.IntN(len(userIcons))],
		}

		// 挿入するデータが完成したのでここから挿入していく
		// ユーザーテーブルに一般ユーザーを追加する
		err = orm.CreateUser(tx, user)
		if err != nil {
			logging.LogError("failed to create user", err)
			return err
		}

		// 一般ユーザーテーブルに追加情報を追加する
		err = orm.CreateGeneralUser(tx, genUser)
		if err != nil {
			logging.LogError("failed to create general user", err)
			return err
		}

		// トークンを発行する
		res.JtiToken, err = auth.GenerateToken(tx, &user)
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
	err = orm.Transaction(func(tx *gorm.DB) error {
		//特定のUUIDに一致するレストランの情報を取得
		restaurantDetail, err := orm.GetRestaurantDetail(tx, uuid)
		if err != nil {
			logging.LogError("failed to get restaurant detail", err)
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return custom_error.NewError(http.StatusBadRequest, "restaurant user is not exist by restaurant uuid")
			}
			return err
		}

		// 電話番号を取得する
		tell, err := orm.GetTellByRestaurantUuid(tx, restaurantDetail.RestaurantUuid)
		if err != nil {
			return err
		}

		// 画像IDにプレフィックスをつける
		var prefixedImages []string
		for _, image := range restaurantDetail.Images {
			prefixedImages = append(prefixedImages, prefix.ImagePrefixRestaurant+image)
		}

		//取得したデータをレスポンスの構造体にバインドする
		res = dto.RestaurantDetail{
			RestaurantUuid: restaurantDetail.RestaurantUuid,
			RestaurantName: restaurantDetail.RestaurantName,
			Address:        restaurantDetail.Address,
			Tell:           tell,
			Images:         prefixedImages,
			Url:            restaurantDetail.Url,
			Summary:        restaurantDetail.Summary,
			BusinessHours:  restaurantDetail.BusinessHours,
			BusyStatus:     dto.BusyStatus(restaurantDetail.BusyStatus),
		}
		return nil
	})
	return
}

// チェックイン処理を行う
func (s *GeneralUserService) PostCheckInRestaurant(userUuid string, restaurantUuid string, latlong dto.CheckInLocation) (err error) {
	// トランザクションを開始する
	err = orm.Transaction(func(tx *gorm.DB) error {
		// 位置情報の構造体をorb.Point型に変換する
		location := orb.Point{latlong.Location.Latitude, latlong.Location.Longitude}

		// チェックインを行う店舗の位置情報を取得
		err := orm.Transaction(func(tx *gorm.DB) error {
			// レストランの詳細を取得
			restaurant, err := orm.GetRestaurantDetail(tx, restaurantUuid)
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

			if distance < parameters.CHECK_IN_RADIUS {
				// チェックインテーブルの構造体
				userVisited := model.UserVisitedRestaurant{
					UserUuid:       userUuid,
					RestaurantUuid: restaurantUuid,
				}

				// 訪れたことがあるかを確認する
				ok, err := orm.ExistsUserVisitedRestaurant(tx, userVisited)
				fmt.Println("行ったことがある", ok)
				fmt.Println("err", err)
				if err != nil {
					return err
				}
				if ok {
					// 最終訪問日の更新を行う
					err = orm.UpdateLastVisitedTime(tx, userVisited)
					if err != nil {
						return err
					}
				} else {
					fmt.Println("ここやってる")
					// 初回訪問なのでレコードの追加を行う
					err = orm.CreateUserVisitedRestaurant(tx, userVisited)
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

	err = orm.Transaction(func(tx *gorm.DB) error {
		result, err := orm.PutShareStatus(tx, generalUser)
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

// 訪れた店の詳細情報をリストで取得する
func (service *GeneralUserService) GetIsReviewedRestaurants(isReviewed bool, userUuid string) (res []dto.RestaurantDetail, err error) {
	//トランザクションを開始する
	err = orm.Transaction(func(tx *gorm.DB) error {
		// 自身のレビューからレストランのUUID一覧を取得する
		reviewedRestaurantUuids, err := orm.ListRestaurantUuidsByUserUuidFromReview(tx, userUuid)
		if err != nil {
			logging.LogError("failed get reviewed restaurant uuid list", err)
			return err
		}

		var restaurantUuids []string
		fmt.Println(reviewedRestaurantUuids)
		if isReviewed {
			// レビューを書いている店舗の場合はそのまま代入する
			restaurantUuids = reviewedRestaurantUuids
		} else {
			// レビューを書いていないものの場合は訪問店舗リストからレビューを書いていない店のリストを取得する
			restaurantUuids, err = orm.ListFilterRestaurantUuidsByUserUuidNotInRestaurantUuids(tx, userUuid, reviewedRestaurantUuids)
			if err != nil {
				return err
			}
		}

		// レストランUUIDから詳細を取得する
		restaurants, err := orm.ListRestaurantByRestaurantUuids(tx, restaurantUuids)
		if err != nil {
			return err
		}

		// 不足している情報を追加する
		for _, restaurant := range restaurants {
			// 画像にプレフィックスをつける
			var prefixedImages []string
			for _, image := range restaurant.Images {
				prefixedImages = append(prefixedImages, prefix.ImagePrefixRestaurant+image)
			}

			// 電話番号を取得
			tell, err := orm.GetTellByRestaurantUuid(tx, restaurant.RestaurantUuid)
			if err != nil {
				return err
			}

			// 1件のレストラン情報を作成
			item := dto.RestaurantDetail{
				RestaurantUuid: restaurant.RestaurantUuid,
				RestaurantName: restaurant.RestaurantName,
				Address:        restaurant.Address,
				Tell:           tell,
				Images:         prefixedImages,
				Url:            restaurant.Url,
				Summary:        restaurant.Summary,
				BusinessHours:  restaurant.BusinessHours,
			}

			// レスポンスの構造体に追加
			res = append(res, item)
		}

		return nil
	})
	return
}

func (service *GeneralUserService) GetWentPlaces(uuid string) ([]dto.WentPlaces, error) {
	var res []dto.WentPlaces

	err := orm.Transaction(func(tx *gorm.DB) error {
		wentPlaces, err := orm.GetWentPlaces(tx, uuid)
		if err != nil {
			logging.LogError("failed to get went places", err)
			return err
		}

		res = wentPlaces
		return nil
	})

	return res, err
}
