/*
Copyright © 2024 Brayden Cloud <brayden14cloud@gmail.com>

*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new <filename> [items]...",
	Short: "Create a new list",
	Long: `Create a new list at the specified location. If the file already exists, it will be overwritten.
If an input file and items are provided, the items will be added after the input file content.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("No filename provided")
		}

		items := ProcessList(args[1:])

		filename := args[0]
		newFile, err := os.Create(filename)
		if err != nil {
			return fmt.Errorf("Error creating file: %s", err)
		}
		defer newFile.Close()

		n, toWrite := CombineLists(InputFileContent, items)

		if _, err := newFile.WriteString(strings.Join(toWrite, "\n")); err != nil {
			return fmt.Errorf("Error writing to file: %s", err)
		}

		if absolutePath, err := filepath.Abs(filename); err == nil {
			fmt.Println("Created new list at", absolutePath, "with", n, "items")
		} else {
			return fmt.Errorf("Error getting absolute path: %s", err)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
