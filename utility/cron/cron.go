package cron

import (
	"github.com/robfig/cron/v3"
)

func Run() {
	c := cron.New()
	// ジョブを設定する
	c.AddFunc("@every 5m", ProivideBoost)
	c.Start()
}
