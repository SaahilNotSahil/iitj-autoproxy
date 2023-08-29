package commands

import (
	"log"

	"github.com/XanderWatson/iitj-autoproxy/pkg"
	"github.com/XanderWatson/iitj-autoproxy/pkg/daemon"
)

func LogoutCmd() {
	token, err := daemon.GetCurrentKeepaliveToken()
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

	err = daemon.Logout(token)
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
