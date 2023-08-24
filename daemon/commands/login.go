package commands

import (
	"log"
	"os"

	"github.com/XanderWatson/iitj-autoproxy/pkg"
	"github.com/XanderWatson/iitj-autoproxy/pkg/daemon"
	"github.com/XanderWatson/iitj-autoproxy/pkg/keystore"

	"github.com/spf13/viper"
)

func LoginCmd() {
	err := viper.ReadInConfig()
	if err != nil {
		pkg.Logger.Println("Error reading config file")
		os.Exit(1)
	}

	username, _ := keystore.Get("username")
	password, _ := keystore.Get("password")

	if username == "" || password == "" {
		pkg.Logger.Println(
			"Please configure the application using the config command",
		)
		log.Println("Please configure the application using the config command")
	}

	err = daemon.Login(viper.GetString("base_url"), username, password)
	if err != nil {
		pkg.Logger.Println("Login failed")
		log.Fatal("Login failed")
	}

	pkg.Logger.Println("Login successful")
	log.Println("Login successful")
}
