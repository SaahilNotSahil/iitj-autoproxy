package cmd

import (
	"fmt"

	"github.com/XanderWatson/iitj-autoproxy/pkg/cli"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(loginCmd)
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to your firewall authentication",
	Long:  "Login to your firewall authentication",
	Run: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(cli.SendCommandToDaemon("login"))
		fmt.Println(cli.CreateNamedPipeAndReceiveMessage(false))
	},
}
