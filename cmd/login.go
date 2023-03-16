package cmd

import (
	"log"

	"github.com/XanderWatson/iitj-autoproxy/pkg"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(loginCmd)
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to your firewall authentication",
	Long:  "Login to your firewall authentication",
	Run: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(viper.ReadInConfig())

		username := viper.GetString("username")
		password := viper.GetString("password")

		if username == "" || password == "" {
			pkg.Logger.Println("Please configure the application using the config command")
			log.Println("Please configure the application using the config command")
		}

		pkg.Scheduler.Clear()

		go pkg.RunLoginScheduler(username, password)
		select {}
	},
}
