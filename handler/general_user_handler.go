package handler

import (
	"errors"
	"fmt"
	"net/http"

	logging "food-shuffle-api/log"
	"food-shuffle-api/model"
	"food-shuffle-api/utility/conversion"
	"food-shuffle-api/utility/custom_error"

	"food-shuffle-api/service"

	"github.com/gin-gonic/gin"
)

// サービス層のメソッドは構造体と紐づいて管理されているため、処理を投げる構造体を呼び出す
var GeneralUserService = service.GeneralUserService{}

// 一般ユーザーのアカウントを作成する
func GeneralUserRegisterHandler(ctx *gin.Context) {
	// ヘッダーのContent-Typeにapplication/jsonが含まれているか確認
	if ctx.GetHeader("Content-Type") != "application/json" {
		logging.LogError("Content-Type is not application/json", nil)

		// エラーレスポンスを返す
		conversion.ResponseJson(ctx, http.StatusUnsupportedMediaType, nil)
		return
	}

	// リクエストをバインドする
	var user model.User
	if err := ctx.ShouldBindBodyWithJSON(&user); err != nil {
		// リクエストのバインドに失敗した場合は、400レスポンスを返す
		logging.LogError("Error binding JSON user:", err)
		fmt.Println(err.Error())
		conversion.ResponseJson(ctx, http.StatusBadRequest, nil)
		return
	}

	// 一般ユーザーの追加テーブルにもバインドする
	var generalUser model.GeneralUser
	if err := ctx.ShouldBindBodyWithJSON(&generalUser); err != nil {
		// リクエストのバインドに失敗した場合は、400レスポンスを返す
		logging.LogError("Error binding JSON general user:", err)
		conversion.ResponseJson(ctx, http.StatusBadRequest, nil)
		return
	}

	// サービス層に処理を投げる
	result, err := GeneralUserService.Register(user, generalUser)
	if err != nil {
		// エラーログを書き込む
		logging.LogError("Error registering general user:", err)

		// エラーハンドリング
		var customErr *custom_error.CustomError
		if errors.As(err, &customErr) { // カスタムエラーの場合
			conversion.ResponseJson(ctx, customErr.StatusCode(), nil)
			return
		} else { // TODO: カスタムエラー以外の場合のハンドリングを行う
			conversion.ResponseJson(ctx, http.StatusInternalServerError, nil)
			return
		}

	}

	// 成功レスポンス
	conversion.ResponseJson(ctx, http.StatusOK, gin.H{"Token": result})
}
