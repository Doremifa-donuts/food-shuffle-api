package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	logging "food-shuffle-api/log"
	"log"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) (bool, error) {
	// DBのテーブルが存在する場合はマイグレートをスキップする
	if db.Migrator().HasTable(&User{}) {
		return false, nil
	} else {
		// DBのマイグレーションを1つずつ実行する
		err := db.AutoMigrate(&User{})
		if err != nil {
			log.Fatalf("failed to migrate User: %v", err)
		}

		err = db.AutoMigrate(&RestaurantUser{})
		if err != nil {
			log.Fatalf("failed to migrate RestaurantUser: %v", err)
		}

		err = db.AutoMigrate(&GeneralUser{})
		if err != nil {
			log.Fatalf("failed to migrate GeneralUser: %v", err)
		}

		err = db.AutoMigrate(&PopupGroup{})
		if err != nil {
			log.Fatalf("failed to migrate PopupGroup: %v", err)
		}

		err = db.AutoMigrate(&PopupGroupSubmission{})
		if err != nil {
			log.Fatalf("failed to migrate PopupGroupSubmission: %v", err)
		}

		err = db.AutoMigrate(&Review{})
		if err != nil {
			log.Fatalf("failed to migrate Review: %v", err)
		}

		err = db.AutoMigrate(&ReviewLike{})
		if err != nil {
			log.Fatalf("failed to migrate ReviewLike: %v", err)
		}

		err = db.AutoMigrate(&ReviewArchive{})
		if err != nil {
			log.Fatalf("failed to migrate ReviewArchive: %v", err)
		}

		err = db.AutoMigrate(&ReviewReceive{})
		if err != nil {
			log.Fatalf("failed to migrate ReviewReceive: %v", err)
		}

		err = db.AutoMigrate(&PopupGroupSharedReviews{})
		if err != nil {
			log.Fatalf("failed to migrate PopupGroupSharedReviews: %v", err)
		}

		err = db.AutoMigrate(&Course{})
		if err != nil {
			log.Fatalf("failed to migrate Course: %v", err)
		}

		err = db.AutoMigrate(&UrgentCampaign{})
		if err != nil {
			log.Fatalf("failed to migrate UrgentCampaign: %v", err)
		}

		err = db.AutoMigrate(&Reservation{})
		if err != nil {
			log.Fatalf("failed to migrate Reservation: %v", err)
		}

		fmt.Println("Database migrated")
		return true, err
	}
}

func InsertSampleData(db *gorm.DB) error {
	// サンプルデータを挿入する
	err := db.Create(sampleUsers).Error
	if err != nil {
		logging.LogError("Error inserting sample data for Users", err)
	}

	err = db.Create(sampleRestaurantUsers).Error
	if err != nil {
		logging.LogError("Error inserting sample data for RestaurantUsers", err)
	}

	err = db.Create(sampleGeneralUsers).Error
	if err != nil {
		logging.LogError("Error inserting sample data for GeneralUsers", err)
	}

	err = db.Create(samplePopupGroups).Error
	if err != nil {
		logging.LogError("Error inserting sample data for PopupGroups", err)
	}

	err = db.Create(samplePopupGroupSubmissions).Error
	if err != nil {
		logging.LogError("Error inserting sample data for PopupGroupSubmissions", err)
	}

	err = db.Create(sampleReviews).Error
	if err != nil {
		logging.LogError("Error inserting sample data for Reviews", err)
	}

	err = db.Create(sampleReviewLikes).Error
	if err != nil {
		logging.LogError("Error inserting sample data for ReviewLikes", err)
	}

	err = db.Create(sampleReviewArchives).Error
	if err != nil {
		logging.LogError("Error inserting sample data for ReviewArchives", err)
	}

	err = db.Create(sampleReviewReceives).Error
	if err != nil {
		logging.LogError("Error inserting sample data for ReviewReceives", err)
	}

	err = db.Create(samplePopupGroupSharedReviews).Error
	if err != nil {
		logging.LogError("Error inserting sample data for PopupGroupSharedReviews", err)
	}

	err = db.Create(sampleCourses).Error
	if err != nil {
		logging.LogError("Error inserting sample data for Courses", err)
	}

	err = db.Create(sampleUrgentCampaigns).Error
	if err != nil {
		logging.LogError("Error inserting sample data for UrgentCampaigns", err)
	}

	err = db.Create(sampleReservations).Error
	if err != nil {
		logging.LogError("Error inserting sample data for Reservations", err)
	}

	return nil
}

// JSON配列を格納する構造体
type StringArray []string

// Scan は database/sql.Scanner インターフェースを実装
func (a *StringArray) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), a)
}

// Value は driver.Valuer インターフェースを実装
func (a StringArray) Value() (driver.Value, error) {
	return json.Marshal(a)
}
