package pkg

import (
	"log"
	"os"
	"runtime"
)

func GetCTDNamedPipe() string {
	var pipeName string

	homeDir, err := os.UserHomeDir()
	if err != nil {
		Logger.Println(err)
		log.Fatal(err)
	}

	if runtime.GOOS == "windows" {
		pipeName = `\\.\pipe\autoproxy-ctd`
	} else {
		pipeName = homeDir + "/autoproxy-ctd"
	}

	return pipeName
}

func GetDTCNamedPipe() string {
	var pipeName string

	homeDir, err := os.UserHomeDir()
	if err != nil {
		Logger.Println(err)
		log.Fatal(err)
	}

	if runtime.GOOS == "windows" {
		pipeName = `\\.\pipe\autoproxy-dtc`
	} else {
		pipeName = homeDir + "/autoproxy-dtc"
	}

	return pipeName
}
