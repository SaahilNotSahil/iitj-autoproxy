package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/SaahilNotSahil/iitj-autoproxy/pkg/cli"
)

func init() {
	scheduleCmd.Flags().BoolVarP(
		&isDummy, "dummy", "d", false, "Run a dummy version of the command",
	)
	rootCmd.AddCommand(scheduleCmd)
}

var scheduleCmd = &cobra.Command{
	Use:   "schedule",
	Short: "Schedule your firewall authentication",
	Long:  "Schedule your firewall authentication",
	Run: func(cmd *cobra.Command, args []string) {
		if isDummy {
			fmt.Println("Dummy schedule invoked")

			cobra.CheckErr(cli.SendCommandToDaemon("scheduleDummy"))
			fmt.Println(cli.CreateNamedPipeAndReceiveMessage(false))
		} else {
			cobra.CheckErr(cli.SendCommandToDaemon("schedule"))
			fmt.Println(cli.CreateNamedPipeAndReceiveMessage(false))
		}
	},
}
