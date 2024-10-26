package service

import (
	bcrypt "food-shuffle-api/crypto"
	"food-shuffle-api/model"
	"food-shuffle-api/repository"
	"food-shuffle-api/utility/auth"
	"food-shuffle-api/utility/custom_error"
)

type UserService struct {}

// ログイン処理を行う
func (userService *UserService) Login(user model.User) (string, error) {
	// トランザクションを開始する
	db := repository.GetDB().Begin()

	// dbから取得したパスワードと比較するために、パスワードを保持しておく
	inputPassword := user.Password

	// メールアドレスを元にユーザーが存在するかを確認する
	user, err := repository.GetUserByMailAddress(db, user.MailAddress)
	if err != nil {
		return "", custom_error.NewError(custom_error.ResourceNotFoundError)
	}

	// パスワードが一致するか確認する
	err = bcrypt.CheckPasswordHash(user.Password, inputPassword)
	if err != nil {
		return "", custom_error.NewError(custom_error.UnauthorizedError)
	}

	// メールアドレスとパスワードが一致した場合、jwtトークンを発行する
	tokenString, err := auth.GenerateToken(&user)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}