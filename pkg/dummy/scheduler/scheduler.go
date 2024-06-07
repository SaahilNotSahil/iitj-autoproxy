package scheduler

import (
	"time"

	"github.com/SaahilNotSahil/iitj-autoproxy/pkg"
)

var SchedulerRunning = false

func ping(pingChannel chan bool, kill chan bool) {
	for {
		select {
		case <-kill:
			return
		default:
			internet := pkg.InternetAvailable()

			if !internet {
				pingChannel <- internet
			}

			time.Sleep(5 * time.Second)
		}
	}
}

func login(
	pingChannel chan bool,
	logChannel chan string,
	kill chan bool,
) {
	for {
		select {
		case <-kill:
			return
		case <-pingChannel:
			logChannel <- "Logged in successfully"
		}
	}
}

func logToFile(logChannel chan string, kill chan bool) {
	for {
		select {
		case <-kill:
			return
		case <-logChannel:
			pkg.Logger.Printf("No internet; %s", <-logChannel)
		}
	}
}

func RunLoginScheduler(kill chan bool) bool {
	if SchedulerRunning {
		return false
	}

	SchedulerRunning = true

	pingChannel := make(chan bool)
	logChannel := make(chan string)

	go ping(pingChannel, kill)
	go login(pingChannel, logChannel, kill)
	go logToFile(logChannel, kill)

	return true
}
