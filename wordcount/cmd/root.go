/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/achal1304/One2N_GoBootcamp/wordcount/contract"
	"github.com/achal1304/One2N_GoBootcamp/wordcount/wcerrors"
	"github.com/achal1304/One2N_GoBootcamp/wordcount/wchandler"
	"github.com/spf13/cobra"
)

var (
	flagsOptions = contract.WcFlags{}
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wc [flags] [file]",
	Short: "Perform word count operations",
	Args:  cobra.ExactArgs(1), // Ensure exactly one file is provided
	RunE: func(cmd *cobra.Command, args []string) error {
		fileName := args[0]
		// Process file based on the flags
		wcValues, err := wchandler.ProcessWCCommand(fileName, flagsOptions)
		if err != nil {
			return wcerrors.HandleErrors(err, fileName)
		}

		wchandler.PrintStdOut(wchandler.GenerateOutput(wcValues, flagsOptions))
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
	flagsOptions = contract.NewFlags()
	rootCmd.Flags().BoolVarP(&flagsOptions.LineCount, "lines", "l", false, "count lines in the file")
	rootCmd.Flags().BoolVarP(&flagsOptions.WordCount, "word", "w", false, "count words in the file")

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.One2N_GoBootcamp.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
