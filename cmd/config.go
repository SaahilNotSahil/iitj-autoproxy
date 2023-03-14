package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		var password string
		fmt.Scanln(&password)

		viper.Set("username", username)
		viper.Set("password", password)
		viper.Set("token", "")

		cobra.CheckErr(viper.WriteConfig())
	},
}
