package story6

import "github.com/achal1304/One2N_GoBootcamp/basics/utils"

func OddMultiplesOfThree(nums []int) []int {
	res := make([]int, 0, len(nums))
	for _, num := range nums {
		if utils.IsOdd(num) && utils.IsDivisibleBy(num, 3) && utils.IsGreaterThanTen(num) {
			res = append(res, num)
		}
	}
	return res
}
