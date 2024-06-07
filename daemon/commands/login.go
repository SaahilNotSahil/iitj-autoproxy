package commands

import (
	"log"

	"github.com/spf13/viper"

	"github.com/SaahilNotSahil/iitj-autoproxy/pkg"
	"github.com/SaahilNotSahil/iitj-autoproxy/pkg/daemon"
	"github.com/SaahilNotSahil/iitj-autoproxy/pkg/keystore"
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

	username, err = keystore.Get("username")
	if err != nil {
		pkg.Logger.Println(err)
		log.Println(err)

		err = daemon.SendMessageToCLI(
			err.Error(),
		)
		if err != nil {
			pkg.Logger.Println(err)
			log.Println(err)

			return
		}

		return
	}

	password, err = keystore.Get("password")
	if err != nil {
		pkg.Logger.Println(err)
		log.Println(err)

		err = daemon.SendMessageToCLI(
			err.Error(),
		)
		if err != nil {
			pkg.Logger.Println(err)
			log.Println(err)

			return
		}

		return
	}

	err = pkg.Login(viper.GetString("base_url"), username, password)
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

	_, err = keystore.Get("username")
	if err != nil {
		pkg.Logger.Println(err)
		log.Println(err)

		err = daemon.SendMessageToCLI(
			err.Error(),
		)
		if err != nil {
			pkg.Logger.Println(err)
			log.Println(err)

			return
		}

		return
	}

	_, err = keystore.Get("password")
	if err != nil {
		pkg.Logger.Println(err)
		log.Println(err)

		err = daemon.SendMessageToCLI(
			err.Error(),
		)
		if err != nil {
			pkg.Logger.Println(err)
			log.Println(err)

			return
		}

		return
	}

	err = daemon.SendMessageToCLI("Dummy login successful")
	if err != nil {
		pkg.Logger.Println(err)
		log.Println(err)

		return
	}
}
