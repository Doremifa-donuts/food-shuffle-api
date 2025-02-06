package cron

import (
	"github.com/robfig/cron/v3"
)

func Run() {
	c := cron.New()
	// ジョブを設定する
	c.AddFunc("@every 5m", ProivideBoost)
	// c.AddFunc("@every 10s", func() {
	// 	ws.SetBoost([]string{"91a78381-f472-496b-90e3-2c66a33391d1"}, ws.BoostContent{RestaurantName: "test", BoostUuid: "testUuid"})
	// })
	c.Start()
}
