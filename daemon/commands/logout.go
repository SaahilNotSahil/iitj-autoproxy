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
	}

	if token == "" {
		pkg.Logger.Println("User not logged in")
		log.Println("User not logged in")
	}

	err = daemon.Logout(token)
	if err != nil {
		pkg.Logger.Println(err)
		log.Println(err)
	}

	pkg.Logger.Println("Logged out successfully")
	log.Println("Logged out successfully")
}
