package cmd

import (
	"github.com/XanderWatson/iitj-autoproxy/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

func init() {
	rootCmd.AddCommand(logoutCmd)
}

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout of your firewall authentication",
	Long:  "Logout of your firewall authentication",
	Run: func(cmd *cobra.Command, args []string) {
		token := viper.GetString("token")

		if token == "" {
			pkg.Logger.Println("User not logged in")
			log.Println("User not logged in")
		}

		cobra.CheckErr(pkg.Logout(token))

		pkg.Logger.Println("Logged out successfully")
		log.Println("Logged out successfully")
	},
}
