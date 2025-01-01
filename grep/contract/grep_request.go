package contract

type GrepRequest struct {
	IsCaseSensitive bool
	SearchString    []byte
	FileName        string
}

type GrepResponse struct {
	SearchedText map[string][][]byte
	Err          error
}
