package repository

import (
	"errors"
	"food-shuffle-api/model"
	"food-shuffle-api/utility/custom_error"
	"net/http"
	"food-shuffle-api/dto"

	"gorm.io/gorm"
)

// レストランUUIDからレストラン名を取得する
func GetRestaurantNameByRestaurantUuid(db *gorm.DB, uuid string) (string, error) {
	var restoUser model.RestaurantUser
	err := db.Where("restaurant_uuid = ?", uuid).Find(&restoUser).Error
	return restoUser.RestaurantName, err
}

// 混雑状況が満席でないことを確認する
func CheckNotPackedStatusByRestaurantUuid(db *gorm.DB, restaurantUuid string) error {
	return db.Where("restaurant_uuid = ? and busy_status <> ?", restaurantUuid, model.Packed).First(&model.RestaurantUser{}).Error
}

// レストランの詳細情報を取得する
func GetRestaurantDetail(db *gorm.DB, RestaurantUuid string) (model.RestaurantUser, error) {
	var restaurantUser model.RestaurantUser
	err := db.Where("restaurant_uuid = ?", RestaurantUuid).First(&restaurantUser).Error
	return restaurantUser, err
}

// レストランUUIDのリストから商法を取得する
func ListRestaurantByRestaurantUuids(db *gorm.DB, restaurantUuids []string) ([]model.RestaurantUser, error) {
	var restaurants []model.RestaurantUser
	err := db.Where("restaurant_uuid in (?)", restaurantUuids).Find(&restaurants).Error
	return restaurants, err
}

// 混雑状況のステータスを更新する
func PutBusyStatus(db *gorm.DB, restaurantUser model.RestaurantUser) (bool, error) {
	result := db.Model(&model.RestaurantUser{}).Where("restaurant_uuid = ?", restaurantUser.RestaurantUuid).Update("busy_status", restaurantUser.BusyStatus)
	return result.RowsAffected == 1, result.Error
}

// レストランUUIDが存在することを確かめる
func ExistsRestaurantByRestaurantUuid(db *gorm.DB, restaurantUuid string) error {
	err := db.Where("restaurant_uuid", restaurantUuid).First(&model.RestaurantUser{}).Error
	// リソースなしエラーはカスタムエラーとして返す
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return custom_error.NewError(http.StatusBadRequest, "restaurant is not found")
	}
	return err
}

func GetWentPlaces(db *gorm.DB, uuid string) ([]dto.WentPlaces, error) {
	var wentPlaces []dto.WentPlaces
	
	err := db.Table("user_visited_restaurants").
	Select("restaurant_users.restaurant_uuid, restaurant_users.restaurant_name, restaurant_users.latitude, restaurant_users.longitude").
	Joins("JOIN restaurant_users ON user_visited_restaurants.restaurant_uuid = restaurant_users.restaurant_uuid").
	Where("user_visited_restaurants.user_uuid = ?", uuid).
	Find(&wentPlaces).Error

	return wentPlaces, err
}