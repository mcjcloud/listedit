/*
Copyright © 2024 Brayden Cloud <brayden14cloud@gmail.com>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// isSortedCmd represents the isSorted command
var isSortedCmd = &cobra.Command{
	Use:   "is-sorted",
	Short: "Checks if the lines of a file are sorted",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("no filename provided")
		}

		fileItems, err := ReadList(args[0])
		if err != nil {
			return fmt.Errorf("error reading file: %s", err)
		}

		if IsSorted(fileItems) {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
			os.Exit(1)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(isSortedCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// isSortedCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// isSortedCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
