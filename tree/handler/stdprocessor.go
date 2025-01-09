package handler

import (
	"fmt"
	"io"

	"github.com/achal1304/One2N_GoBootcamp/tree/contract"
)

func PrintResponse(writer io.Writer, req contract.TreeRequest, response contract.TreeResponse) {
	root := response.Root
	PrintStdOut(writer, root.Name)
	PrintTree(writer, req, root, 0, fmt.Sprint(""))
	finalCount := fmt.Sprintf("\n%d directories, %d files", response.DirectoryCount, response.FileCount)
	PrintStdOut(writer, finalCount)
}

func PrintTree(writer io.Writer, req contract.TreeRequest, response *contract.TreeNode, iteration int, printer string) {
	if response == nil {
		return
	}
	var newPrinter string
	iteration++
	for i, node := range response.NextDir {
		if !req.Flags.RelatviePath {
			newPrinter = printer + "|-- " + node.Name
		} else {
			newPrinter = printer + "|-- " + node.RelativePath
		}

		if node.IsDir && i < len(response.NextDir)-1 {
			PrintStdOut(writer, newPrinter)
			PrintTree(writer, req, node, iteration, printer+"|   ")
		} else if node.IsDir && i >= len(response.NextDir)-1 {
			PrintStdOut(writer, newPrinter)
			PrintTree(writer, req, node, iteration, printer+"    ")
		}
		if !node.IsDir {
			PrintStdOut(writer, newPrinter)
		}
	}
	return
}

func PrintStdOut(writer io.Writer, text string) {
	fmt.Fprintln(writer, text)
}
