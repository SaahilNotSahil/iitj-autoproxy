package commands

import (
	"log"
	"os"

	"github.com/XanderWatson/iitj-autoproxy/pkg"
	"github.com/XanderWatson/iitj-autoproxy/pkg/daemon/scheduler"

	"github.com/spf13/viper"
)

var kill = make(chan bool)

func ScheduleCmd() {
	err := viper.ReadInConfig()
	if err != nil {
		pkg.Logger.Println("Error reading config file")
		os.Exit(1)
	}

	username := viper.GetString("username")
	password := viper.GetString("password")
	base_url := viper.GetString("base_url")

	if username == "" || password == "" {
		pkg.Logger.Println("Please configure the application using the config command")
		log.Println("Please configure the application using the config command")
	}

	if !scheduler.RunLoginScheduler(base_url, username, password, kill) {
		pkg.Logger.Println("Scheduler already running")
		log.Println("Scheduler already running")
	} else {
		pkg.Logger.Println("Scheduler started")
		log.Println("Scheduler started")
	}
}

func KillScheduler() {
	if scheduler.SchedulerRunning {
		pkg.Logger.Println("Stopping the scheduler")
		log.Println("Stopping the scheduler")

		scheduler.SchedulerRunning = false

		kill <- true
	}
}
