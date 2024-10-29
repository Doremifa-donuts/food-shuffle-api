package service

import (
	"errors"
	"food-shuffle-api/bcrypto"
	logging "food-shuffle-api/log"
	"food-shuffle-api/model"
	"food-shuffle-api/repository"
	"food-shuffle-api/utility/auth"
	"food-shuffle-api/utility/custom_error"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GeneralUserService struct{}

// 一般ユーザーのアカウントを作成し、トークンを返す
func (service *GeneralUserService) Register(user model.User, generalUser model.GeneralUser) (string, error) {

	// レスポンスの型を初期化する
	var tokenString string

	// トランザクションを開始する
	err := repository.Transaction(func(tx *gorm.DB) error {

		// メールアドレスを元にユーザーが存在するかを確認する
		_, err := repository.GetUserByMailAddress(tx, user.MailAddress)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) { // 存在しない以外のエラーがある場合
			logging.LogError("failed to get user", err)
			return err
		}
		if err == nil { // メールアドレスが一致するユーザーが存在する場合
			logging.LogError("users mail address already exists", err)
			return custom_error.NewError(custom_error.ConflictError)
		}

		// UUIDを生成する
		uuid, err := uuid.NewRandom()
		if err != nil {
			logging.LogError("failed to generate user uuid", err)
			return err
		}

		// それぞれのテーブルにuuidを挿入する
		user.UserUuid = uuid.String()
		generalUser.UserUuid = uuid.String()

		// パスワードをハッシュ化する
		hashedPassword, err := bcrypto.GetHashPassword(user.Password)
		if err != nil {
			logging.LogError("failed to hash password", err)
			return err
		}
		// ハッシュ化したパスワードに入れ替える
		user.Password = hashedPassword

		// ユーザータイプを一般に設定する
		user.UserType = model.General

		// 挿入するデータが完成したのでここから挿入していく
		// ユーザーテーブルに一般ユーザーを追加する
		err = repository.CreateUser(tx, user)
		if err != nil {
			return err
		}

		// 一般ユーザーテーブルに追加情報を追加する
		err = repository.CreateGeneralUser(tx, generalUser)
		if err != nil {
			return err
		}

		// トークンを発行する
		tokenString, err = auth.GenerateToken(tx, &user)
		if err != nil {
			return err
		}

		return nil

	})

	// ユーザーテーブルに共通事項を追加する
	return tokenString, err
}
