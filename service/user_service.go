package service

import (
	"food-shuffle-api/bcrypto"
	"food-shuffle-api/model"
	"food-shuffle-api/repository"
	"food-shuffle-api/utility/auth"
	"food-shuffle-api/utility/custom_error"

	"gorm.io/gorm"
)

type UserService struct{}

// ログイン処理を行う
func (userService *UserService) Login(user model.User) (string, error) {

	// dbから取得したパスワードと比較するために、パスワードを保持しておく
	inputPassword := user.Password

	// レスポンスのを初期化する
	var tokenString string

	// トランザクションを開始する
	err := repository.Transaction(func(tx *gorm.DB) error {
		// メールアドレスを元にユーザーが存在するかを確認する
		user, err := repository.GetUserByMailAddress(tx, user.MailAddress)
		if err != nil {
			return custom_error.NewError(custom_error.ResourceNotFoundError)
		}

		// パスワードが一致するか確認する
		err = bcrypto.CheckPasswordHash(user.Password, inputPassword)
		if err != nil {
			return custom_error.NewError(custom_error.UnauthorizedError)
		}

		// メールアドレスとパスワードが一致した場合、jwtトークンを発行する
		tokenString, err = auth.GenerateToken(tx, &user)
		if err != nil {
			return err
		}

		// エラーがなければnilを返し、トランザクションをコミットさせる
		return nil
	})

	// 結果を返す
	// トランザクションが失敗した場合はエラーを返す
	if err != nil {
		return "", err
	}

	// 正常の返り値
	return tokenString, nil
}
