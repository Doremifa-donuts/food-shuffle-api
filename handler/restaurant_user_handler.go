package handler

import (
	"net/http"

	logging "food-shuffle-api/log"
	"food-shuffle-api/model"
	"food-shuffle-api/service"
	"food-shuffle-api/utility/conversion"

	"github.com/gin-gonic/gin"
)

var RestaurantUserService = service.RestaurantUserService{}

// レストラン情報を取得する（busy_status確認用）
func GetRestaurantByUuidHandler(ctx *gin.Context) {
	uuid, ok := ctx.Get("uuid")
	if !ok {
		logging.LogError("uuid not found", nil)
		// エラーレスポンスを返す
		conversion.ResponseJson(ctx, http.StatusInternalServerError, nil)
		ctx.Abort()
		return
	}
	// レストラン情報を取得する
	resto, err := RestaurantUserService.GetRestaurantByUuid(uuid.(string))

	if err != nil {
		logging.LogError("get restaurant by uuid failed", err)
		// エラーレスポンスを返す
		conversion.ResponseJson(ctx, http.StatusInternalServerError, nil)
		ctx.Abort()
		return
	}
	// レストラン情報を返す
	conversion.ResponseJson(ctx, http.StatusOK, resto)
}

// 混雑状況を更新する
func UpdateBusyStatusHandler(ctx *gin.Context) {
	uuid, ok := ctx.Get("uuid")
	if !ok {
		logging.LogError("uuid not found", nil)
		// エラーレスポンスを返す
		conversion.ResponseJson(ctx, http.StatusInternalServerError, nil)
		ctx.Abort()
		return
	}

	busyStatusParam := ctx.Param("busy_status")          // 混雑状況をstring型でlinkから取得
	busyStatusInput := model.BusyStatus(busyStatusParam) // string型をBusyStatus型に変換

	message, err := RestaurantUserService.UpdateBusyStatus(uuid.(string), busyStatusInput)
	if err != nil {
		logging.LogError("update busy status failed", err)
		// エラーレスポンスを返す
		conversion.ResponseJson(ctx, http.StatusBadRequest, nil)
		ctx.Abort()
		return
	}

	// 混雑状況を返す
	conversion.ResponseJson(ctx, http.StatusOK, message)
}
