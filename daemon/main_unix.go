//go:build linux || darwin
// +build linux darwin

package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"path"
	"runtime"
	"syscall"

	"github.com/spf13/viper"

	"github.com/SaahilNotSahil/iitj-autoproxy/daemon/commands"
	"github.com/SaahilNotSahil/iitj-autoproxy/pkg"
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

	scheduler_running_state := viper.GetBool("scheduler_running_state")

	if scheduler_running_state {
		commands.ScheduleCmd()
	}

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

	var baseConfigPath string

	if runtime.GOOS == "darwin" {
		baseConfigPath = "/opt/homebrew/etc/iitj-autoproxy/autoproxy.config"
	} else {
		baseConfigPath = "/etc/iitj-autoproxy/autoproxy.config"
	}

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
