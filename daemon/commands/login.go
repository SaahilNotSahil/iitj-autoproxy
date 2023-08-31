package commands

import (
	"log"

	"github.com/XanderWatson/iitj-autoproxy/pkg"
	"github.com/XanderWatson/iitj-autoproxy/pkg/daemon"
	"github.com/XanderWatson/iitj-autoproxy/pkg/keystore"

	"github.com/spf13/viper"
)

func LoginCmd() {
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

	err = daemon.Login(viper.GetString("base_url"), username, password)
	if err != nil {
		pkg.Logger.Println(err)
		log.Println(err)

		err = daemon.SendMessageToCLI("Login failed")
		if err != nil {
			pkg.Logger.Println(err)
			log.Println(err)
		}
	} else {
		pkg.Logger.Println("Login successful")
		log.Println("Login successful")

		err = daemon.SendMessageToCLI("Login successful")
		if err != nil {
			pkg.Logger.Println(err)
			log.Println(err)
		}
	}
}
