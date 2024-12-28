package contract

type WcFlags struct {
	LineCount bool
	WordCount bool
}

func NewFlags() WcFlags {
	return WcFlags{}
}
