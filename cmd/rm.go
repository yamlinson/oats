// Package cmd implements application subcommands
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yamlinson/oats/internal/db"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm [\"list name\"] [\"item name\"]",
	Short: "Removes an item from a list",
	Args:  cobra.ExactArgs(2),
	Run: func(_ *cobra.Command, args []string) {
		err := db.RemoveItem(args[1], args[0])
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s removed from %s\n", args[1], args[0])
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
