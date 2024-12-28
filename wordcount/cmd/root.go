/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/achal1304/One2N_GoBootcamp/wordcount/wcerrors"
	"github.com/achal1304/One2N_GoBootcamp/wordcount/wchandler"
	"github.com/spf13/cobra"
)

type WcFlags struct {
	lineCount bool
}

var (
	flagsOptions WcFlags
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wc [flags] [file]",
	Short: "Perform word count operations",
	Args:  cobra.ExactArgs(1), // Ensure exactly one file is provided
	RunE: func(cmd *cobra.Command, args []string) error {
		fileName := args[0]

		// Process file based on the flags
		if flagsOptions.lineCount {
			count, err := wchandler.ProcessWCCommand(fileName)
			if err != nil {
				return wcerrors.HandleErrors(err, fileName)
			}

			fmt.Printf("%8d %s\n", count, fileName)
		}
		return nil
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

func init() {
	rootCmd.SilenceUsage = true
	rootCmd.Flags().BoolVarP(&flagsOptions.lineCount, "lines", "l", false, "count lines in the file")

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.One2N_GoBootcamp.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
