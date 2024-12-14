package service

import (
	"food-shuffle-api/bcrypto"
	"food-shuffle-api/dto"
	logging "food-shuffle-api/log"
	"food-shuffle-api/model"
	"food-shuffle-api/repository"
	"food-shuffle-api/utility/auth"
	"food-shuffle-api/utility/custom_error"
	"food-shuffle-api/utility/prefix"
	"net/http"

	"gorm.io/gorm"
)

type UserService struct{}

// ログイン処理を行う
func (userService *UserService) Login(bUser model.User) (res dto.LoginUser, err error) {
	// トランザクションを開始する
	err = repository.Transaction(func(tx *gorm.DB) error {
		// メールアドレスを元にユーザーが存在するかを確認する
		user, err := repository.GetUserByMailAddress(tx, bUser.MailAddress)
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
	err = repository.Transaction(func(tx *gorm.DB) error {
		// レストランが存在することを確かめる
		err := repository.ExistsRestaurantByRestaurantUuid(tx, restaurantUuid)
		if err != nil {
			return err
		}

		//店舗UUIDに一致するコースを全件取得する
		courses, err := repository.GetCourses(tx, restaurantUuid)
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
