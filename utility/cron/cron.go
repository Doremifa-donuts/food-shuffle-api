package cron

import (
	"github.com/robfig/cron/v3"
)

func Run() {
	c := cron.New()
	// ジョブを設定する
	c.AddFunc("@every 5m", ProivideBoost)
	// c.AddFunc("@every 10s", func() {
	// 	ws.SetBoost(map[string][]string{
	// 		"test": {"91a78381-f472-496b-90e3-2c66a33391d1"},
	// 	})
	// })
	c.Start()
}
