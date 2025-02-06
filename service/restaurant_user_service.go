package service

import (
	"food-shuffle-api/dto"
	logging "food-shuffle-api/log"
	"food-shuffle-api/repository/model"
	"food-shuffle-api/repository/orm"
	"food-shuffle-api/utility/custom_error"
	"net/http"

	"gorm.io/gorm"
)

type RestaurantUserService struct{}

func (s *RestaurantUserService) PutBusyStatus(restaurantUser model.RestaurantUser) (err error) {

	err = orm.Transaction(func(tx *gorm.DB) error {
		result, err := orm.PutBusyStatus(tx, restaurantUser)
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

func (s *RestaurantUserService) GetRestaurantPlace(uuid string) (restaurant dto.RestaurantPlace, err error) {

	err = orm.Transaction(func(tx *gorm.DB) error {
		restaurant, err = orm.GetRestaurantPlace(tx, uuid)
		if err != nil {
			logging.LogError("failed to get restaurant place", err)
			return err
		}
		return nil
	})
	return
}
