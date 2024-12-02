package handler

import (
	"errors"
	"net/http"

	logging "food-shuffle-api/log"
	"food-shuffle-api/model"
	"food-shuffle-api/utility/conversion"
	"food-shuffle-api/utility/custom_error"
	"food-shuffle-api/utility/prefix"

	"food-shuffle-api/service"

	"github.com/gin-gonic/gin"
)

// サービス層のメソッドは構造体と紐づいて管理されているため、処理を投げる構造体を呼び出す
var GeneralUserService = service.GeneralUserService{}

// 一般ユーザーのアカウントを作成する
func GeneralUserRegisterHandler(ctx *gin.Context) {
	// リクエストをバインドする
	var user model.User
	// 一般ユーザーの追加テーブルにもバインドする
	var generalUser model.GeneralUser

	customErr := conversion.BindJSON(ctx, &user, &generalUser)
	if customErr != nil {
		logging.LogError("failed bind json", customErr)
		conversion.ResponseJson(ctx, customErr.StatusCode(), nil)
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
	conversion.ResponseJson(ctx, http.StatusOK, result)
}

func GetRestaurantDetailHandler(ctx *gin.Context) {
	uuid := ctx.Param("restaurantUuid")
	if uuid == "" {
		logging.LogError("uuid not found", nil)
		// エラーレスポンスを返す
		conversion.ResponseJson(ctx, http.StatusBadRequest, nil)
		ctx.Abort()
		return
	}

	detail, err := GeneralUserService.GetRestaurantDetail(uuid)

	if err != nil {
		logging.LogError("get restaurant detail failed", err)
		// エラーレスポンスを返す
		conversion.ResponseJson(ctx, http.StatusInternalServerError, nil)
		ctx.Abort()
		return
	}

	//画像があればプレフィックスを付ける
	if len(detail.Images) > 0 {
		// 画像のプレフィックス処理
		prefixedImages := make([]string, len(detail.Images))
		for i, image := range detail.Images {
			if image == "" {
				//画像の文字列が空、もしくは予期しないエラーが発生した場合
				logging.LogError("image not found or unexpected error", nil)
				conversion.ResponseJson(ctx, http.StatusInternalServerError, nil)
				ctx.Abort()
				return
			}

			prefixedImages[i] = prefix.ImagePrefixRestaurant + image
		}
		detail.Images = prefixedImages
	}

	conversion.ResponseJson(ctx, http.StatusOK, detail)
}
