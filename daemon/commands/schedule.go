package commands

import (
	"log"

	"github.com/XanderWatson/iitj-autoproxy/pkg"
	"github.com/XanderWatson/iitj-autoproxy/pkg/daemon"
	"github.com/XanderWatson/iitj-autoproxy/pkg/daemon/scheduler"
	"github.com/XanderWatson/iitj-autoproxy/pkg/keystore"

	"github.com/spf13/viper"
)

var kill = make(chan bool)

func ScheduleCmd() {
	err := viper.ReadInConfig()
	if err != nil {
		pkg.Logger.Println(err)
		log.Println(err)

		err = daemon.SendMessageToCLI(
			"Error reading config file. Please make sure the file exists and is valid",
		)
		if err != nil {
			pkg.Logger.Println(err)
			log.Println(err)

			return
		}

		return
	}

	username, err := keystore.Get("username")
	if err != nil {
		pkg.Logger.Println(err)
		log.Println(err)

		err = daemon.SendMessageToCLI(
			"Error fetching the username from the OS keyring",
		)
		if err != nil {
			pkg.Logger.Println(err)
			log.Println(err)

			return
		}

		return
	}

	password, err := keystore.Get("password")
	if err != nil {
		pkg.Logger.Println(err)
		log.Println(err)

		err = daemon.SendMessageToCLI(
			"Error fetching the password from the OS keyring",
		)
		if err != nil {
			pkg.Logger.Println(err)
			log.Println(err)

			return
		}

		return
	}

	base_url := viper.GetString("base_url")

	if username == "" || password == "" {
		pkg.Logger.Println(
			"Please configure the application using the config command",
		)
		log.Println(
			"Please configure the application using the config command",
		)

		err = daemon.SendMessageToCLI(
			"Please configure the application using the config command",
		)
		if err != nil {
			pkg.Logger.Println(err)
			log.Println(err)

			return
		}

		return
	}

	if !scheduler.RunLoginScheduler(base_url, username, password, kill) {
		pkg.Logger.Println("Scheduler already running")
		log.Println("Scheduler already running")

		err = daemon.SendMessageToCLI("Scheduler already running")
		if err != nil {
			pkg.Logger.Println(err)
			log.Println(err)
		}
	} else {
		pkg.Logger.Println("Scheduler started")
		log.Println("Scheduler started")

		err = daemon.SendMessageToCLI("Scheduler started")
		if err != nil {
			pkg.Logger.Println(err)
			log.Println(err)
		}
	}
}

func KillScheduler() {
	if scheduler.SchedulerRunning {
		pkg.Logger.Println("Stopping the scheduler...")
		log.Println("Stopping the scheduler...")

		err := daemon.SendMessageToCLI("Stopping the scheduler...")
		if err != nil {
			pkg.Logger.Println(err)
			log.Println(err)
		}

		scheduler.SchedulerRunning = false

		kill <- true
	}
}
