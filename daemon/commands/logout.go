package commands

import (
	"log"

	"github.com/SaahilNotSahil/iitj-autoproxy/pkg"
	"github.com/SaahilNotSahil/iitj-autoproxy/pkg/daemon"
)

func LogoutCmd() {
	token, err := pkg.GetCurrentKeepaliveToken()
	if err != nil {
		pkg.Logger.Println(err)
		log.Println(err)

		err = daemon.SendMessageToCLI("User not logged in")
		if err != nil {
			pkg.Logger.Println(err)
			log.Println(err)

			return
		}

		return
	}

	if token == "" {
		pkg.Logger.Println("User not logged in")
		log.Println("User not logged in")

		err := daemon.SendMessageToCLI("User not logged in")
		if err != nil {
			pkg.Logger.Println(err)
			log.Println(err)
		}

		return
	}

	err = pkg.Logout(token)
	if err != nil {
		pkg.Logger.Println(err)
		log.Println(err)

		err = daemon.SendMessageToCLI(err.Error())
		if err != nil {
			pkg.Logger.Println(err)
			log.Println(err)
		}
	} else {
		pkg.Logger.Println("Logged out successfully")
		log.Println("Logged out successfully")

		err = daemon.SendMessageToCLI("Logged out successfully")
		if err != nil {
			pkg.Logger.Println(err)
			log.Println(err)
		}

		KillScheduler()
	}
}

func LogoutDummyCmd() {
	err := daemon.SendMessageToCLI("Dummy logout successful")
	if err != nil {
		pkg.Logger.Println(err)
		log.Println(err)

		return
	}

	KillDummyScheduler()
}
