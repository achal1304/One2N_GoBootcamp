package handler

import (
	"fmt"
	"io"

	"github.com/achal1304/One2N_GoBootcamp/grep/contract"
)

func PrintResponseStdOut(writer io.Writer, response contract.GrepResponse) {
	for fileName, resp := range response.SearchedText {
		for i, text := range resp {
			if response.Flags.FolderCheck {
				fmt.Fprint(writer, fileName+fmt.Sprintf(":%s", string(text)))
			} else {
				fmt.Fprint(writer, string(text))
			}
			if (response.Flags.AfterSearch > 0 ||
				response.Flags.BeforeSearch > 0 ||
				response.Flags.BetweenSearch > 0) && !(i >= len(resp)-1) {
				fmt.Fprintln(writer, "--")
			}
		}
	}
}

func PrintStdOut(writer io.Writer, text string) {
	fmt.Fprintln(writer, text)
}
