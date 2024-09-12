/*
Copyright © 2024 Brayden Cloud <brayden14cloud@gmail.com>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// addToCmd represents the add command
var addToCmd = &cobra.Command{
	Use:   "add-to <filename> <item>...",
	Short: "Add items to a list in a text file",
	Long: `Add items to a list in a text file. If the list is sorted, the items will be added in sorted order.
Exmaple usage:
	- listedit add-to <filename> "item 1" "item 2" "item 3" # adds items to the list if they are not already present (use --force to add if present)
	- listedit add-to <filename> --file <input-filename> # add all lines from input-filename to the list`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("no filename provided")
		}

		filename := args[0]
		items := ProcessList(args[1:])
		fileItems, err := ReadAndProcessList(filename)
		if err != nil {
			return fmt.Errorf("Error reading file: %s", err)
		}

		toWrite := make([]string, len(fileItems))
		copy(toWrite, fileItems)

		if IsSorted(fileItems) {
			Sort = true
		}

		var totalAdded, n int
		n, toWrite = CombineLists(toWrite, InputFileContent)
		totalAdded += n
		n, toWrite = CombineLists(toWrite, items)
		totalAdded += n

		WriteList(filename, toWrite)

		fmt.Printf("Added %d new lines\n", totalAdded)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addToCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
