package story3

import "testing"

func TestIsPrime(t *testing.T) {
	t.Run("LessThanTwo_ReturnsFalse", func(t *testing.T) {
		input := 1
		expectedOutput := false

		isPrimeOutput := isPrime(input)

		if expectedOutput != isPrimeOutput {
			t.Errorf("Expected %v got %v", expectedOutput, isPrimeOutput)
		}
	})

	t.Run("IsNotPrime_ReturnsFalse", func(t *testing.T) {
		input := 8
		expectedOutput := false

		isPrimeOutput := isPrime(input)

		if expectedOutput != isPrimeOutput {
			t.Errorf("Expected %v got %v", expectedOutput, isPrimeOutput)
		}
	})

	t.Run("IsPrime_ReturnsTrue", func(t *testing.T) {
		input := 7
		expectedOutput := true

		isPrimeOutput := isPrime(input)

		if expectedOutput != isPrimeOutput {
			t.Errorf("Expected %v got %v", expectedOutput, isPrimeOutput)
		}
	})
}

func TestPrimeNumbers(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expectedOutput := []int{2, 3, 5, 7}

	actualOutput := PrimeNumbers(input)

	if len(actualOutput) != len(expectedOutput) {
		t.Errorf("Invalid length")
	}

	for i := 0; i < len(actualOutput); i++ {
		if actualOutput[i] != expectedOutput[i] {
			t.Errorf("Expected %d obtained %d", expectedOutput[i], actualOutput[i])
		}
	}

}
