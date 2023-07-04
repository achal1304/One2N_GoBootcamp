package story1

import "testing"

func TestEvenNumbers(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}
	expectedOutput := []int{2, 4, 6}

	actualOutput := EvenNumbers(input)

	for i := 0; i < len(actualOutput); i++ {
		if actualOutput[i] != expectedOutput[i] {
			t.Errorf("Expected %d obtained %d", expectedOutput[i], actualOutput[i])
		}
	}
	if len(actualOutput) != len(expectedOutput) {
		t.Errorf("Invalid length")
	}

}
