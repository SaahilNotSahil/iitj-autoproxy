//go:build linux || darwin
// +build linux darwin

package cli

import (
	"log"
	"os"
	"syscall"

	"github.com/XanderWatson/iitj-autoproxy/pkg"
)

func CreateNamedPipeAndReceiveMessage(isHealthCheck bool) string {
	pipeName := pkg.GetDTCNamedPipe()

	err := syscall.Mkfifo(pipeName, 0666)
	if err != nil {
		if isHealthCheck {
			return "Unhealthy"
		}
		
		pkg.Logger.Println(err)
		log.Fatal(err)
	}
	defer os.Remove(pipeName)

	pipe, err := os.OpenFile(pipename, os.O_RDONLY, os.ModeNamedPipe)
	if err != nil {
		if isHealthCheck {
			return "Unhealthy"
		}

		pkg.Logger.Println(err)
		log.Fatal(err)
	}
	defer pipe.Close()

	buf := make([]byte, 1024)

	num_bytes, _ := pipe.Read(buf)

	message := string(buf[:num_bytes])

	return message
}
