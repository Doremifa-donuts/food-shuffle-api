package handler

import (
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
		logging.LogError("Content-Type is not application/json", custom_error.NewError(custom_error.InvalidDataError))

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

	ctx.Get("uuid")

	//reservation_register_serviseへ処理を投げる
	result, err := ReservationService.ResevationRegister(reservation)
	if err != nil {
		// エラーログを書き込む
		logging.LogError("Error logging in:", err)
	}
	//成功レスポンス
	conversion.ResponseJson(ctx, http.StatusOK, gin.H{"Token": result})
}
