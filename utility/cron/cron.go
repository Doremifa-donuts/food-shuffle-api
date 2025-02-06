package cron

import (
	"github.com/robfig/cron/v3"
)

func Run() {
	c := cron.New()
	// ジョブを設定する
	c.AddFunc("@every 5m", ProivideBoost)
	// お助けブーストの通知テスト @every ?? で通知の間隔を設定できる 他にも様々な実行タイミングがあるので調べてみてね
	// https://pkg.go.dev/github.com/robfig/cron?utm_source=godoc#section-documentation
	// c.AddFunc("@every 10s", func() {
	// 	ws.SetBoost([]string{"91a78381-f472-496b-90e3-2c66a33391d1"}, ws.BoostContent{RestaurantName: "test", BoostUuid: "0193a8ee-6972-7a4e-bc20-71de6517b565"})
	// })
	c.Start()
}
