package contract

type TreeFlags struct {
}

type TreeRequest struct {
	FolderName string
	Flags      TreeFlags
}

type TreeResponse struct {
	DirectoryCount int
	FileCount      int
	Root           *TreeNode
}

type TreeNode struct {
	Name         string
	Path         string
	IsDir        bool
	RelativePath string
	NextDir      []*TreeNode
}
