//go:build windows
// +build windows

package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"path"
	"syscall"

	"github.com/spf13/viper"
	"gopkg.in/natefinch/npipe.v2"

	"github.com/SaahilNotSahil/iitj-autoproxy/daemon/commands"
	"github.com/SaahilNotSahil/iitj-autoproxy/pkg"
)

func main() {
	initConfig()
	cleanup()

	pipePath := pkg.GetCTDNamedPipe()

	ln, err := npipe.Listen(pipePath)
	if err != nil {
		pkg.Logger.Println(err)
		log.Fatal(err)
	}

	scheduler_running_state := viper.GetBool("scheduler_running_state")

	if scheduler_running_state {
		commands.ScheduleCmd()
	}

	buf := make([]byte, 1024)
	for {
		conn, err := ln.Accept()
		if err != nil {
			pkg.Logger.Println(err)
			log.Fatal(err)
		}

		num_bytes, err := conn.Read(buf)
		if err != nil {
			pkg.Logger.Println(err)
			log.Fatal(err)
		}

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
	case "loginDummy":
		commands.LoginDummyCmd()
	case "logout":
		commands.LogoutCmd()
	case "logoutDummy":
		commands.LogoutDummyCmd()
	case "schedule":
		commands.ScheduleCmd()
	case "scheduleDummy":
		commands.ScheduleDummyCmd()
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

	baseConfigPath := "C:/ProgramData/IITJ Autoproxy/autoproxy.config"

	configName := ".autoproxyconfig"

	targetConfig := path.Join(home, configName)

	viper.AddConfigPath(home)
	viper.SetConfigName(configName)
	viper.SetConfigType("json")

	_, err = os.Stat(targetConfig)
	if err != nil {
		pkg.Logger.Println(err)
		log.Println(err)

		_, err = copy(baseConfigPath, targetConfig)
		if err != nil {
			pkg.Logger.Println(err)
			log.Fatal(err)
		}
	}

	err = viper.ReadInConfig()
	if err != nil {
		pkg.Logger.Println(err)
		log.Fatal(err)
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
		return 0, fmt.Errorf("failed to stat source file %s: %w", src, err)
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, fmt.Errorf("failed to open source file %s: %w", src, err)
	}
	defer source.Close()

	destination, err := os.OpenFile(dst, os.O_RDWR|os.O_CREATE|os.O_TRUNC, sourceFileStat.Mode())
	if err != nil {
		return 0, fmt.Errorf("failed to create destination file %s: %w", dst, err)
	}
	defer destination.Close()

	bytesCopied, err := io.Copy(destination, source)
	if err != nil {
		return 0, fmt.Errorf("failed to copy data from %s to %s: %w", src, dst, err)
	}

	return bytesCopied, nil
}
