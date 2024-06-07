package commands

import (
	"log"

	"github.com/SaahilNotSahil/iitj-autoproxy/pkg"
	"github.com/SaahilNotSahil/iitj-autoproxy/pkg/daemon"
)

func HealthCheckCmd() {
	err := daemon.SendMessageToCLI("Healthy")
	if err != nil {
		errMsg := err.Error()
		errMsg = "Health Check Command " + errMsg

		pkg.Logger.Println(errMsg)
		log.Println(errMsg)
	} else {
		pkg.Logger.Println("Healthy")
		log.Println("Healthy")
	}
}
