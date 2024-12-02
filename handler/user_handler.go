package handler

import (
	"errors"
	"net/http"

	logging "food-shuffle-api/log"
	"food-shuffle-api/model"
	"food-shuffle-api/utility/conversion"
	"food-shuffle-api/utility/custom_error"
	"food-shuffle-api/utility/prefix"

	"github.com/gin-gonic/gin"

	"food-shuffle-api/service"
)

// サービス層のメソッドは構造体と紐づいて管理されているため、処理を投げる構造体を呼び出す
var UserService = service.UserService{}

// ログイン処理
func LoginHandler(ctx *gin.Context) {

	// 取得したパラメータを格納する構造体を宣言
	var user model.User
	customErr := conversion.BindJSON(ctx, &user)
	if customErr != nil {
		logging.LogError("failed bind json", customErr)
		conversion.ResponseJson(ctx, customErr.StatusCode(), nil)
		return
	}

	// ログイン処理の流れはサービス層で行う
	result, err := UserService.Login(user)
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
	conversion.ResponseJson(ctx, http.StatusOK, result) // レスポンスにトークンを返す(tokenString)
}

func GetCoursesHandler(ctx *gin.Context) {

	uuid := ctx.Param("restaurantUuid")
	if uuid == "" {
		logging.LogError("uuid not found", nil)
		// エラーレスポンスを返す
		conversion.ResponseJson(ctx, http.StatusBadRequest, nil)
		ctx.Abort()
		return
	}

	courses, err := UserService.GetCourses(uuid)
	if err != nil {
		logging.LogError("get courses failed", err)
		// エラーレスポンスを返す
		conversion.ResponseJson(ctx, http.StatusInternalServerError, nil)
		ctx.Abort()
		return
	}

	//画像があればプレフィックスを付ける
	for i := range courses {
		if len(courses[i].Images) > 0 {
			// 画像のプレフィックス処理
			prefixedImages := make([]string, len(courses[i].Images))
			for j, image := range courses[i].Images {
				if image == "" {
					//画像の文字列が空、もしくは予期しないエラーが発生した場合
					logging.LogError("image not found or unexpected error", nil)
					conversion.ResponseJson(ctx, http.StatusInternalServerError, nil)
					ctx.Abort()
					return
				}

				prefixedImages[j] = prefix.ImagePrefixCourse + image
			}
			courses[i].Images = prefixedImages
		}
	}
	conversion.ResponseJson(ctx, http.StatusOK, courses)
}
