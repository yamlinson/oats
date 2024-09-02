// Package cmd implements application subcommands
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yamlinson/oats/internal/db"
)

var (
	rmCurrent bool
	// rmCmd represents the rm command
	rmCmd = &cobra.Command{
		Use:   "rm [\"list name\"] [\"item name\"]",
		Short: "Removes an item from a list",
		Args: func(cmd *cobra.Command, args []string) error {
			if rmCurrent {
				if err := cobra.ExactArgs(0)(cmd, args); err != nil {
					return err
				}
			} else {
				if err := cobra.ExactArgs(2)(cmd, args); err != nil {
					return err
				}
			}
			return nil
		},
		Run: func(_ *cobra.Command, args []string) {
			if rmCurrent {
				item, err := db.GetCurrent()
				if err != nil {
					panic(err)
				}
				err = db.RemoveItem(item.Name, item.List)
				if err != nil {
					panic(err)
				}
				fmt.Printf("%s removed from %s\n", item.Name, item.List)
			} else {
				err := db.RemoveItem(args[1], args[0])
				if err != nil {
					panic(err)
				}
				fmt.Printf("%s removed from %s\n", args[1], args[0])
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(rmCmd)
	rmCmd.Flags().BoolVarP(&rmCurrent, "current", "c", false, "remove item marked current")
}
