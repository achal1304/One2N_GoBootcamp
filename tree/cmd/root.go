/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/achal1304/One2N_GoBootcamp/tree/contract"
	"github.com/achal1304/One2N_GoBootcamp/tree/handler"
	"github.com/spf13/cobra"
)

var TreeFlags contract.TreeFlags

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tree",
	Short: "print directory structures in tree format",
	RunE: func(cmd *cobra.Command, args []string) error {
		var req contract.TreeRequest
		req.Flags = TreeFlags
		if len(args) > 0 {
			req.FolderName = args[0]
		}
		resp, err := handler.ProcessTreeRequest(req)
		if err != nil {
			handler.PrintStdOut(os.Stderr, err.Error())
			os.Exit(1)
		}

		handler.PrintResponse(os.Stdout, req, resp)

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
	TreeFlags = contract.TreeFlags{}
	rootCmd.Flags().BoolVarP(&TreeFlags.RelatviePath, "relativePath", "f", false, "print relative path")
}
