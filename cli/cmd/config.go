package cmd

import (
	"fmt"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/term"
)

func init() {
	rootCmd.AddCommand(configCmd)
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure the application",
	Long:  "Set the username and password for authentication",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Username:")

		var username string
		fmt.Scanln(&username)

		fmt.Println("Password:")

		bytePassword, err := term.ReadPassword(int(syscall.Stdin))
		cobra.CheckErr(err)
		password := string(bytePassword)

		viper.Set("username", username)
		viper.Set("password", password)
		viper.Set("token", "")

		cobra.CheckErr(viper.WriteConfig())
	},
}
