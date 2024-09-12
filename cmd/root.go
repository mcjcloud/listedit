/*
Copyright © 2024 Brayden Cloud <brayden14cloud@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile, filePath string

var InputFileContent []string
var Force, Sort, Dedup bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "listedit <sort|add-to|remove-from|find-in> <filename> [args]...",
	Short: "Edit a list of items in a text file",
	Long: `Edit a list of items in a text file. Support for adding, removing, and sorting items in a list.
Example usage:
	- listedit new <filename> "item 1" "item 2" # creates a new file with items
	- listedit sort <filename> # sorts in place
	- listedit add-to <filename> "item 1" "item 2" "item 3" # adds items to the list if they are not already present (use --force to add if present)
	- listedit remove-from <filename> "item 1" "item 2" # removes all instances of items
	- listedit find-in <filename> "item" # finds all instances of item in the list`,
	// Run: func(cmd *cobra.Command, args []string) {
	// },
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if filePath == "" {
			return nil
		}

		var err error
		InputFileContent, err = ReadAndProcessList(filePath)
		if err != nil {
			return fmt.Errorf("Error reading input file: %s", err)
		}

		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	DedupLookup = make(map[string]bool)
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&filePath, "file", "", "input file")
	rootCmd.PersistentFlags().BoolVar(&Force, "force", false, "force action")
	rootCmd.PersistentFlags().BoolVar(&Sort, "sort", false, "sort the resulting list")
	rootCmd.PersistentFlags().BoolVar(&Dedup, "dedup", false, "remove duplicates")

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.listedit.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".listedit" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".listedit")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
