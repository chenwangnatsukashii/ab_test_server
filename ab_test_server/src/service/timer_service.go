package service

import (
	"github.com/robfig/cron/v3"
	"log"
)

func StartTimer(interval int) {
	c := cron.New()
	c.AddFunc("@every 5s", func() {
		if err := FinalResult(); err != nil {
			log.Println(err)
		}
	})
	go c.Start()
}
