package daemon

import (
	"log"
	"os"

	"github.com/XanderWatson/iitj-autoproxy/pkg"
)

func SendMessageToCLI(message string) error {
	pipename := pkg.GetDTCNamedPipe()

	pipe, err := os.OpenFile(pipename, os.O_WRONLY, os.ModeNamedPipe)
	if err != nil {
		pkg.Logger.Println(err)
		log.Println(err)

		return err
	}
	defer pipe.Close()

	_, err = pipe.WriteString(message)

	return err
}
