// Package cmd implements application subcommands
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "oats",
	Short: "One-at-a-time To-do's",
	Long:  "Oats is a todo list manager which only presents one task at a time",
	Run: func(cmd *cobra.Command, _ []string) {
		cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
