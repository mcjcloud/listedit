/*
Copyright © 2024 Brayden Cloud <brayden14cloud@gmail.com>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// binSearchCmd represents the binSearch command
var binSearchCmd = &cobra.Command{
	Use:   "bin-search <filename> <item>",
	Short: "Perform binary search in file for at item",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return fmt.Errorf("Usage: listedit bin-search <filename> <item>")
		}

		fileItems, err := ReadList(args[0])
		if err != nil {
			return fmt.Errorf("error reading file: %s", err)
		}
		item := args[1]

		return nil
	},
}

func init() {
	rootCmd.AddCommand(binSearchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// binSearchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// binSearchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
