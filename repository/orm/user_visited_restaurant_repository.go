package orm

import (
	"errors"
	"food-shuffle-api/repository/model"

	"gorm.io/gorm"
)

// レストランにユーザーが訪れたことがあるかをチェックする 存在しないばあいにも処理を行うことがある場合、コードが複雑化するので、ここでハンドリングを行う
func ExistsUserVisitedRestaurant(db *gorm.DB, userVisitedRestaurant model.UserVisitedRestaurant) (bool, error) {
	err := db.First(&userVisitedRestaurant).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

// レストランにチェックインする
func CreateUserVisitedRestaurant(db *gorm.DB, userVisitedRestaurant model.UserVisitedRestaurant) error {
	return db.Create(&userVisitedRestaurant).Error
}

// 最終訪問日の更新を行う
func UpdateLastVisitedTime(db *gorm.DB, userVisitedRestaurant model.UserVisitedRestaurant) error {
	return db.Updates(&userVisitedRestaurant).Error
}

// ユーザーIDが一致し、レストランUUIDのリストには含まれないもののみを取得
func ListFilterRestaurantUuidsByUserUuidNotInRestaurantUuids(db *gorm.DB, userUuid string, restaurantUuids []string) ([]string, error) {
	var filteredUuids []string
	var err error
	if len(restaurantUuids) == 0 {
		// restaurantUuids が空の場合は単純に user_uuid の条件だけで検索
		err = db.Model(&model.UserVisitedRestaurant{}).
			Where("user_uuid = ?", userUuid).
			Pluck("restaurant_uuid", &filteredUuids).Error
	} else {
		// restaurantUuids が存在する場合は not in 条件も含める
		err = db.Model(&model.UserVisitedRestaurant{}).
			Where("user_uuid = ? and restaurant_uuid not in (?)", userUuid, restaurantUuids).
			Pluck("restaurant_uuid", &filteredUuids).Error
	}
	return filteredUuids, err
}

func ListRestaurantUuidsByUserUuid(db *gorm.DB, userUuid string) ([]string, error) {
	var filteredUuids []string
	err := db.Model(&model.UserVisitedRestaurant{}).Where("user_uuid = ?", userUuid).Pluck("restaurant_uuid", &filteredUuids).Error
	return filteredUuids, err

}
