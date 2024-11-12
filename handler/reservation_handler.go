package handler

import (
	"errors"
	"net/http"

	logging "food-shuffle-api/log"
	"food-shuffle-api/model"
	"food-shuffle-api/utility/conversion"
	"food-shuffle-api/utility/custom_error"

	"food-shuffle-api/service"

	"github.com/gin-gonic/gin"
)

var ReservationService = service.ReservationService{}

func ReservationRegistorHandler(ctx *gin.Context) {
	// ヘッダーのContent-Typeにapplication/jsonが含まれているか確認
	if ctx.GetHeader("Content-Type") != "application/json" {
		logging.LogError("Content-Type is not application/json", nil)

		// エラーレスポンスを返す
		conversion.ResponseJson(ctx, http.StatusUnsupportedMediaType, nil)
		return
	}

	// リクエストをバインドする
	var reservation model.Reservation
	if err := ctx.ShouldBindJSON(&reservation); err != nil {
		// エラーログを書き込む
		logging.LogError("Error binding JSON:", err)
		conversion.ResponseJson(ctx, http.StatusBadRequest, nil)
		return
	}

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

		// エラーが発生しなかった時のレスポンス
		conversion.ResponseJson(ctx, http.StatusInternalServerError, nil)
		return
	}

	//成功レスポンス
	conversion.ResponseJson(ctx, http.StatusOK, gin.H{"ReservationUuid": result})
}

// 予約一覧を取得するハンドラー
func GetReservationsHandler(ctx *gin.Context) {
	uuid, ok := ctx.Get("uuid")
	if !ok {
		logging.LogError("uuid not found", nil)
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
		conversion.ResponseJson(ctx, http.StatusInternalServerError, nil)
		ctx.Abort()
		return
	}

	// 予約一覧を返す
	conversion.ResponseJson(ctx, http.StatusOK, reviews)
}
