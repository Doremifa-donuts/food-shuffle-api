package service

import (
	"errors"
	"food-shuffle-api/dto"
	logging "food-shuffle-api/log"
	"food-shuffle-api/repository/model"
	"food-shuffle-api/repository/orm"
	"food-shuffle-api/utility/custom_error"
	"net/http"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UrgentCampaignService struct{}

func (service *UrgentCampaignService) UrgentCampaignRegister(urgentCampaign model.UrgentCampaign) (res dto.CreateUrgentCampaign, err error) {
	// トランザクションを開始する
	err = orm.Transaction(func(tx *gorm.DB) error {
		// キャンペーンUUIDを生成する
		CampaignUuid, err := uuid.NewV7()
		if err != nil {
			logging.LogError("failed to generate CampaignUuid", err)
			return err
		}
		res.CampaignUuid = CampaignUuid.String()
		urgentCampaign.CampaignUuid = CampaignUuid.String()

		// キャンペーンを新規追加する
		err = orm.CreateUrgentCampaign(tx, urgentCampaign)
		if err != nil {
			logging.LogError("failed to create UrgentCampaign", err)
			return err
		}

		return nil
	})

	return
}

func (service *UrgentCampaignService) GetUrgentCampaign(uuid string) (res dto.GetUrgentCampaigns, err error) {
	// トランザクションを開始する
	err = orm.Transaction(func(tx *gorm.DB) error {
		//特定のUUIDに一致するキャンペーンを取得
		campaign, err := orm.GetUrgentCampaign(tx, uuid)
		if err != nil {
			logging.LogError("failed to get UrgentCampaign", err)
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return custom_error.NewError(http.StatusBadRequest, "campaign uuid is not match")
			}
			return err
		}

		// レスポンスの構造体に格納
		res = dto.GetUrgentCampaigns{
			CampaignUuid:   campaign.CampaignUuid,
			RestaurantUuid: campaign.RestaurantUuid,
			StartAt:        campaign.StartAt,
			EndAt:          campaign.EndAt,
			Description:    campaign.Description,
			DiscountOffer:  campaign.DiscountOffer,
			CreatedAt:      campaign.CreatedAt,
		}

		return nil
	})
	return
}
