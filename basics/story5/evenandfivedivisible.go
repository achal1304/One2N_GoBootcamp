package story5

import "github.com/achal1304/One2N_GoBootcamp/basics/utils"

func EvenAndDivisibleByFive(nums []int) []int {
	res := make([]int, 0, len(nums))
	for _, num := range nums {
		if utils.IsEven(num) && utils.IsDivisibleBy(num, 5) {
			res = append(res, num)
		}
	}
	return res
}
