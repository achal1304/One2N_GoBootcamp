package story2

import "github.com/achal1304/One2N_GoBootcamp/basics/utils"

func OddNumbers(num []int) []int {
	out := make([]int, 0, len(num))
	for _, v := range num {
		if utils.IsOdd(v) {
			out = append(out, v)
		}
	}
	return out
}
