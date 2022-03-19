/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cli

import (
	"fmt"
	"oh-my-posh/environment"
	"time"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get [shell|cache-path|millies]",
	Short: "Get a value from oh-my-posh",
	Long: `Get a value from oh-my-posh.
This command is used to get the value of the following variables:

- shell
- millis`,
	ValidArgs: []string{
		"millis",
		"shell",
	},
	Args: cobra.OnlyValidArgs,
	Run: func(cmd *cobra.Command, args []string) {
		env := &environment.ShellEnvironment{
			Version: cliVersion,
		}
		env.Init(false)
		defer env.Close()
		switch args[0] {
		case "millis":
			fmt.Print(time.Now().UnixNano() / 1000000)
		case "shell":
			fmt.Println(env.Shell())
			// case "cache-path":
			// 	fmt.Print(env.CachePath())
		}
	},
}

func init() { // nolint:gochecknoinits
	rootCmd.AddCommand(getCmd)
}