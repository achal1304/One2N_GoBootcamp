package contract

type WcFlags struct {
	LineCount      bool
	WordCount      bool
	CharacterCount bool
}

func NewFlags() WcFlags {
	return WcFlags{}
}
