package handler

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/achal1304/One2N_GoBootcamp/tree/contract"
)

type JSONNode struct {
	Type     string     `json:"type"`
	Name     string     `json:"name"`
	Mode     string     `json:"mode,omitempty"`
	Prot     string     `json:"prot,omitempty"`
	Contents []JSONNode `json:"contents,omitempty"`
}

type Report struct {
	Type        string `json:"type"`
	Directories int    `json:"directories"`
	Files       *int   `json:"files,omitempty"`
}

// WriteJSON generates the JSON output for the tree and writes it to the provided writer
func WriteJSON(writer io.Writer, req contract.TreeRequest, response contract.TreeResponse) error {
	treeJSON := buildJSONTree(req, response.Root)
	reportJSON := Report{
		Type:        "report",
		Directories: response.DirectoryCount,
	}

	if !req.Flags.DirectoryPrint {
		reportJSON.Files = &response.FileCount
	}
	// Combine tree and report
	output := []interface{}{treeJSON, reportJSON}

	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(output); err != nil {
		return fmt.Errorf("error writing JSON: %v", err)
	}

	return nil
}

func buildJSONTree(req contract.TreeRequest, node *contract.TreeNode) JSONNode {
	jsonNode := JSONNode{
		Type: "file",
		Name: node.Name,
	}

	if node.IsDir {
		jsonNode.Type = "directory"
		for _, child := range node.NextDir {
			jsonNode.Contents = append(jsonNode.Contents, buildJSONTree(req, child))
		}
	}

	if req.Flags.Permission {
		jsonNode.Mode = node.PermissionOctal
		jsonNode.Prot = node.Permission
	}

	if req.Flags.RelativePath {
		jsonNode.Name = node.RelativePath
	}

	return jsonNode
}
