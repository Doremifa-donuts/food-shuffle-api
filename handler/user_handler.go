package handler

import (
	"errors"
	"net/http"

	logging "food-shuffle-api/log"
	"food-shuffle-api/model"
	"food-shuffle-api/utility/custom_error"
	"food-shuffle-api/utility/enbox"

	"github.com/gin-gonic/gin"

	"food-shuffle-api/service"
)

// サービス層のメソッドは構造体と紐づいて管理されているため、処理を投げる構造体を呼び出す
var UserService = service.UserService{}

// ログイン処理
func LoginHandler(c *gin.Context) {
	// リクエストボディを取得する

	// 取得したパラメータを格納する構造体を宣言
	var user model.User
	// リクエストボディを構造体にバインドする
	if err := c.ShouldBindJSON(&user); err != nil {
		// エラーログを書き込む
		logging.LogError("Error binding JSON:", err)

		// エラーレスポンスを返す
		enbox.ResponseJson(c, http.StatusBadRequest, nil)
		return
	}

		// ログイン処理の流れはサービス層で行う
	tokenString, err :=UserService.Login(user)
	if err != nil {
		// エラーログを書き込む
		logging.LogError("Error logging in:", err)

		// エラーによって異なるレスポンスを返す
		// カスタムエラーの変数を宣言
		var customErr *custom_error.CustomError

		// カスタムエラーにキャスト可能か確認する
		// キャスト可能な場合、自身で定義したビジネスロジック上のエラーなので、適切なレスポンスを返す
		if errors.As(err, &customErr) {
			switch customErr.Code() {
				case custom_error.ResourceNotFoundError:	// メールアドレスが違っていた場合
					enbox.ResponseJson(c, http.StatusUnauthorized, nil)
					return
				case custom_error.UnauthorizedError:		// パスワードが一致しなかった場合
					enbox.ResponseJson(c, http.StatusUnauthorized, nil)
					return
				default:	// 設定したカスタムエラー処理に抜けがある場合の処理
					enbox.ResponseJson(c, http.StatusBadRequest, nil)
					return
			}
		} else {
			enbox.ResponseJson(c, http.StatusInternalServerError, nil)
			return
		}
	}

	// 正常に終了した場合のレスポンス
	enbox.ResponseJson(c, http.StatusOK, tokenString)

}