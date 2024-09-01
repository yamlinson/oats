// Package cmd implements application subcommands
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yamlinson/oats/internal/db"
)

var (
	all       bool
	allInList bool
	last      bool
	first     bool
	random    bool
	randomAll bool
	// getCmd represents the get command
	getCmd = &cobra.Command{
		Use:   "get",
		Short: "Get lists and their items",
		Args: func(cmd *cobra.Command, args []string) error {
			first = !all && !allInList && !last && !random && !randomAll
			if allInList {
				if err := cobra.MaximumNArgs(1)(cmd, args); err != nil {
					return err
				}
			}
			if first || last || random {
				if err := cobra.ExactArgs(1)(cmd, args); err != nil {
					return err
				}
			}
			if all || randomAll {
				if err := cobra.ExactArgs(0)(cmd, args); err != nil {
					return err
				}
			}
			return nil
		},
		Run: func(_ *cobra.Command, args []string) {
			first = !all && !allInList && !last && !random && !randomAll
			if all {
				items, err := db.GetAllItems()
				if err != nil {
					panic(err)
				}
				for _, element := range *items {
					fmt.Printf("%s: %s\n", element.List, element.Name)
				}
			}
			if allInList {
				switch len(args) {
				case 0:
					// Return all lists
					lists, err := db.GetLists()
					if err != nil {
						panic(err)
					}
					for _, element := range *lists {
						fmt.Println(element)
					}
				case 1:
					// Return all items of list
					items, err := db.GetAllItemsByList(args[0])
					if err != nil {
						panic(err)
					}
					for _, element := range *items {
						fmt.Printf("%s: %s\n", element.List, element.Name)
					}
				}
			}
			if first {
				item, err := db.GetItem(args[0], false, false)
				if err != nil {
					panic(err)
				}
				fmt.Printf("%s: %s\n", item.List, item.Name)
			}
			if last {
				item, err := db.GetItem(args[0], false, true)
				if err != nil {
					panic(err)
				}
				fmt.Printf("%s: %s\n", item.List, item.Name)
			}
			if random {
				fmt.Println("Return random")
			}
			if randomAll {
				fmt.Println("Random any")
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.Flags().BoolVarP(&all, "all", "A", false, "get all items from all lists")
	getCmd.Flags().BoolVarP(&allInList, "all-in-list", "a", false, "get all items from a given list, or all lists if none is supplied")
	getCmd.Flags().BoolVarP(&last, "last", "l", false, "get the most recently created item instead of the oldest")
	getCmd.Flags().BoolVarP(&random, "random", "r", false, "get a random item from the specified list")
	getCmd.Flags().BoolVarP(&randomAll, "any-random", "R", false, "get a random item from any list")
	getCmd.MarkFlagsMutuallyExclusive("all", "all-in-list", "last", "random", "any-random")
}
