package handler

import (
	"fmt"
	"io"

	"github.com/achal1304/One2N_GoBootcamp/tree/contract"
)

func PrintResponseStdOut(writer io.Writer, response contract.TreeResponse) {
	root := response.Root
	fmt.Fprintln(writer, root.Name)
	PrintTree(writer, root, 0, fmt.Sprint(""))
	finalCount := fmt.Sprintf("\n%d directories, %d files", response.DirectoryCount, response.FileCount)
	fmt.Fprintln(writer, finalCount)
}

func PrintTree(writer io.Writer, response *contract.TreeNode, iteration int, printer string) {
	if response == nil {
		return
	}
	var newPrinter string
	iteration++
	for i, node := range response.NextDir {
		if node.IsDir && i < len(response.NextDir)-1 {
			newPrinter = printer + "|-- " + node.Name
			fmt.Fprintln(writer, newPrinter)
			PrintTree(writer, node, iteration, printer+"|   ")
		} else if node.IsDir && i >= len(response.NextDir)-1 {
			newPrinter = printer + "|-- " + node.Name
			fmt.Fprintln(writer, newPrinter)
			PrintTree(writer, node, iteration, printer+"    ")
		}
		if !node.IsDir {
			newPrinter = printer + "|-- " + node.Name
			fmt.Fprintln(writer, newPrinter)
		}
	}
	return
}

func PrintStdOut(writer io.Writer, text string) {
	fmt.Fprintln(writer, text)
}
