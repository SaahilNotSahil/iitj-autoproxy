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

	var username string
	var password string

	keys, _ := keystore.Keys()
	log.Println(keys)

	if len(keys) == 1 {
		username = keys[0]
		password, err = keystore.Get(username)
		if err != nil || password == "" {
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
	} else if len(keys) == 0 {
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
	} else {
		err := keystore.Reset()
		if err != nil {
			pkg.Logger.Println(err)
			log.Println(err)

			err = daemon.SendMessageToCLI(
				"Error resetting the OS keyring",
			)
			if err != nil {
				pkg.Logger.Println(err)
				log.Println(err)

				return
			}

			return
		}

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
