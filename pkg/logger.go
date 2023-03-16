package pkg

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var Logger *log.Logger

func init() {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)
	file, err := os.OpenFile(home+"/.autoproxy.logs", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	cobra.CheckErr(err)
	Logger = log.New(file, "", log.LstdFlags)
}
