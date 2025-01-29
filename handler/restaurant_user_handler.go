package handler

import (
	"errors"
	"net/http"

	logging "food-shuffle-api/log"
	"food-shuffle-api/repository/model"
	"food-shuffle-api/utility/conversion"
	"food-shuffle-api/utility/custom_error"

	"food-shuffle-api/service"

	"github.com/gin-gonic/gin"
)

var RestaurantUserService = service.RestaurantUserService{}

// 混雑状況の切り替え
func PutBusyStatusHandler(ctx *gin.Context) {
	//リクエストを構造体にバインド
	var restaurantUser model.RestaurantUser

	//ユーザーIDを取得
	uuid, _ := ctx.Get("uuid")
	restaurantUser.RestaurantUuid = uuid.(string)

	//変更後の混雑状況を取得
	Status := ctx.Param("status")
	switch Status {
	case "Free":
		restaurantUser.BusyStatus = model.Free
	case "Spare":
		restaurantUser.BusyStatus = model.Spare
	case "Packed":
		restaurantUser.BusyStatus = model.Packed
	default:
		logging.LogError("status not found", nil)
		// エラーレスポンスを返す
		conversion.ResponseJson(ctx, http.StatusBadRequest, nil)
		return
	}

	err := RestaurantUserService.PutBusyStatus(restaurantUser)
	if err != nil {
		var customErr *custom_error.CustomError
		if errors.As(err, &customErr) { // カスタムエラーの場合
			// エラーレスポンスを返す
			conversion.ResponseJson(ctx, customErr.StatusCode(), nil)
		} else {
			conversion.ResponseJson(ctx, http.StatusInternalServerError, nil)
		}
		return
	}
	// レスポンスを返す
	conversion.ResponseJson(ctx, http.StatusOK, nil)

}

// 自身の店舗情報を取得する
func GetOwnRestaurantDetailHandler(ctx *gin.Context) {
	// 自身のUUIDを取得
	restaurantUuid, _ := ctx.Get("uuid")
	idAdjusted := restaurantUuid.(string)

	detail, err := GeneralUserService.GetRestaurantDetail(idAdjusted)
	if err != nil {
		logging.LogError("get restaurant detail failed", err)
		var customErr *custom_error.CustomError
		if errors.As(err, &customErr) {
			conversion.ResponseJson(ctx, customErr.StatusCode(), nil)
			return
		}
		// エラーレスポンスを返す
		conversion.ResponseJson(ctx, http.StatusInternalServerError, nil)
		return
	}

	// 成功レスポンス
	conversion.ResponseJson(ctx, http.StatusOK, detail)
}
