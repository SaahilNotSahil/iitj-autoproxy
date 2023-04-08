package cmd

import (
	"github.com/XanderWatson/iitj-autoproxy/pkg/cli"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(scheduleCmd)
}

var scheduleCmd = &cobra.Command{
	Use:   "schedule",
	Short: "Schedule your firewall authentication",
	Long:  "Schedule your firewall authentication",
	Run: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(cli.SendCommandToDaemon("schedule"))
	},
}
