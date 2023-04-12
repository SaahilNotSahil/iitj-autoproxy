package main

import (
	"github.com/XanderWatson/iitj-autoproxy/daemon/commands"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	initConfig()
	cleanup()

	err := syscall.Mkfifo("autoproxy-ctd", 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove("autoproxy-ctd")

	pipe, err := os.OpenFile("autoproxy-ctd", os.O_RDONLY, os.ModeNamedPipe)
	if err != nil {
		log.Fatal(err)
	}
	defer pipe.Close()

	buf := make([]byte, 1024)
	for {
		num_bytes, _ := pipe.Read(buf)
		command := string(buf[:num_bytes])

		if num_bytes > 0 {
			log.Println("Running command: ", command)
			execute(command)
		}
	}
}

func execute(command string) {
	switch command {
	case "login":
		commands.LoginCmd()
	case "logout":
		commands.LogoutCmd()
	case "schedule":
		commands.ScheduleCmd()
	}
}

func initConfig() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	viper.AddConfigPath(home)

	viper.SetConfigType("json")
	viper.SetConfigName(".autoproxy.config")

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func cleanup() {
	signal_channel := make(chan os.Signal, 1)
	signal.Notify(signal_channel, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signal_channel

		os.Remove("autoproxy-ctd")
		os.Exit(0)
	}()
}
