package story3

import "testing"

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
