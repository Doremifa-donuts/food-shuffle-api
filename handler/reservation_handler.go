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

var ReservationService = service.ReservationService{}

func ReservationRegisterHandler(ctx *gin.Context) {
	// リクエストをバインドする
	var reservation model.Reservation
	customErr := conversion.BindJSON(ctx, &reservation)
	if customErr != nil {
		conversion.ResponseJson(ctx, customErr.StatusCode(), nil)
		return
	}
	// パスパラメータに含まれるrestaurant uuidを取得
	restaurantUuid := ctx.Param("restaurant_uuid")
	if restaurantUuid == "" {
		conversion.ResponseJson(ctx, http.StatusBadRequest, nil)
		return
	}
	reservation.RestaurantUuid = restaurantUuid

	// ミドルウェアが解析したuuidを構造体に格納
	uuid, _ := ctx.Get("uuid")
	reservation.UserUuid = uuid.(string)

	//reservation_register_serviseへ処理を投げる
	result, err := ReservationService.ResevationRegister(reservation)
	if err != nil {
		// エラーログを書き込む
		logging.LogError("Error logging in:", err)
		// カスタムエラーのエラー型を宣言
		var customErr *custom_error.CustomError
		if errors.As(err, &customErr) {
			conversion.ResponseJson(ctx, customErr.StatusCode(), nil)
			return
		}

		// カスタムエラー以外のレスポンス
		conversion.ResponseJson(ctx, http.StatusInternalServerError, nil)
		return
	}

	//成功レスポンス
	conversion.ResponseJson(ctx, http.StatusOK, result)
}

// 予約一覧を取得するハンドラー
func GetReservationsHandler(ctx *gin.Context) {
	uuid, ok := ctx.Get("uuid")
	if !ok {
		logging.LogError("uuid not found", nil)
		// エラーレスポンスを返す
		conversion.ResponseJson(ctx, http.StatusBadRequest, nil)
		return
	}

	// 予約一覧を取得する
	reviews, err := ReservationService.GetReservationsByRestaurant(uuid.(string))

	if err != nil {
		logging.LogError("get reservation failed", err)
		// エラーレスポンスを返す
		conversion.ResponseJson(ctx, http.StatusInternalServerError, nil)
		ctx.Abort()
		return
	}

	// 予約一覧を返す
	conversion.ResponseJson(ctx, http.StatusOK, reviews)
}
