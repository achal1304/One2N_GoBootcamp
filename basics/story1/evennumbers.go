package story1

func EvenNumbers(num []int) []int {
	out := make([]int, 0, len(num))
	for _, v := range num {
		if v%2 == 0 {
			out = append(out, v)
		}
	}
	return out
}
