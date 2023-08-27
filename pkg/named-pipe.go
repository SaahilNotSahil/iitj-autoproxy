package pkg

import (
	"log"
	"os"
)

func GetCTDNamedPipe() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	return homeDir + "/autoproxy-ctd"
}
