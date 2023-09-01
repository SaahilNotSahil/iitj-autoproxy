//go:build linux || darwin
// +build linux darwin

package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"

	"github.com/XanderWatson/iitj-autoproxy/daemon/commands"
	"github.com/XanderWatson/iitj-autoproxy/pkg"
)

func main() {
	initConfig()
	cleanup()

	pipePath := pkg.GetCTDNamedPipe()

	err := syscall.Mkfifo(pipePath, 0666)
	if err != nil {
		pkg.Logger.Println(err)
		log.Fatal(err)
	}
	defer os.Remove(pipePath)

	pipe, err := os.OpenFile(pipePath, os.O_RDONLY, os.ModeNamedPipe)
	if err != nil {
		pkg.Logger.Println(err)
		log.Fatal(err)
	}
	defer pipe.Close()

	buf := make([]byte, 1024)
	for {
		num_bytes, _ := pipe.Read(buf)
		command := string(buf[:num_bytes])

		if num_bytes > 0 {
			pkg.Logger.Println("Running command: ", command)
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
	case "hc":
		commands.HealthCheckCmd()
	}
}

func initConfig() {
	home, err := os.UserHomeDir()
	if err != nil {
		pkg.Logger.Println(err)
		log.Fatal(err)
	}

	baseConfigPath := "/etc/iitj-autoproxy/autoproxy.config"

	configName := ".autoproxy.config"

	viper.AddConfigPath(home)

	viper.SetConfigType("json")
	viper.SetConfigName(configName)

	err = viper.ReadInConfig()
	if err != nil {
		pkg.Logger.Println(err)
		log.Println(err)

		_, err = copy(baseConfigPath, home+"/"+configName)
		if err != nil {
			pkg.Logger.Println(err)
			log.Fatal(err)
		}

		err = viper.ReadInConfig()
		if err != nil {
			pkg.Logger.Println(err)
			log.Fatal(err)
		}
	}
}

func cleanup() {
	signal_channel := make(chan os.Signal, 1)
	signal.Notify(signal_channel, syscall.SIGINT, syscall.SIGTERM)

	pipePath := pkg.GetCTDNamedPipe()

	go func() {
		<-signal_channel

		os.Remove(pipePath)
		os.Exit(0)
	}()
}

func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()

	return io.Copy(destination, source)
}
