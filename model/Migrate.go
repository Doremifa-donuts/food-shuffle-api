package model

import (
	"fmt"
	logging "food-shuffle-api/log"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {
	// DBのマイグレーションを実行する
	if err := db.AutoMigrate(&User{}, &RestoUser{}, &Group{}, &GroupSubmission{}, &Review{}, &ReviewFavorite{}, &ReviewArchive{}, &Course{}, &UrgentCampaign{}, &GroupSharedReviews{}, &Reservation{}); err != nil {
		logging.LogError("Error migrating database", err)
	} else {
		fmt.Println("Migrated database")
	}
}
