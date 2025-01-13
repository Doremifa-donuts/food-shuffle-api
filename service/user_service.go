package service

import (
	"fmt"
	"food-shuffle-api/bcrypto"
	"food-shuffle-api/dto"
	logging "food-shuffle-api/log"
	"food-shuffle-api/repository/model"
	"food-shuffle-api/repository/orm"
	"food-shuffle-api/utility/auth"
	"food-shuffle-api/utility/custom_error"
	"food-shuffle-api/utility/prefix"
	"net/http"
	"strings"

	"gorm.io/gorm"
)

type UserService struct{}

// ログイン処理を行う
func (userService *UserService) Login(bUser model.User) (res dto.LoginUser, err error) {
	// トランザクションを開始する
	err = orm.Transaction(func(tx *gorm.DB) error {
		// メールアドレスを元にユーザーが存在するかを確認する
		user, err := orm.GetUserByMailAddress(tx, bUser.MailAddress)
		if err != nil {
			logging.LogError("failed to get user", err)
			return custom_error.NewError(http.StatusNotFound, "User not found")
		}

		// パスワードが一致するか確認する
		err = bcrypto.CheckPasswordHash(user.Password, bUser.Password)
		if err != nil {
			logging.LogError("failed to check password", err)
			return custom_error.NewError(http.StatusUnauthorized, "Invalid password")
		}

		// メールアドレスとパスワードが一致した場合、jwtトークンを発行する
		res.JtiToken, err = auth.GenerateToken(tx, &user)
		if err != nil {
			logging.LogError("failed to generate token", err)
			return err
		}

		// エラーがなければnilを返し、トランザクションをコミットさせる
		return nil
	})

	return
}

// 店舗ごとのコースの一覧を取得する
func (service *UserService) GetCourses(restaurantUuid string) (res []dto.GetCourses, err error) {
	// トランザクションを開始する
	err = orm.Transaction(func(tx *gorm.DB) error {
		// レストランが存在することを確かめる
		err := orm.ExistsRestaurantByRestaurantUuid(tx, restaurantUuid)
		if err != nil {
			return err
		}

		//店舗UUIDに一致するコースを全件取得する
		courses, err := orm.GetCourses(tx, restaurantUuid)
		if err != nil {
			logging.LogError("failed to get courses", err)
			return err
		}

		// 構造体に格納していく
		for _, course := range courses {
			// 画像にプレフィックスをつける
			var prefixedImages []string
			for _, image := range course.Images {
				prefixedImages = append(prefixedImages, prefix.ImagePrefixCourse+image)
			}

			// レスポンスの構造体に格納する
			res = append(res, dto.GetCourses{
				CourseUuid:     course.CourseUuid,
				RestaurantUuid: course.RestaurantUuid,
				CourseName:     course.CourseName,
				Discription:    course.Description,
				Images:         prefixedImages,
				Price:          course.Price,
				Minimum:        course.Minimum,
			})
		}

		// エラーがなければnilを返し、トランザクションをコミットさせる
		return nil
	})

	return
}

// ユーザーがリクエストりした画像の閲覧権限のチェックと画像パスの生成
func (s *UserService) CheckImageAccessPermission(userUuid string, imageId string) (res string, err error) {
	// トランザクションの開始
	// FIXME: 何もチェックしてない
	err = orm.Transaction(func(tx *gorm.DB) error {
		//　ユーザーIDからユーザータイプを取得する
		userType, err := orm.GetUserType(tx, userUuid)
		if err != nil {
			fmt.Println(err)
			return err
		}

		// 一般ユーザーならば
		if userType == model.General {

		} else { // レストランユーザーならば

		}
		strings := strings.SplitAfter(imageId, "_")
		// プレフィックスから画像の種類を判断する
		switch strings[0] {
		case prefix.ImagePrefixCourse:
			res = "public/images/courses/"
		case prefix.ImagePrefixRestaurant:
			res = "public/images/restaurants/"
		case prefix.ImagePrefixReview:
			res = "public/images/reviews/"
		case prefix.ImagePrefixUserIcon:
			res = "public/images/icons/"
		}

		// 画像のUUIDをくっつける
		res = res + strings[1]

		return nil
	})

	return
}
