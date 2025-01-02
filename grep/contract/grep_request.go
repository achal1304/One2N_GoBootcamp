package contract

type GrepFlags struct {
	OutputFile bool
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
	Err          error
}
