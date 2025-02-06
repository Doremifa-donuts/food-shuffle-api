package cron

import (
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
	orm.Transaction(func(tx *gorm.DB) error {
		//　開始地獄付近のお助けブーストの情報を取得
		urgentCampaigns, err := orm.ListUrgentCampaignByStartAt(tx, time.Now())
		if err != nil {
			logging.LogError("failed to get urgent campaigns", err)
			return err
		}

		// 通知対象1件ずつに対して通知を行う
		for _, urgentCampaign := range urgentCampaigns {

			// お知らせブースト対象の店舗を取得する
			restaurant, err := orm.GetRestaurantDetail(tx, urgentCampaign.RestaurantUuid)
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

			// WebSocketで通知する
			ws.SetBoost(userUuids, ws.BoostContent{RestaurantName: restaurant.RestaurantName, BoostUuid: urgentCampaign.CampaignUuid})
		}

		return nil
	})

}
