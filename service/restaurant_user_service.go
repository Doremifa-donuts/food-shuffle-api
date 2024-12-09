package service

import (
	"food-shuffle-api/model"
	"food-shuffle-api/repository"

	"gorm.io/gorm"
)

type RestaurantUserService struct{}

type restaurant struct {
	RestoName  string
	BusyStatus model.BusyStatus
}

// レストランを取得する(busy_status確認用)
func (s RestaurantUserService) GetRestaurantByUuid(uuid string) (restaurant, error) {
	// 返り値
	var restaurantResponse restaurant

	// トランザクションを開始する
	err := repository.Transaction(func(tx *gorm.DB) error {
		resto, err := repository.GetRestaurantUserByUuid(tx, uuid)
		if err != nil {
			return err
		}
		restaurantResponse.RestoName = resto.RestoName
		restaurantResponse.BusyStatus = resto.BusyStatus
		return nil
	})
	return restaurantResponse, err
}

// レストランの混雑状況を更新する
func (s RestaurantUserService) UpdateBusyStatus(uuid string, busyStatusInput model.BusyStatus) (string, error) {
	message := "yeahh"
	// トランザクションを開始する
	err := repository.Transaction(func(tx *gorm.DB) error {
		return repository.UpdateBusyStatus(tx, uuid, busyStatusInput)
	})
	return message, err
}
