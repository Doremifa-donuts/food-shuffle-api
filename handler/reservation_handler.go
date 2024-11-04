package handler

import (
	"net/http"

	logging "food-shuffle-api/log"
	"food-shuffle-api/service"
	"food-shuffle-api/utility/conversion"
	"food-shuffle-api/utility/custom_error"

	"github.com/gin-gonic/gin"
)

// サービス層のメソッドは構造体と紐づいて管理されているため、処理を投げる構造体を呼び出す
var ReservationService = service.ReservationService{}

// 予約一覧を取得するハンドラー
func GetReservationsHandler(ctx *gin.Context) {
	uuid, ok := ctx.Get("uuid")
	if !ok {
		logging.LogError("uuid not found", custom_error.NewError(custom_error.ResourceNotFoundError))
		// エラーレスポンスを返す
		conversion.ResponseJson(ctx, http.StatusBadRequest, nil)
		ctx.Abort()
		return
	}

	// 予約一覧を取得する
	reviews, err := ReservationService.GetReservationsByRestaurant(uuid.(string))

	if err != nil {
		logging.LogError("get reservation failed", err)
		// エラーレスポンスを返す
		conversion.ResponseJson(ctx, http.StatusBadRequest, nil)
		ctx.Abort()
		return
	}

	// 予約一覧を返す
	conversion.ResponseJson(ctx, http.StatusOK, reviews)
}
