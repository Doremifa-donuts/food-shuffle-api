package handler

import (
	"errors"
	"fmt"
	"net/http"

	logging "food-shuffle-api/log"
	"food-shuffle-api/repository/model"
	"food-shuffle-api/utility/conversion"
	"food-shuffle-api/utility/custom_error"

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

// コース一覧の取得
func GetCoursesHandler(ctx *gin.Context) {
	// 対象のレストランUUIDを取得する
	uuid := ctx.Param("restaurant_uuid")
	if uuid == "" {
		logging.LogError("uuid not found", nil)
		// エラーレスポンスを返す
		conversion.ResponseJson(ctx, http.StatusBadRequest, nil)
		return
	}

	courses, err := UserService.GetCourses(uuid)
	if err != nil {
		logging.LogError("get courses failed", err)
		// エラーレスポンスを返す
		var customErr *custom_error.CustomError
		if errors.As(err, &customErr) {
			conversion.ResponseJson(ctx, customErr.StatusCode(), nil)
			return
		}
		// その他のエラーのレスポンス
		conversion.ResponseJson(ctx, http.StatusInternalServerError, nil)
		return
	}

	// 正常レスポンス
	conversion.ResponseJson(ctx, http.StatusOK, courses)
}

// 画像を取得するエンドポイント
func GetImagesHandler(ctx *gin.Context) {
	// ユーザーUUIDの取得
	userUuid, _ := ctx.Get("uuid")
	idAdjusted := userUuid.(string)

	// 画像IDの取得
	imageId := ctx.Param("image_id")
	if imageId == "" {
		logging.LogError("image id is not set", nil)
		conversion.ResponseJson(ctx, http.StatusBadRequest, nil)
		return
	}

	// 閲覧権限があるかどうかを確かめる処理
	path, err := UserService.CheckImageAccessPermission(idAdjusted, imageId)
	fmt.Println(path)

	if err != nil {
		// カスタムエラーの分類
		var customErr *custom_error.CustomError
		if errors.As(err, &customErr) {
			conversion.ResponseJson(ctx, customErr.StatusCode(), nil)
			return
		}
		conversion.ResponseJson(ctx, http.StatusInternalServerError, nil)
		return
	}
	ctx.File(path)
}
