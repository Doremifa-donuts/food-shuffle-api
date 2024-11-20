package service

import (
	"food-shuffle-api/bcrypto"
	"food-shuffle-api/dto"
	logging "food-shuffle-api/log"
	"food-shuffle-api/model"
	"food-shuffle-api/repository"
	"food-shuffle-api/utility/auth"
	"food-shuffle-api/utility/custom_error"
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
