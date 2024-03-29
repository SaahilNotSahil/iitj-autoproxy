package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "autoproxy",
	Short: "Use IITJ internet hassle-free",
	Long: `A Fast CLI Autoproxy Tool built for IITJ fraternity. 
	
			With love, by XanderWatson.`,
	Run: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(cmd.Help())
	},
}

func Execute() {
	err := rootCmd.Execute()
	cobra.CheckErr(err)
}
