package story3

import "github.com/achal1304/One2N_GoBootcamp/basics/utils"

func PrimeNumbers(nums []int) []int {
	res := make([]int, 0, len(nums))
	for _, n := range nums {
		if utils.IsPrime(n) {
			res = append(res, n)
		}
	}
	return res
}
