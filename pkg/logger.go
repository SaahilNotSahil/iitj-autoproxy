package pkg

import (
	"bufio"
	"log"
	"os"
)

var Logger *log.Logger
var Writer *bufio.Writer

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Println(err)

		os.Exit(1)
	}

	file, err := os.OpenFile(
		home+"/.autoproxy.logs", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644,
	)
	if err != nil {
		log.Println(err)

		os.Exit(1)
	}

	Logger = log.New(file, "", log.LstdFlags)
	Writer = bufio.NewWriter(file)

	// Logger.SetOutput(Writer)
}
