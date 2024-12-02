package service

import (
	logging "food-shuffle-api/log"
	"food-shuffle-api/model"
	"food-shuffle-api/repository"
	"food-shuffle-api/utility/custom_error"
	"net/http"

	"gorm.io/gorm"
)

type RestaurantUserService struct{}

func (s *RestaurantUserService) PutBusyStatus(restaurantUser model.RestaurantUser) (err error) {

	err = repository.Transaction(func(tx *gorm.DB) error {
		result, err := repository.PutBusyStatus(tx, restaurantUser)
		if err != nil {
			logging.LogError("failed to put busy status", err)
			return err
		}
		if !result {
			logging.LogError("failed to put busy status", err)
			return custom_error.NewError(http.StatusBadRequest, "Restaurant user not found")
		}

		// トランザクションを終了する
		return nil
	})
	return

}
