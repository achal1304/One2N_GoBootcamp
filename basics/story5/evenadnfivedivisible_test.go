package story5

import "testing"

func TestEvenAndDivisibleByFive(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 15, 20}
	expectedOutput := []int{10, 20}

	actualOutput := EvenAndDivisibleByFive(input)

	if len(actualOutput) != len(expectedOutput) {
		t.Errorf("Invalid length")
	}

	for i := 0; i < len(actualOutput); i++ {
		if actualOutput[i] != expectedOutput[i] {
			t.Errorf("Expected %d obtained %d", expectedOutput[i], actualOutput[i])
		}
	}

}
