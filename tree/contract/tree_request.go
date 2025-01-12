package contract

type TreeFlags struct {
	RelativePath     bool
	DirectoryPrint   bool
	Permission       bool
	Levels           int
	RecentlyModified bool
	XmlOutput        bool
	JsonOutput       bool
}

const MaxLevel = 999999999

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
	Permission   string
	NextDir      []*TreeNode
}
