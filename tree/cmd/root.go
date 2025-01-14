/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
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
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		handleCombinedFlags(cmd)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		var req contract.TreeRequest
		req.Flags = TreeFlags
		if len(args) > 0 {
			req.FolderName = args[0]
		}

		if TreeFlags.Levels <= 0 {
			handler.PrintStdOut(os.Stderr, errors.New("tree: Invalid level, must be greater than 0.").Error())
			os.Exit(1)
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
	rootCmd.Flags().BoolVarP(&TreeFlags.RelativePath, "relativePath", "f", false, "print relative path")
	rootCmd.Flags().BoolVarP(&TreeFlags.DirectoryPrint, "printDirectories", "d", false, "print directories only")
	rootCmd.Flags().BoolVarP(&TreeFlags.Permission, "permission", "p", false, "print permissions")
	rootCmd.Flags().BoolVarP(&TreeFlags.RecentlyModified, "recentlyModified", "t", false, "print recently modified first")
	rootCmd.Flags().BoolVarP(&TreeFlags.XmlOutput, "xmlOutput", "X", false, "print xml output")
	rootCmd.Flags().BoolVarP(&TreeFlags.JsonOutput, "jsonOutput", "J", false, "print json output")
	rootCmd.Flags().BoolVarP(&TreeFlags.Graphics, "graphicsOption", "i", false, "print without indentation")
	rootCmd.Flags().IntVarP(&TreeFlags.Levels, "nestedLevels", "L", contract.MaxLevel, "print nested levels only")
}

func handleCombinedFlags(cmd *cobra.Command) {
	for _, arg := range os.Args {
		if len(arg) > 2 && arg[0] == '-' && arg[1] != '-' {
			// Loop through combined flags (e.g., -if)
			for _, flag := range arg[1:] {
				switch flag {
				case 'i':
					_ = cmd.Flags().Set("graphicsOption", "true")
				case 'f':
					_ = cmd.Flags().Set("relativePath", "true")
				}
			}
		}
	}
}
