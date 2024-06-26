package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// camerasCmd represents the cameras command
var camerasCmd = &cobra.Command{
	Use:   "cameras",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cameras called")
	},
}

func init() {
	rootCmd.AddCommand(camerasCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// camerasCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// camerasCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
