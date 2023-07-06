package utils

func IsPrime(num int) bool {
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

func IsEven(num int) bool {
	return num%2 == 0
}

func IsOdd(num int) bool {
	return num%2 != 0
}

func IsDivisibleBy(num int, divisor int) bool {
	if divisor == 0 {
		return false
	}

	return num%divisor == 0
}

func IsGreaterThanTen(num int) bool {
	return num > 10
}
