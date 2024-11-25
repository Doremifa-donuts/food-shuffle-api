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

func (service *UrgentCampaignService) UrgentCampaignRegister(urgentCampaign model.UrgentCampaign) (res dto.UrgentCampaign, err error) {

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