package handler

import (
	"fmt"
	"io"
	"os"

	"github.com/achal1304/One2N_GoBootcamp/tree/contract"
)

func PrintResponse(writer io.Writer, req contract.TreeRequest, response contract.TreeResponse) {
	if req.Flags.XmlOutput {
		WriteXML(writer, req, response)
	} else if req.Flags.JsonOutput {
		err := WriteJSON(writer, req, response)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing JSON: %v\n", err)
			os.Exit(1)
		}
	} else {
		WritePlainText(writer, req, response)
	}
}

func WritePlainText(writer io.Writer, req contract.TreeRequest, response contract.TreeResponse) {
	root := response.Root
	var finalCount string
	PrintStdOut(writer, getPrinter(req, root, "", ""))
	if !req.Flags.Graphics {
		PrintTree(writer, req, root, 0, "")
	} else {
		PrintTreeWithoutGraphics(writer, req, root, "")
	}

	if req.Flags.DirectoryPrint {
		finalCount = fmt.Sprintf("\n%d directories", response.DirectoryCount)
	} else {
		finalCount = fmt.Sprintf("\n%d directories, %d files", response.DirectoryCount, response.FileCount)
	}
	PrintStdOut(writer, finalCount)
}

// func PrintTree(writer io.Writer, req contract.TreeRequest, response *contract.TreeNode, iteration int, printer string) {
// 	if response == nil {
// 		return
// 	}
// 	iteration++
// 	for i, node := range response.NextDir {
// 		if node.IsDir && i < len(response.NextDir)-1 {
// 			PrintStdOut(writer, getPrinter(req, node, printer, "|-- "))
// 			PrintTree(writer, req, node, iteration, printer+"|   ")
// 		} else if node.IsDir && i >= len(response.NextDir)-1 {
// 			PrintStdOut(writer, getPrinter(req, node, printer, "|-- "))
// 			PrintTree(writer, req, node, iteration, printer+"    ")
// 		}
// 		if !node.IsDir {
// 			PrintStdOut(writer, getPrinter(req, node, printer, "|-- "))
// 		}
// 	}
// 	return
// }

func PrintTree(writer io.Writer, req contract.TreeRequest, response *contract.TreeNode, iteration int, printer string) {
	if response == nil {
		return
	}
	iteration++

	for i, node := range response.NextDir {
		graphicPrefix := "|-- "
		nextGraphicPrefix := "|   "
		if i >= len(response.NextDir)-1 {
			nextGraphicPrefix = "    "
		}
		PrintStdOut(writer, getPrinter(req, node, printer, graphicPrefix))
		if node.IsDir {
			PrintTree(writer, req, node, iteration, printer+nextGraphicPrefix)
		}
	}
}

func PrintTreeWithoutGraphics(writer io.Writer, req contract.TreeRequest, response *contract.TreeNode, printer string) {
	if response == nil {
		return
	}

	for _, node := range response.NextDir {
		PrintStdOut(writer, getPrinter(req, node, printer, ""))
		if node.IsDir {
			PrintTreeWithoutGraphics(writer, req, node, printer)
		}
	}
}

func PrintStdOut(writer io.Writer, text string) {
	fmt.Fprintln(writer, text)
}

func getPrinter(req contract.TreeRequest, node *contract.TreeNode, printer string, seperator string) string {
	newPrinter := ""
	if !req.Flags.RelativePath && !req.Flags.Permission {
		newPrinter = printer + seperator + node.Name
	} else if req.Flags.RelativePath && !req.Flags.Permission {
		newPrinter = printer + seperator + node.RelativePath
	} else if !req.Flags.RelativePath && req.Flags.Permission {
		newPrinter = printer + seperator + "[" + node.Permission + "]" + " " + node.Name
	} else {
		newPrinter = printer + seperator + "[" + node.Permission + "]" + " " + node.RelativePath
	}
	return newPrinter
}
