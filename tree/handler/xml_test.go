package handler

import (
	"bytes"
	"testing"

	"github.com/achal1304/One2N_GoBootcamp/tree/contract"
)

func TestWriteXML(t *testing.T) {
	tree := &contract.TreeNode{
		Name:  "root",
		IsDir: true,
		NextDir: []*contract.TreeNode{
			{
				Name:  "dir1",
				IsDir: true,
				NextDir: []*contract.TreeNode{
					{Name: "file1", IsDir: false},
					{Name: "file2", IsDir: false},
				},
			},
			{
				Name:  "dir2",
				IsDir: true,
				NextDir: []*contract.TreeNode{
					{Name: "file3", IsDir: false},
				},
			},
			{
				Name:  "file4",
				IsDir: false,
			},
		},
	}

	req := contract.TreeRequest{
		Flags: contract.TreeFlags{
			DirectoryPrint: false,
		},
	}

	resp := contract.TreeResponse{
		Root:           tree,
		DirectoryCount: 2,
		FileCount:      4,
	}

	var buf bytes.Buffer
	WriteXML(&buf, req, resp)

	expected := `<?xml version="1.0" encoding="UTF-8"?>
<tree>
  <directory name="root">
    <directory name="dir1">
      <file name="file1"></file>
      <file name="file2"></file>
    </directory>
    <directory name="dir2">
      <file name="file3"></file>
    </directory>
    <file name="file4"></file>
  </directory>
  <report>
    <directories>2</directories>
    <files>4</files>
  </report>
</tree>`

	output := buf.String()
	if output != expected {
		t.Errorf("Expected:\n%s\nBut got:\n%s", expected, output)
	}
}

func TestWriteXMLDirectoryOnly(t *testing.T) {
	tree := &contract.TreeNode{
		Name:  "root",
		IsDir: true,
		NextDir: []*contract.TreeNode{
			{
				Name:    "dir1",
				IsDir:   true,
				NextDir: []*contract.TreeNode{},
			},
			{
				Name:    "dir2",
				IsDir:   true,
				NextDir: []*contract.TreeNode{},
			},
		},
	}

	req := contract.TreeRequest{
		Flags: contract.TreeFlags{
			DirectoryPrint: true,
		},
	}

	resp := contract.TreeResponse{
		Root:           tree,
		DirectoryCount: 2,
		FileCount:      0,
	}

	var buf bytes.Buffer
	WriteXML(&buf, req, resp)

	expected := `<?xml version="1.0" encoding="UTF-8"?>
<tree>
  <directory name="root">
    <directory name="dir1"></directory>
    <directory name="dir2"></directory>
  </directory>
  <report>
    <directories>2</directories>
  </report>
</tree>`

	output := buf.String()
	if output != expected {
		t.Errorf("Expected:\n%s\nBut got:\n%s", expected, output)
	}
}

func TestWriteXMLPermissionAndRelativePath(t *testing.T) {
	tree := &contract.TreeNode{
		Name:            "root",
		IsDir:           true,
		Permission:      "drwxrwxrwx",
		PermissionOctal: "0777",
		RelativePath:    "root",
		NextDir: []*contract.TreeNode{
			{Name: "file1.txt", IsDir: false, Permission: "-rw-rw-rw", PermissionOctal: "0666",
				RelativePath: "root/file1.txt",
			},
		},
	}

	req := contract.TreeRequest{
		Flags: contract.TreeFlags{
			DirectoryPrint: false,
			Permission:     true,
			RelativePath:   true,
		},
	}

	resp := contract.TreeResponse{
		Root:           tree,
		DirectoryCount: 1,
		FileCount:      1,
	}

	var buf bytes.Buffer
	WriteXML(&buf, req, resp)

	expected := `<?xml version="1.0" encoding="UTF-8"?>
<tree>
  <directory name="root" mode="0777" prot="drwxrwxrwx">
    <file name="root/file1.txt" mode="0666" prot="-rw-rw-rw"></file>
  </directory>
  <report>
    <directories>1</directories>
    <files>1</files>
  </report>
</tree>`

	output := buf.String()
	if output != expected {
		t.Errorf("Expected:\n%s\nBut got:\n%s", expected, output)
	}
}
