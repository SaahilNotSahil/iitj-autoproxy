package commands

import (
	"log"
	"os"

	"github.com/XanderWatson/iitj-autoproxy/pkg"
	"github.com/XanderWatson/iitj-autoproxy/pkg/daemon"

	"github.com/spf13/viper"
)

func ScheduleCmd() {
	err := viper.ReadInConfig()
	if err != nil {
		pkg.Logger.Println("Error reading config file")
		os.Exit(1)
	}

	username := viper.GetString("username")
	password := viper.GetString("password")

	if username == "" || password == "" {
		pkg.Logger.Println("Please configure the application using the config command")
		log.Println("Please configure the application using the config command")
	}

	go daemon.RunLoginScheduler(username, password)
}
