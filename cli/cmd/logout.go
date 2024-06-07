package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/SaahilNotSahil/iitj-autoproxy/pkg/cli"
)

func init() {
	logoutCmd.Flags().BoolVarP(
		&isDummy, "dummy", "d", false, "Run a dummy version of the command",
	)
	rootCmd.AddCommand(logoutCmd)
}

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout of your firewall authentication",
	Long:  "Logout of your firewall authentication",
	Run: func(cmd *cobra.Command, args []string) {
		if isDummy {
			fmt.Println("Dummy logout invoked")

			cobra.CheckErr(cli.SendCommandToDaemon("logoutDummy"))
			fmt.Println(cli.CreateNamedPipeAndReceiveMessage(false))
		} else {
			cobra.CheckErr(cli.SendCommandToDaemon("logout"))
			fmt.Println(cli.CreateNamedPipeAndReceiveMessage(false))
		}
	},
}
