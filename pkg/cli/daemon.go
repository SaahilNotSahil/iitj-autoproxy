package cli

import (
	"log"
	"os"

	"github.com/SaahilNotSahil/iitj-autoproxy/pkg"
)

func SendCommandToDaemon(command string) error {
	pipe, err := os.OpenFile(
		pkg.GetCTDNamedPipe(), os.O_WRONLY, os.ModeNamedPipe,
	)
	if err != nil {
		pkg.Logger.Println(err)
		log.Fatal("The autoproxy daemon is not running.")
	}
	defer pipe.Close()

	_, err = pipe.WriteString(command)

	return err
}
