// Package cmd implements application subcommands
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/yamlinson/oats/internal/db"
)

var (
	create bool
	// addCmd represents the add command
	addCmd = &cobra.Command{
		Use:   "add [\"list name\"] [\"item name\"]",
		Short: "Adds an item to a list",
		Args:  cobra.ExactArgs(2),
		Run: func(_ *cobra.Command, args []string) {
			item := db.Item{
				Name:    args[1],
				List:    args[0],
				Created: time.Now(),
			}
			err := db.AddItem(item)
			if err != nil {
				panic(err)
			}
			fmt.Printf("%s added to %s\n", item.Name, item.List)
		},
	}
)

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().BoolVarP(&create, "create", "c", false, "create list if it does not exist")
}
