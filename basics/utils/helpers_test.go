package utils

import (
	"testing"
)

func TestIsPrime(t *testing.T) {
	t.Run("LessThanTwo_ReturnsFalse", func(t *testing.T) {
		input := 1
		expectedOutput := false

		isPrimeOutput := IsPrime(input)

		if expectedOutput != isPrimeOutput {
			t.Errorf("Expected %v got %v", expectedOutput, isPrimeOutput)
		}
	})

	t.Run("IsNotPrime_ReturnsFalse", func(t *testing.T) {
		input := 8
		expectedOutput := false

		isPrimeOutput := IsPrime(input)

		if expectedOutput != isPrimeOutput {
			t.Errorf("Expected %v got %v", expectedOutput, isPrimeOutput)
		}
	})

	t.Run("IsPrime_ReturnsTrue", func(t *testing.T) {
		input := 7
		expectedOutput := true

		isPrimeOutput := IsPrime(input)

		if expectedOutput != isPrimeOutput {
			t.Errorf("Expected %v got %v", expectedOutput, isPrimeOutput)
		}
	})
}

func TestIsEven(t *testing.T) {
	t.Run("True", func(t *testing.T) {
		input := 4
		actualOutput := IsEven(input)
		if actualOutput != true {
			t.Errorf("Expected %v got %v", true, actualOutput)
		}
	})

	t.Run("False", func(t *testing.T) {
		input := 3
		actualOutput := IsEven(input)
		if actualOutput != false {
			t.Errorf("Expected %v got %v", false, actualOutput)
		}
	})
}

func TestIsOdd(t *testing.T) {
	t.Run("False", func(t *testing.T) {
		input := 4
		actualOutput := IsOdd(input)
		if actualOutput != false {
			t.Errorf("Expected %v got %v", false, actualOutput)
		}
	})

	t.Run("True", func(t *testing.T) {
		input := 3
		actualOutput := IsOdd(input)
		if actualOutput != true {
			t.Errorf("Expected %v got %v", true, actualOutput)
		}
	})
}

func TestIsDivisibleBy(t *testing.T) {
	t.Run("True", func(t *testing.T) {
		input := 4
		divisor := 2
		actualOutput := IsDivisibleBy(input, divisor)
		if actualOutput != true {
			t.Errorf("Expected %v got %v", true, actualOutput)
		}
	})

	t.Run("False", func(t *testing.T) {
		input := 3
		divisor := 2
		actualOutput := IsDivisibleBy(input, divisor)
		if actualOutput != false {
			t.Errorf("Expected %v got %v", false, actualOutput)
		}
	})

	t.Run("InvalidAndFalse", func(t *testing.T) {
		input := 3
		divisor := 0
		actualOutput := IsDivisibleBy(input, divisor)
		if actualOutput != false {
			t.Errorf("Expected %v got %v", false, actualOutput)
		}
	})
}

func TestIsGreaterThanTen(t *testing.T) {
	t.Run("True", func(t *testing.T) {
		input := 11
		actualOutput := IsGreaterThanTen(input)
		if actualOutput != true {
			t.Errorf("Expected %v got %v", true, actualOutput)
		}
	})

	t.Run("False", func(t *testing.T) {
		input := 3
		actualOutput := IsGreaterThanTen(input)
		if actualOutput != false {
			t.Errorf("Expected %v got %v", false, actualOutput)
		}
	})
}
