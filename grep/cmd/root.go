/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/achal1304/One2N_GoBootcamp/grep/contract"
	"github.com/achal1304/One2N_GoBootcamp/grep/handler"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "grep",
	Short: "filter out by searching texts",
	RunE: func(cmd *cobra.Command, args []string) error {
		req := contract.GrepRequest{}
		if len(args) >= 2 {
			req.SearchString = []byte(args[0])
			req.FileName = args[1]
			response, err := handler.ProcessGrepRequest(req)
			if err != nil {
				handler.PrintStdOut(os.Stderr, err.Error())
				os.Exit(1)
			}

			handler.PrintResponseStdOut(os.Stdout, response)
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
}
