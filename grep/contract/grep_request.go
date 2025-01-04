package contract

type GrepFlags struct {
	OutputFile      bool
	CaseInsensitive bool
	FolderCheck     bool
}

type GrepRequest struct {
	IsCaseSensitive bool
	SearchString    []byte
	FileName        string
	OutputFileName  string
	Flags           GrepFlags
}

type GrepResponse struct {
	SearchedText map[string][][]byte
	Flags        GrepFlags
	Err          error
}
