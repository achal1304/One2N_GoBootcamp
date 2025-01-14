package handler

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"

	"github.com/achal1304/One2N_GoBootcamp/tree/contract"
)

func WriteXML(writer io.Writer, req contract.TreeRequest, response contract.TreeResponse) {
	encoder := xml.NewEncoder(writer)
	encoder.Indent("", "  ") // Pretty print

	fmt.Fprint(writer, xml.Header)

	// Start <tree> tag
	if err := startElement(encoder, "tree"); err != nil {
		return
	}

	// Print the directory and file structure
	if err := PrintXMLTree(req, response.Root, encoder); err != nil {
		logXMLError("Error writing XML tree", err)
		return
	}

	// Write <report> with directory and file counts
	if err := writeReport(encoder, response, req.Flags.DirectoryPrint); err != nil {
		logXMLError("Error writing report", err)
		return
	}

	// Close </tree> tag
	if err := endElement(encoder, "tree"); err != nil {
		return
	}

	// Flush encoder to ensure all data is written
	if err := encoder.Flush(); err != nil {
		logXMLError("Error flushing encoder", err)
	}
}

func PrintXMLTree(req contract.TreeRequest, node *contract.TreeNode, encoder *xml.Encoder) error {
	tagName := "file"
	if node.IsDir {
		tagName = "directory"
	}

	if err := startElementWithAttr(req, encoder, tagName, "name", node); err != nil {
		return err
	}

	for _, child := range node.NextDir {
		if err := PrintXMLTree(req, child, encoder); err != nil {
			return err
		}
	}

	return endElement(encoder, tagName)
}

func writeReport(encoder *xml.Encoder, response contract.TreeResponse, directoryOnly bool) error {
	if err := startElement(encoder, "report"); err != nil {
		return err
	}

	// Write <directories> count
	if err := writeSimpleElement(encoder, "directories", fmt.Sprintf("%d", response.DirectoryCount)); err != nil {
		return err
	}

	// Write <files> count if not in directory-only mode
	if !directoryOnly {
		if err := writeSimpleElement(encoder, "files", fmt.Sprintf("%d", response.FileCount)); err != nil {
			return err
		}
	}

	return endElement(encoder, "report")
}

func startElement(encoder *xml.Encoder, name string) error {
	return encoder.EncodeToken(xml.StartElement{Name: xml.Name{Local: name}})
}

func endElement(encoder *xml.Encoder, name string) error {
	return encoder.EncodeToken(xml.EndElement{Name: xml.Name{Local: name}})
}

func startElementWithAttr(req contract.TreeRequest, encoder *xml.Encoder, name, attrKey string, node *contract.TreeNode) error {
	nameValue := node.Name
	if req.Flags.RelativePath {
		nameValue = node.RelativePath
	}

	elem := xml.StartElement{
		Name: xml.Name{Local: name},
		Attr: []xml.Attr{
			{Name: xml.Name{Local: attrKey}, Value: nameValue},
		},
	}

	if req.Flags.Permission {
		elem.Attr = append(elem.Attr, []xml.Attr{
			{Name: xml.Name{Local: "mode"}, Value: node.PermissionOctal},
			{Name: xml.Name{Local: "prot"}, Value: node.Permission},
		}...)
	}

	return encoder.EncodeToken(elem)
}

func writeSimpleElement(encoder *xml.Encoder, tagName, value string) error {
	if err := startElement(encoder, tagName); err != nil {
		return err
	}
	if err := encoder.EncodeToken(xml.CharData(value)); err != nil {
		return err
	}
	return endElement(encoder, tagName)
}

func logXMLError(message string, err error) {
	PrintStdOut(os.Stderr, fmt.Sprintf("tree: %s: %v", message, err))
}
