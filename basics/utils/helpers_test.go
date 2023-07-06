package utils

import "testing"

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
