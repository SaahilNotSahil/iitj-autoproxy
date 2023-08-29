package cmd

import (
	"fmt"

	"github.com/XanderWatson/iitj-autoproxy/pkg/cli"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(healthCheckCmd)
}

var healthCheckCmd = &cobra.Command{
	Use:   "hc",
	Short: "Check connection with daemon",
	Long:  "Check connection with daemon",
	Run: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(cli.SendCommandToDaemon("hc"))
		fmt.Println(cli.CreateNamedPipeAndReceiveMessage(true))
	},
}
