package handler

import (
	"net/http"

	logging "food-shuffle-api/log"
	"food-shuffle-api/service"
	"food-shuffle-api/utility/conversion"

	"github.com/gin-gonic/gin"
)

// サービス層のメソッドは構造体と紐づいて管理されているため、処理を投げる構造体を呼び出す
var ReservationService = service.ReservationService{}

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

// 予約を取得するハンドラー
func GetReservationDetailHandler(ctx *gin.Context) {
	uuid, ok := ctx.Get("uuid")
	if !ok {
		logging.LogError("uuid not found", nil)
		// エラーレスポンスを返す
		conversion.ResponseJson(ctx, http.StatusBadRequest, nil)
		ctx.Abort()
		return
	}
	// 予約UUIDを取得
	reservation_uuid := ctx.Param("reservation_uuid")

	// 予約を取得する
	review, err := ReservationService.GetReservationDetailByReservation(uuid.(string), reservation_uuid)
	if err != nil {
		logging.LogError("get reservation failed", err)
		// エラーレスポンスを返す
		conversion.ResponseJson(ctx, http.StatusInternalServerError, nil)
		ctx.Abort()
		return
	}

	// 予約を返す
	conversion.ResponseJson(ctx, http.StatusOK, review)
}

// 予約を承認するハンドラー
// next time, make uuid as array or slice or whatever that is
func ApproveReservationHandler(ctx *gin.Context) {
	uuid, ok := ctx.Get("uuid")
	if !ok {
		logging.LogError("uuid not found", nil)
		// エラーレスポンスを返す
		conversion.ResponseJson(ctx, http.StatusBadRequest, nil)
		ctx.Abort()
		return
	}
	// 予約UUIDを取得
	reservation_uuid := ctx.Param("reservation_uuid")

	// 予約を承認する
	err := ReservationService.ApproveReservation(uuid.(string), reservation_uuid)
	if err != nil {
		logging.LogError("approve reservation failed", err)
		ctx.Abort()
		return
	}
	// 予約を返す
	conversion.ResponseJson(ctx, http.StatusOK, nil)
}
