package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var isDummy bool

func initConfig() {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	viper.AddConfigPath(home)

	viper.SetConfigName(".autoproxyconfig")
	viper.SetConfigType("json")

	cobra.CheckErr(viper.ReadInConfig())
}

func init() {
	cobra.OnInitialize(initConfig)
}

var rootCmd = &cobra.Command{
	Use:   "autoproxy",
	Short: "Use IITJ internet hassle-free",
	Long: `A Fast CLI Autoproxy Tool built for IITJ fraternity. 
	
			With love, by SaahilNotSahil.`,
	Run: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(cmd.Help())
	},
}

func Execute() {
	err := rootCmd.Execute()
	cobra.CheckErr(err)
}
