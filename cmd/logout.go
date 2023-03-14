package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func init() {
	rootCmd.AddCommand(logoutCmd)
}

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		token := viper.GetString("token")

		if token == "" {
			cobra.CheckErr("User not logged in")
		}

		cobra.CheckErr(Logout(token))
		log.Println("Logged out successfully")
	},
}

func Logout(token string) error {
	url := viper.GetString("login_base_url") + "logout?" + token

	_, err := http.Get(url)

	return err
}
