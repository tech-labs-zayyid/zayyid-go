package cron

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

func Run() {
	s := gocron.NewScheduler(time.Local)

	s.Every(1).Second().Do(func() {
		fmt.Println("cron is start...")
	})

	s.StartAsync()
}
