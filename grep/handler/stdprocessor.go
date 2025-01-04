package handler

import (
	"fmt"
	"io"

	"github.com/achal1304/One2N_GoBootcamp/grep/contract"
)

func PrintResponseStdOut(writer io.Writer, response contract.GrepResponse) {
	for fileName, resp := range response.SearchedText {
		for _, text := range resp {
			if response.Flags.FolderCheck {
				fmt.Fprintln(writer, fileName+fmt.Sprintf(":%s", string(text)))
			} else {
				fmt.Fprintln(writer, string(text))
			}
		}
	}
}

func PrintStdOut(writer io.Writer, text string) {
	fmt.Fprintln(writer, text)
}
