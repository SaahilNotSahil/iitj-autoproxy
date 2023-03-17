package pkg

import (
	"github.com/go-co-op/gocron"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"time"
)

var Scheduler *gocron.Scheduler

func init() {
	Scheduler = gocron.NewScheduler(time.Local)
}

func RunLoginScheduler(username string, password string) {
	viper.Set("pid", os.Getpid())
	viper.WriteConfig()

	_, err := Scheduler.Every(10001).Seconds().Do(schedule, username, password)
	cobra.CheckErr(err)

	Logger.Println("Starting Scheduler")
	Scheduler.StartBlocking()
}

func schedule(username string, password string) {
	for i := 0; i < 5; i++ {
		err := Login(viper.GetString("base_url"), username, password)
		if err == nil {
			Logger.Println("Logged in successfully")
			break
		} else {
			Logger.Printf("Attempted login %d", i+1)
		}

		time.Sleep(5 * time.Second)
	}
}
