package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/SaahilNotSahil/iitj-autoproxy/pkg/cli"
)

func init() {
	loginCmd.Flags().BoolVarP(&isDummy, "dummy", "d", false, "")
	rootCmd.AddCommand(loginCmd)
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to your firewall authentication",
	Long:  "Login to your firewall authentication",
	Run: func(cmd *cobra.Command, args []string) {
		if isDummy {
			fmt.Println("Dummy login invoked")

			cobra.CheckErr(cli.SendCommandToDaemon("loginDummy"))
			fmt.Println(cli.CreateNamedPipeAndReceiveMessage(false))
		} else {
			cobra.CheckErr(cli.SendCommandToDaemon("login"))
			fmt.Println(cli.CreateNamedPipeAndReceiveMessage(false))
		}
	},
}
