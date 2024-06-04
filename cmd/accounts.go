package cmd

import (
	"fmt"

	"github.com/juiroy/ucams-cli/pkg/ucams"
	"github.com/spf13/cobra"
)

func addAccountRunE(cmd *cobra.Command, args []string) error {
	username, _ := cmd.Flags().GetString("username")
	password, _ := cmd.Flags().GetString("password")

	client := ucams.Create()

	if err := client.AddAccount(username, password); err != nil {
		fmt.Println(err)
	}

	return nil
}

var accountsCommand = &cobra.Command{
	Use: "accounts",
}

var addAccountCommand = &cobra.Command{
	Use:  "add",
	RunE: addAccountRunE,
}

func init() {
	addAccountCommand.Flags().StringP("username", "u", "", "Account username")
	addAccountCommand.Flags().StringP("password", "p", "", "Account password")

	addAccountCommand.MarkFlagRequired("username")
	addAccountCommand.MarkFlagRequired("password")

	accountsCommand.AddCommand(addAccountCommand)
	rootCmd.AddCommand(accountsCommand)
}
