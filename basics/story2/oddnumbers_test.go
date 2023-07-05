package story2

import "testing"

func TestOddNumbers(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}
	expectedOutput := []int{1, 3, 5}

	actualOutput := OddNumbers(input)

	for i := 0; i < len(actualOutput); i++ {
		if actualOutput[i] != expectedOutput[i] {
			t.Errorf("Expected %d obtained %d", expectedOutput[i], actualOutput[i])
		}
	}
	if len(actualOutput) != len(expectedOutput) {
		t.Errorf("Invalid length")
	}

}
