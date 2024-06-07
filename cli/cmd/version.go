package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/SaahilNotSahil/iitj-autoproxy/pkg"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "IITJ Autoproxy version",
	Long:  `IITJ Autoproxy version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("IITJ Autoproxy Version - v%s\n", pkg.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
