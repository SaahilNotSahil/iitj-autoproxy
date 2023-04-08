package daemon

import (
	"github.com/go-co-op/gocron"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"time"
	"github.com/XanderWatson/iitj-autoproxy/pkg"
)

var Scheduler *gocron.Scheduler

func init() {
	Scheduler = gocron.NewScheduler(time.Local)
}

func RunLoginScheduler(username string, password string) {
	viper.Set("pid1", os.Getpid())
	viper.WriteConfig()

	_, err := Scheduler.Every(5).Seconds().Do(schedule, username, password)
	cobra.CheckErr(err)

	pkg.Logger.Println("Starting Scheduler")

	go goroutine()
	select {}
}

func schedule(username string, password string) {
	err := Login(viper.GetString("base_url"), username, password)
	if err == nil {
		pkg.Logger.Println("Logged in successfully")
	} else {
		pkg.Logger.Println("Attempted login")
	}
}

func goroutine() {
	viper.Set("pid2", os.Getpid())
	viper.WriteConfig()
	Scheduler.StartBlocking()
}
