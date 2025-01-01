package contract

import "github.com/achal1304/One2N_GoBootcamp/wordcount/wcerrors"

type WcValues struct {
	FileName       string
	LineCount      int
	WordCount      int
	CharacterCount int
	Err            *wcerrors.WcError
}
