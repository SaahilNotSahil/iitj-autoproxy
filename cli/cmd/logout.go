package cmd

import (
	"github.com/XanderWatson/iitj-autoproxy/pkg/cli"
	
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(logoutCmd)
}

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout of your firewall authentication",
	Long:  "Logout of your firewall authentication",
	Run: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(cli.SendCommandToDaemon("logout"))
	},
}
