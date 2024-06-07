package commands

import (
	"log"

	"github.com/spf13/viper"

	"github.com/SaahilNotSahil/iitj-autoproxy/pkg"
	"github.com/SaahilNotSahil/iitj-autoproxy/pkg/daemon"
)

func LoginCmd() {
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

			return
		}

		return
	}

	err = pkg.Login(baseURL, username, password)
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

func LoginDummyCmd() {
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

	_ = viper.GetString("username")
	_ = viper.GetString("password")
	_ = viper.GetString("base_url")

	err = daemon.SendMessageToCLI("Dummy login successful")
	if err != nil {
		pkg.Logger.Println(err)
		log.Println(err)

		return
	}
}
