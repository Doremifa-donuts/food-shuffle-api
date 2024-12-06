package repository

import (
	"errors"
	"food-shuffle-api/model"

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
