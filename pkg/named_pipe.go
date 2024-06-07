package pkg

import (
	"log"
	"os"
	"runtime"
)

func GetCTDNamedPipe() string {
	var pipeName string

	if runtime.GOOS == "windows" {
		pipeName = `\\.\pipe\autoproxy-ctd`
	} else {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			Logger.Println(err)
			log.Fatal(err)
		}

		pipeName = homeDir + "/autoproxy-ctd"
	}

	return pipeName
}

func GetDTCNamedPipe() string {
	var pipeName string

	if runtime.GOOS == "windows" {
		pipeName = `\\.\pipe\autoproxy-dtc`
	} else {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			Logger.Println(err)
			log.Fatal(err)
		}

		pipeName = homeDir + "/autoproxy-dtc"
	}

	return pipeName
}
