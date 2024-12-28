package wchandler

import (
	"fmt"
	"io"

	"github.com/achal1304/One2N_GoBootcamp/wordcount/contract"
)

func GenerateOutput(wcValues contract.WcValues, wcFlags contract.WcFlags) string {
	output := ""
	if wcFlags.LineCount || wcFlags.WordCount || wcFlags.CharacterCount {
		if wcFlags.LineCount {
			output += fmt.Sprintf("%8d", wcValues.LineCount)
		}

		if wcFlags.WordCount {
			output += fmt.Sprintf("%8d", wcValues.WordCount)
		}

		if wcFlags.CharacterCount {
			output += fmt.Sprintf("%8d", wcValues.CharacterCount)
		}
		output += fmt.Sprintf(" %s", wcValues.FileName)
	} else {
		output += fmt.Sprintf("%8d", wcValues.LineCount)
		output += fmt.Sprintf("%8d", wcValues.WordCount)
		output += fmt.Sprintf("%8d", wcValues.CharacterCount)
		output += fmt.Sprintf(" %s", wcValues.FileName)
	}

	return output
}

func PrintStdOut(writer io.Writer, text string) {
	fmt.Fprint(writer, text)
}
