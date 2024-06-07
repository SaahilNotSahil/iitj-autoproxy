//go:build windows
// +build windows

package cli

import (
	"log"

	"gopkg.in/natefinch/npipe.v2"

	"github.com/SaahilNotSahil/iitj-autoproxy/pkg"
)

func CreateNamedPipeAndReceiveMessage(isHealthCheck bool) string {
	pipename := pkg.GetDTCNamedPipe()

	ln, err := npipe.Listen(pipename)
	if err != nil {
		if isHealthCheck {
			return "Unhealthy"
		}

		pkg.Logger.Println(err)
		log.Fatal(err)
	}

	conn, err := ln.Accept()
	if err != nil {
		if isHealthCheck {
			return "Unhealthy"
		}

		pkg.Logger.Println(err)
		log.Fatal(err)
	}
	defer conn.Close()

	buf := make([]byte, 1024)

	num_bytes, _ := conn.Read(buf)

	message := string(buf[:num_bytes])

	return message
}
