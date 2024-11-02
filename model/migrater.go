package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	logging "food-shuffle-api/log"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {
	// DBのテーブルが存在する場合はマイグレートをスキップする
	if db.Migrator().HasTable(&User{}) {
		fmt.Println("Database already migrated")
		return nil
	} else {
		// DBのマイグレーションを実行する
		err := db.AutoMigrate(
			&User{},
			&RestaurantUser{},
			&GeneralUser{},
			&PopupGroup{},
			&PopupGroupSubmission{},
			&Review{},
			&ReviewFavorite{},
			&ReviewArchive{},
			&ReviewReceive{},
			&Course{},
			&UrgentCampaign{},
			&PopupGroupSharedReviews{},
			&Reservation{},
		)
		if err != nil {
			logging.LogError("Error migrating database", err)
			return err
		} else {
			fmt.Println("Migrated database")
		}
		// サンプルデータを挿入する
		insertSampleData(db)
		return nil
	}
}

func insertSampleData(db *gorm.DB) {

	// サンプルデータを挿入する
	db.Create(SampleUsers)
	db.Create(SampleRestaurantUsers)
	db.Create(SampleGeneralUsers)
	fmt.Println("Inserted sample data")
}

type StringArray []string

// Scan は database/sql.Scanner インターフェースを実装
func (a *StringArray) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), a)
}

// Value は driver.Valuer インターフェースを実装
func (a StringArray) Value() (driver.Value, error) {
	return json.Marshal(a)
}
