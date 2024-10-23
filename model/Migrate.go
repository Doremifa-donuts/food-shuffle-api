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
		return
	} else {
		fmt.Println("Migrated database")
	}

	// サンプルデータを挿入する
	insertSampleData(db)
}

func insertSampleData(db *gorm.DB) {

	// サンプルデータを挿入する
	db.Create(SampleUsers)
	fmt.Println("Users inserted")
}
