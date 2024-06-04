package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "ucams-cli",
	Long: `ucams-cli should be an cli client for UCAMS VMS
	But first of all it designed as quick and simple utility for obtaining camera feed URL from UCAMS
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	var cfgFile string
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ucams-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringP("host", "s", "https://ucams.ufanet.ru", "UCAMS API host")
	rootCmd.PersistentFlags().StringP("username", "u", "", "Account username")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Activate for verbose")

	_ = rootCmd.MarkPersistentFlagRequired("username")
}
