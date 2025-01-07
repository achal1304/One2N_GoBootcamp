package handler

import (
	"fmt"
	"io"

	"github.com/achal1304/One2N_GoBootcamp/tree/contract"
)

func PrintResponseStdOut(writer io.Writer, response contract.TreeResponse) {
	root := response.Root
	fmt.Println(root.Name, root.RelativePath)
	PrintTree(writer, root, 0)
	finalCount := fmt.Sprintf("%d directories, %d files", response.DirectoryCount, response.FileCount)
	fmt.Fprintln(writer, finalCount)
}

func PrintTree(writer io.Writer, response *contract.TreeNode, iteration int) {

	if response == nil {
		return
	}
	iteration++
	for _, node := range response.NextDir {
		fmt.Fprintln(writer, node.RelativePath)
		PrintTree(writer, node, iteration)
	}
	return
}

func PrintStdOut(writer io.Writer, text string) {
	fmt.Fprintln(writer, text)
}
