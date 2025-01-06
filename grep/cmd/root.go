/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"io"
	"os"

	"github.com/achal1304/One2N_GoBootcamp/grep/contract"
	"github.com/achal1304/One2N_GoBootcamp/grep/handler"
	"github.com/achal1304/One2N_GoBootcamp/grep/helper"
	"github.com/spf13/cobra"
)

var GrepFlags contract.GrepFlags

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "grep",
	Short: "filter out by searching texts",
	RunE: func(cmd *cobra.Command, args []string) error {
		req := contract.GrepRequest{
			Flags: GrepFlags,
		}

		if len(args) >= 1 {
			req.SearchString = []byte(args[0])

			var reader io.Reader
			if len(args) == 3 && !GrepFlags.OutputFile {
				handler.PrintStdOut(os.Stderr, errors.New("grep: output file flag not specified but 3 arguments given").Error())
				os.Exit(1)
			}

			switch {
			case !GrepFlags.OutputFile && len(args) == 2:
				req.FileName = args[1]

			case GrepFlags.OutputFile:
				if len(args) == 3 {
					req.FileName = args[1]
					req.OutputFileName = args[2]
				} else {
					handler.PrintStdOut(os.Stderr, errors.New("grep: output file name not specified").Error())
					os.Exit(1)
				}

			default:
				reader = os.Stdin
			}

			response, err := handler.ProcessGrepRequest(req, reader)
			if err != nil {
				handler.PrintStdOut(os.Stderr, err.Error())
				os.Exit(1)
			}

			if !req.Flags.OutputFile {
				handler.PrintResponseStdOut(os.Stdout, response)
			} else {
				file, err := helper.GenerateFile(req.OutputFileName)
				if err != nil {
					handler.PrintStdOut(os.Stderr, err.Error())
					os.Exit(1)
				}
				defer file.Close()
				handler.PrintResponseStdOut(file, response)
			}
		} else {
			handler.PrintStdOut(os.Stderr, "error: not enough arguements")
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
	GrepFlags = contract.GrepFlags{}
	rootCmd.Flags().BoolVarP(&GrepFlags.OutputFile, "outputFile", "o", false, "output to a file")
	rootCmd.Flags().BoolVarP(&GrepFlags.CaseInsensitive, "insensitive", "i", false, "case insensitive search")
	rootCmd.Flags().BoolVarP(&GrepFlags.FolderCheck, "directorysearch", "r", false, "search in directories")
	rootCmd.Flags().IntVarP(&GrepFlags.AfterSearch, "aftersearch", "A", 0, "searched line and nlines after the result")
	rootCmd.Flags().IntVarP(&GrepFlags.BeforeSearch, "beforesearch", "B", 0, "searched line and nlines before the result.")
	rootCmd.Flags().IntVarP(&GrepFlags.BetweenSearch, "betweensearch", "C", 0, "searched line and nlines before, after the result.")
}
