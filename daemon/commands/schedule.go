package commands

import (
	"log"

	"github.com/spf13/viper"

	"github.com/SaahilNotSahil/iitj-autoproxy/pkg"
	"github.com/SaahilNotSahil/iitj-autoproxy/pkg/daemon"
	ds "github.com/SaahilNotSahil/iitj-autoproxy/pkg/dummy/scheduler"
	"github.com/SaahilNotSahil/iitj-autoproxy/pkg/scheduler"
)

var (
	kill      = make(chan bool)
	killDummy = make(chan bool)
)

func ScheduleCmd() {
	err := viper.ReadInConfig()
	if err != nil {
		pkg.Logger.Println("Error reading config file")
		log.Println("Error reading config file")

		err = daemon.SendMessageToCLI(
			"Error reading config file",
		)
		if err != nil {
			pkg.Logger.Println(err)
			log.Println(err)
		}

		return
	}

	username := viper.GetString("username")
	password := viper.GetString("password")
	baseURL := viper.GetString("base_url")

	if username == "" || password == "" {
		pkg.Logger.Println("Please configure the application using the config command")
		log.Println("Please configure the application using the config command")

		err = daemon.SendMessageToCLI(
			"Please configure the application using the config command",
		)
		if err != nil {
			pkg.Logger.Println(err)
			log.Println(err)
		}

		return
	}

	scheduler_running_state := scheduler.RunLoginScheduler(
		baseURL, username, password, kill,
	)
	if !scheduler_running_state {
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

		viper.Set("scheduler_running_state", scheduler_running_state)

		err = viper.WriteConfig()
		if err != nil {
			pkg.Logger.Println(err)
			log.Println(err)

			err = daemon.SendMessageToCLI(err.Error())
			if err != nil {
				pkg.Logger.Println(err)
				log.Println(err)
			}
		}

		err = daemon.SendMessageToCLI("Scheduler started")
		if err != nil {
			pkg.Logger.Println(err)
			log.Println(err)
		}
	}
}

func ScheduleDummyCmd() {
	err := viper.ReadInConfig()
	if err != nil {
		pkg.Logger.Println("Error reading config file")
		log.Println("Error reading config file")

		err = daemon.SendMessageToCLI(
			"Error reading config file",
		)
		if err != nil {
			pkg.Logger.Println(err)
			log.Println(err)
		}

		return
	}

	username := viper.GetString("username")
	password := viper.GetString("password")
	_ = viper.GetString("base_url")

	if username == "" || password == "" {
		pkg.Logger.Println("Please configure the application using the config command")
		log.Println("Please configure the application using the config command")

		err = daemon.SendMessageToCLI(
			"Please configure the application using the config command",
		)
		if err != nil {
			pkg.Logger.Println(err)
			log.Println(err)
		}

		return
	}

	scheduler_running_state := ds.RunLoginScheduler(killDummy)
	if !scheduler_running_state {
		pkg.Logger.Println("Dummy scheduler already running")
		log.Println("Dummy scheduler already running")

		err = daemon.SendMessageToCLI("Dummy scheduler already running")
		if err != nil {
			pkg.Logger.Println(err)
			log.Println(err)
		}
	} else {
		pkg.Logger.Println("Scheduler started")
		log.Println("Scheduler started")

		viper.Set("dummy_scheduler_running_state", scheduler_running_state)

		err = viper.WriteConfig()
		if err != nil {
			pkg.Logger.Println(err)
			log.Println(err)

			err = daemon.SendMessageToCLI(err.Error())
			if err != nil {
				pkg.Logger.Println(err)
				log.Println(err)
			}
		}

		err = daemon.SendMessageToCLI("Dummy scheduler started")
		if err != nil {
			pkg.Logger.Println(err)
			log.Println(err)
		}
	}
}

func KillScheduler() {
	err := viper.ReadInConfig()
	if err != nil {
		pkg.Logger.Println("Error reading config file")
		log.Println("Error reading config file")

		err = daemon.SendMessageToCLI(
			"Error reading config file",
		)
		if err != nil {
			pkg.Logger.Println(err)
			log.Println(err)

			return
		}

		return
	}

	if scheduler.SchedulerRunning {
		pkg.Logger.Println("Stopping the scheduler...")
		log.Println("Stopping the scheduler...")

		err := daemon.SendMessageToCLI("Stopping the scheduler...")
		if err != nil {
			pkg.Logger.Println(err)
			log.Println(err)
		}

		scheduler.SchedulerRunning = false

		viper.Set("scheduler_running_state", false)

		err = viper.WriteConfig()
		if err != nil {
			pkg.Logger.Println(err)
			log.Println(err)

			err = daemon.SendMessageToCLI(err.Error())
			if err != nil {
				pkg.Logger.Println(err)
				log.Println(err)
			}
		}

		kill <- true
	}
}

func KillDummyScheduler() {
	err := viper.ReadInConfig()
	if err != nil {
		pkg.Logger.Println("Error reading config file")
		log.Println("Error reading config file")

		err = daemon.SendMessageToCLI(
			"Error reading config file",
		)
		if err != nil {
			pkg.Logger.Println(err)
			log.Println(err)

			return
		}

		return
	}

	if ds.SchedulerRunning {
		pkg.Logger.Println("Stopping the dummy scheduler...")
		log.Println("Stopping the dummy scheduler...")

		err := daemon.SendMessageToCLI("Stopping the dummy scheduler...")
		if err != nil {
			pkg.Logger.Println(err)
			log.Println(err)
		}

		ds.SchedulerRunning = false

		viper.Set("dummy_scheduler_running_state", false)

		err = viper.WriteConfig()
		if err != nil {
			pkg.Logger.Println(err)
			log.Println(err)

			err = daemon.SendMessageToCLI(err.Error())
			if err != nil {
				pkg.Logger.Println(err)
				log.Println(err)
			}
		}

		killDummy <- true
	}
}
