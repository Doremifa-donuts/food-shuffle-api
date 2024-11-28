package service

import(
	"food-shuffle-api/dto"
	logging "food-shuffle-api/log"
	"food-shuffle-api/model"
	"food-shuffle-api/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UrgentCampaignService struct{}

func (service *UrgentCampaignService) UrgentCampaignRegister(urgentCampaign model.UrgentCampaign) (res dto.CreateUrgentCampaign, err error) {

	err = repository.Transaction(func(tx *gorm.DB) error {
		CampaignUuid, err := uuid.NewV7()
		if err != nil {
			logging.LogError("failed to generate CampaignUuid", err)
			return err
		}
		res.CampaignUuid = CampaignUuid.String()

		urgentCampaign.CampaignUuid = CampaignUuid.String()

		err = repository.CreateUrgentCampaign(tx, urgentCampaign)
		if err != nil {
			logging.LogError("failed to create UrgentCampaign", err)
			return err
		}
		
		return nil
	})

	return
}

func (service *UrgentCampaignService) GetUrgentCampaign(uuid string) (res dto.GetUrgentCampaigns, err error) {
    err = repository.Transaction(func(tx *gorm.DB) error {
        //特定のUUIDに一致するキャンペーンを取得
        urgentCampaigns, err := repository.GetUrgentCampaign(tx, uuid)
        if err != nil {
            logging.LogError("failed to get UrgentCampaign", err)
            return err
        }
        
        // 取得したキャンペーンの中から、指定されたUUIDと一致するものを探す
        for _, campaign := range urgentCampaigns {
            if campaign.CampaignUuid == uuid {
                res = dto.GetUrgentCampaigns{
                    CampaignUuid:    campaign.CampaignUuid,
                    RestaurantUuid:  campaign.RestaurantUuid,
                    StartAt:         campaign.StartAt,
                    EndAt:           campaign.EndAt,
                    Description:     campaign.Description,
                    DiscountOffer:   campaign.DiscountOffer,
                    CreatedAt:       campaign.CreatedAt,
                }
                break // 最初に一致したものだけを取得
            }
        }
        
        return nil
    })
    return
}
