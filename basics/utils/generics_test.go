package utils

import "testing"

func TestAll(t *testing.T) {
	conditions := []func(int) bool{IsEven, IsGreaterThanTen}
	t.Run("True", func(t *testing.T) {
		input := 20
		expectedOutput := true

		actualOutput := All(conditions, input)

		if actualOutput != expectedOutput {
			t.Errorf("Expected %v got %v", expectedOutput, actualOutput)
		}
	})

	t.Run("False", func(t *testing.T) {
		input := 25
		expectedOutput := false

		actualOutput := All(conditions, input)

		if actualOutput != expectedOutput {
			t.Errorf("Expected %v got %v", expectedOutput, actualOutput)
		}
	})
}

func TestAny(t *testing.T) {
	conditions := []func(int) bool{IsEven, IsGreaterThanTen}
	t.Run("True", func(t *testing.T) {
		input := 2
		expectedOutput := true

		actualOutput := Any(conditions, input)

		if actualOutput != expectedOutput {
			t.Errorf("Expected %v got %v", expectedOutput, actualOutput)
		}
	})

	t.Run("False", func(t *testing.T) {
		input := 9
		expectedOutput := false

		actualOutput := Any(conditions, input)

		if actualOutput != expectedOutput {
			t.Errorf("Expected %v got %v", expectedOutput, actualOutput)
		}
	})
}
