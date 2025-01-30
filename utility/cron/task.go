package cron

import (
	"fmt"
	logging "food-shuffle-api/log"
	"food-shuffle-api/repository/orm"
	"food-shuffle-api/repository/redis"
	"food-shuffle-api/utility/parameters"
	"food-shuffle-api/ws"
	"time"

	"gorm.io/gorm"
)

// 開始時刻付近のお助けブーストを周囲のユーザーに告知する
func ProivideBoost() {

	// 店舗名とその付近にいるユーザーのリスト
	boostList := make(map[string][]string)

	orm.Transaction(func(tx *gorm.DB) error {
		//　開始地獄付近のお助けブーストの情報を取得
		urgentCampaigns, err := orm.ListUrgentCampaignByStartAt(tx, time.Now())
		if err != nil {
			logging.LogError("failed to get urgent campaigns", err)
			return err
		}
		fmt.Println(urgentCampaigns)
		// 通知対象1件ずつに対して通知を行う
		for _, ururgentCampaign := range urgentCampaigns {
			// お知らせブースト対象の店舗を取得する
			restaurant, err := orm.GetRestaurantDetail(tx, ururgentCampaign.RestaurantUuid)
			if err != nil {
				logging.LogError("failed to get restaurant detail", err)
				return err
			}
			// レストランの位置情報から通知範囲内にいる人を取得
			userUuids, err := redis.GetUserUuidsByRestaurantBoostRadius(restaurant.Latitude, restaurant.Longitude, parameters.BOOST_RADIUS)
			if err != nil {
				logging.LogError("failed to get user uuid list", err)
				return err
			}

			// レストラン名をユーザーIDをセットで格納する
			boostList[restaurant.RestaurantName] = userUuids
		}
		fmt.Println(err)
		return nil
	})
	// WebSocketで通知する
	ws.SetBoost(boostList)
}
