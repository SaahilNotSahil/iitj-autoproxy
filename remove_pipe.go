package main

import "os"

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	pipePath := homeDir + "/autoproxy-ctd"

	os.Remove(pipePath)
}
