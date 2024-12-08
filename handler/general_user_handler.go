package handler

import (
	"errors"
	"fmt"
	"net/http"

	"food-shuffle-api/dto"
	logging "food-shuffle-api/log"
	"food-shuffle-api/model"
	"food-shuffle-api/service"
	"food-shuffle-api/utility/conversion"
	"food-shuffle-api/utility/custom_error"
	"food-shuffle-api/utility/prefix"

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

// レストランの詳細情報を取得
func GetRestaurantDetailHandler(ctx *gin.Context) {
	// ユーザーのUUIDを取得	//TODO: ユーザーがレストランの情報を見る権限があるかを確認する必要がある
	// userUuid, _ := ctx.Get("uuid")
	// idAdjusted := userUuid.(string)

	// レストランのUUIDを取得
	restaurantUuid := ctx.Param("restaurant_uuid")
	if restaurantUuid == "" {
		logging.LogError("uuid not found", nil)
		// エラーレスポンスを返す
		conversion.ResponseJson(ctx, http.StatusBadRequest, nil)
		return
	}

	detail, err := GeneralUserService.GetRestaurantDetail(restaurantUuid)
	if err != nil {
		logging.LogError("get restaurant detail failed", err)
		// エラーレスポンスを返す
		conversion.ResponseJson(ctx, http.StatusInternalServerError, nil)
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

// 店舗へのチェックインを行う
func PostCheckInRestaurantHandler(ctx *gin.Context) {
	// ユーザーUUIDを取得
	userUuid, _ := ctx.Get("uuid")

	// パスパラメータから店舗UUIDを取得
	restaurantUuid, ok := ctx.Params.Get("restaurant_uuid")
	if !ok {
		conversion.ResponseJson(ctx, http.StatusBadRequest, nil)
		return
	}

	// 位置情報を取得
	var latlong dto.CheckInLocation
	customErr := conversion.BindJSON(ctx, &latlong)
	if customErr != nil {
		conversion.ResponseJson(ctx, customErr.StatusCode(), nil)
		return
	}

	// サービス層に処理を投げる
	err := GeneralUserService.PostCheckInRestaurant(userUuid.(string), restaurantUuid, latlong)
	if err != nil {
		var customErr *custom_error.CustomError
		if errors.As(err, &customErr) {
			conversion.ResponseJson(ctx, customErr.StatusCode(), nil)
			return
		}
		// 切り分けできてないエラー
		conversion.ResponseJson(ctx, http.StatusInternalServerError, nil)
		return
	}

	// 成功レスポンス
	conversion.ResponseJson(ctx, http.StatusOK, nil)
}

// ユーザーの通知モードを変更
func PutShareStatusHandler(ctx *gin.Context) {
	//リクエストを構造体にバインド
	var generalUser model.GeneralUser

	//ユーザーIDを取得する
	uuid, _ := ctx.Get("uuid")
	generalUser.UserUuid = uuid.(string)
	fmt.Println("uuid", generalUser.UserUuid)

	// 変更後のモードを取得
	Status := ctx.Param("status")
	switch Status {
	case "Active":
		generalUser.ShareStatus = model.Active
	case "Silent":
		generalUser.ShareStatus = model.Silent
	case "Disabled":
		generalUser.ShareStatus = model.Disabled
	default:
		logging.LogError("status not found", nil)
		// エラーレスポンスを返す
		conversion.ResponseJson(ctx, http.StatusBadRequest, nil)
		return
	}

	err := GeneralUserService.PutShareStatus(generalUser)
	if err != nil {
		var customErr *custom_error.CustomError
		if errors.As(err, &customErr) { // カスタムエラーの場合
			conversion.ResponseJson(ctx, customErr.StatusCode(), nil)
		} else {
			conversion.ResponseJson(ctx, http.StatusInternalServerError, nil)
		}
		return
	}
	// 成功レスポンス
	conversion.ResponseJson(ctx, http.StatusOK, nil)
}

// レビューを書いた店のリスト
func GetReviewedRestaurantsHandler(ctx *gin.Context) {
	// ユーザーUUIDを取得
	uuid, _ := ctx.Get("uuid")
	uuidAdjusted := uuid.(string)

	// サービス層に処理を投げる
	res, err := GeneralUserService.GetIsReviewedRestaurants(true, uuidAdjusted)
	if err != nil {
		//カスタムエラーの場合
		var customErr *custom_error.CustomError
		if errors.As(err, &customErr) {
			conversion.ResponseJson(ctx, customErr.StatusCode(), nil)
			return
		}
		// その他のエラー
		conversion.ResponseJson(ctx, http.StatusInternalServerError, nil)
		return
	}

	// 成功レスポンス
	conversion.ResponseJson(ctx, http.StatusOK, res)
}

// 訪れただけの店
func GetVisitedRestaurantsHandler(ctx *gin.Context) {
	// ユーザーUUIDを取得
	uuid, _ := ctx.Get("uuid")
	uuidAdjusted := uuid.(string)

	// サービス層に処理を投げる
	res, err := GeneralUserService.GetIsReviewedRestaurants(false, uuidAdjusted)
	if err != nil {
		//カスタムエラーの場合
		var customErr *custom_error.CustomError
		if errors.As(err, &customErr) {
			conversion.ResponseJson(ctx, customErr.StatusCode(), nil)
			return
		}
		// その他のエラー
		conversion.ResponseJson(ctx, http.StatusInternalServerError, nil)
		return
	}

	// 成功レスポンス
	conversion.ResponseJson(ctx, http.StatusOK, res)
}
