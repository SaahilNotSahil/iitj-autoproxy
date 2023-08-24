package cmd

import (
	"fmt"
	"syscall"

	"github.com/spf13/cobra"
	"golang.org/x/term"

	"github.com/XanderWatson/iitj-autoproxy/pkg/keystore"
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

		cobra.CheckErr(keystore.Set("username", username))
		cobra.CheckErr(keystore.Set("password", password))
		cobra.CheckErr(keystore.Set("token", ""))
	},
}
