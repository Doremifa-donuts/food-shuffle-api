package handler

import (
	"errors"
	"net/http"

	logging "food-shuffle-api/log"
	"food-shuffle-api/repository/model"
	"food-shuffle-api/service"
	"food-shuffle-api/utility/conversion"
	"food-shuffle-api/utility/custom_error"

	"github.com/gin-gonic/gin"
)

var UrgentCampaignService = service.UrgentCampaignService{}

func CreateUrgentCampaignHandler(ctx *gin.Context) {
	// リクエストをバインドする
	var urgentCampaign model.UrgentCampaign
	customErr := conversion.BindJSON(ctx, &urgentCampaign)
	if customErr != nil {
		conversion.ResponseJson(ctx, customErr.StatusCode(), nil)
		return
	}

	// アクセスした人を特定する
	restaurantUuid, _ := ctx.Get("uuid")
	urgentCampaign.RestaurantUuid = restaurantUuid.(string)

	// サービス層で処理を行う
	result, err := UrgentCampaignService.UrgentCampaignRegister(urgentCampaign)
	if err != nil {
		logging.LogError("Error logging in:", err)
		var customErr *custom_error.CustomError
		if errors.As(err, &customErr) {
			conversion.ResponseJson(ctx, customErr.StatusCode(), nil)
			return
		}
		// カスタムエラー以外のエラーレスポンス
		conversion.ResponseJson(ctx, http.StatusInternalServerError, nil)
		return
	}

	// 正常レスポンス
	conversion.ResponseJson(ctx, http.StatusOK, result)
}

// お助けブーストの1件取得
func GetUrgentCampaignHandler(ctx *gin.Context) {
	uuid := ctx.Param("campaign_uuid")
	if uuid == "" {
		logging.LogError("uuid not found", nil)
		// エラーレスポンスを返す
		conversion.ResponseJson(ctx, http.StatusBadRequest, nil)
		return
	}

	// サービス層に処理を投げる
	campaign, err := UrgentCampaignService.GetUrgentCampaign(uuid)
	if err != nil {
		logging.LogError("get urgent campaign failed", err)
		// エラーレスポンスを返す
		var customErr *custom_error.CustomError
		if errors.As(err, &customErr) {
			conversion.ResponseJson(ctx, customErr.StatusCode(), nil)
			return
		}
		//　カスタムエラー以外のエラーレスポンス
		conversion.ResponseJson(ctx, http.StatusInternalServerError, nil)
		return
	}

	// 正常レスポンス
	conversion.ResponseJson(ctx, http.StatusOK, campaign)
}
