package cli

import (
	"log"
	"os"
)

func SendCommandToDaemon(command string) error {
	pipe, err := os.OpenFile("autoproxy", os.O_WRONLY, os.ModeNamedPipe)
	if err != nil {
		log.Fatal("The autoproxy daemon is not running.")
	}
	defer pipe.Close()

	_, err = pipe.WriteString(command)

	return err
}
