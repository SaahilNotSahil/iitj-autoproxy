package pkg

import (
	"time"

	"github.com/go-co-op/gocron"
)

var Scheduler *gocron.Scheduler

func init() {
	Scheduler = gocron.NewScheduler(time.Local)
}
