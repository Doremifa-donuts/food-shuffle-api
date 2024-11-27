package handler

import(
	"errors"
	"net/http"

	logging "food-shuffle-api/log"
	"github.com/gin-gonic/gin"
	"food-shuffle-api/model"
	"food-shuffle-api/utility/conversion"
	"food-shuffle-api/utility/custom_error"
	"food-shuffle-api/service"
)

var UrgentCampaignService = service.UrgentCampaignService{}

func CreateUrgentCampaignHandler(ctx *gin.Context) {

	if ctx.GetHeader("Content-Type") != "application/json" {
		logging.LogError("Content-Type is not application/json", nil)

		conversion.ResponseJson(ctx, http.StatusUnsupportedMediaType, nil)
		return
	}

	var urgentCampaign model.UrgentCampaign
	if err := ctx.ShouldBindJSON(&urgentCampaign); err != nil {
		logging.LogError("Error binding JSON:", err)
		conversion.ResponseJson(ctx, http.StatusBadRequest, nil)
		return
	}

	restaurantUuid, _ := ctx.Get("uuid")
	urgentCampaign.RestaurantUuid = restaurantUuid.(string)

	result, err := UrgentCampaignService.UrgentCampaignRegister(urgentCampaign)
	if err != nil {
		logging.LogError("Error logging in:", err)
		var customErr *custom_error.CustomError
		if errors.As(err, &customErr) {
			conversion.ResponseJson(ctx, customErr.StatusCode(), nil)
			return
		}
		conversion.ResponseJson(ctx, http.StatusInternalServerError, nil)
		return
	}

	conversion.ResponseJson(ctx, http.StatusOK, result)
}

func GetUrgentCampaignHandler(ctx *gin.Context) {

	uuid := ctx.Param("campaignUuid")
	if uuid == "" {
		logging.LogError("uuid not found", nil)
		// エラーレスポンスを返す
		conversion.ResponseJson(ctx, http.StatusBadRequest, nil)
		ctx.Abort()
		return
	}

	campaign, err := UrgentCampaignService.GetUrgentCampaign(uuid)

	if err != nil {
		logging.LogError("get urgent campaign failed", err)
		// エラーレスポンスを返す
		conversion.ResponseJson(ctx, http.StatusInternalServerError, nil)
		ctx.Abort()
		return
	}

	conversion.ResponseJson(ctx, http.StatusOK, campaign)
}