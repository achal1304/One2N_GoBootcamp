/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"sync"

	"github.com/achal1304/One2N_GoBootcamp/wordcount/contract"
	"github.com/achal1304/One2N_GoBootcamp/wordcount/wchandler"
	"github.com/spf13/cobra"
)

var (
	flagsOptions = contract.WcFlags{}
)

const MaxFiles = 10

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wc [flags] [file]",
	Short: "Perform word count operations",
	Args:  cobra.MaximumNArgs(MaxFiles), // Ensure exactly one file is provided
	RunE: func(cmd *cobra.Command, args []string) error {
		// fileName := args[0]
		// Process file based on the flags
		total := contract.WcValues{FileName: "total"}
		wcValuesCh := make(chan contract.WcValues)
		exitCode := make(chan int)
		defer close(exitCode)

		done := make(chan struct{})
		var wg sync.WaitGroup
		for _, arg := range args {
			wg.Add(1)
			go wchandler.ProcessWCCommand(&wg, arg, flagsOptions, wcValuesCh)
		}

		go func() {
			exitStatusCode := wchandler.ComputeTotalCount(len(args) > 1, wcValuesCh, flagsOptions, &total, done, os.Stdout)
			exitCode <- exitStatusCode
		}()

		wg.Wait()
		close(wcValuesCh)
		<-done
		exitCodeValue := <-exitCode
		if exitCodeValue != 0 {
			os.Exit(exitCodeValue)
		}
		if len(args) > 1 {
			wchandler.PrintStdOut(os.Stdout, wchandler.GenerateOutput(total, flagsOptions))
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
	flagsOptions = contract.NewFlags()
	rootCmd.Flags().BoolVarP(&flagsOptions.LineCount, "lines", "l", false, "count lines in the file")
	rootCmd.Flags().BoolVarP(&flagsOptions.WordCount, "word", "w", false, "count words in the file")
	rootCmd.Flags().BoolVarP(&flagsOptions.CharacterCount, "character", "c", false, "count characters in the file")

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.One2N_GoBootcamp.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
