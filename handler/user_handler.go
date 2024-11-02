package handler

import (
	"errors"
	"net/http"

	logging "food-shuffle-api/log"
	"food-shuffle-api/model"
	"food-shuffle-api/utility/conversion"
	"food-shuffle-api/utility/custom_error"

	"github.com/gin-gonic/gin"

	"food-shuffle-api/service"
)

// サービス層のメソッドは構造体と紐づいて管理されているため、処理を投げる構造体を呼び出す
var UserService = service.UserService{}

// ログイン処理
func LoginHandler(ctx *gin.Context) {
	// ヘッダーのContent-Typeにapplication/jsonが含まれているか確認
	if ctx.GetHeader("Content-Type") != "application/json" {
		err := custom_error.NewError(http.StatusBadRequest, "Content-Type is not application/json")
		logging.LogError("Content-Type is not application/json", err)

		// エラーレスポンスを返す
		conversion.ResponseJson(ctx, http.StatusUnsupportedMediaType, nil)
		return
	}

	// 取得したパラメータを格納する構造体を宣言
	var user model.User
	// リクエストボディを構造体にバインドする
	if err := ctx.ShouldBindJSON(&user); err != nil {
		// エラーログを書き込む
		logging.LogError("Error binding JSON:", err)

		// エラーレスポンスを返す
		conversion.ResponseJson(ctx, http.StatusBadRequest, nil)
		return
	}

	// ログイン処理の流れはサービス層で行う
	tokenString, err := UserService.Login(user)
	if err != nil {
		// エラーログを書き込む
		logging.LogError("Error logging in:", err)

		// カスタムエラーにキャスト可能か確認する
		// カスタムエラーの変数を宣言
		var customErr *custom_error.CustomError

		if errors.As(err, &customErr) {
			conversion.ResponseJson(ctx, customErr.StatusCode(), nil)
			return
		} else { // TODO: カスタムエラー以外のエラーを分類する
			conversion.ResponseJson(ctx, http.StatusInternalServerError, nil)
			return
		}
	}

	// 正常に終了した場合のレスポンス
	conversion.ResponseJson(ctx, http.StatusOK, gin.H{"Token": tokenString}) // レスポンスにトークンを返す(tokenString)

}
