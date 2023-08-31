package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/XanderWatson/iitj-autoproxy/pkg"
)

func init() {
	rootCmd.AddCommand(cleanupCmd)
}

var cleanupCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "Clean up named pipes",
	Long:  "Clean up named pipes",
	Run: func(cmd *cobra.Command, args []string) {
		pipePath := pkg.GetCTDNamedPipe()
		os.Remove(pipePath)

		pipePath = pkg.GetDTCNamedPipe()
		os.Remove(pipePath)
	},
}
