package story3

func isPrime(num int) bool {
	if num < 2 {
		return false
	}

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func PrimeNumbers(nums []int) []int {
	res := make([]int, 0, len(nums))
	for _, n := range nums {
		if isPrime(n) {
			res = append(res, n)
		}
	}
	return res
}
