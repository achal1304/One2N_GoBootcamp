package story8

import (
	"testing"

	"github.com/achal1304/One2N_GoBootcamp/basics/story7"
	"github.com/achal1304/One2N_GoBootcamp/basics/utils"
)

func TestMultiConditionsChekc(t *testing.T) {
	conditions := []func(int) bool{utils.IsOdd, utils.IsGreaterThanTen, utils.IsPrime}
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 16, 17, 18, 19}
	expectedOutput := []int{1, 2, 3, 5, 7, 9, 11, 12, 13, 14, 16, 17, 18, 19}

	actualOutput := story7.MultiConditionsCheck[int](input, utils.Any[int], conditions)

	if len(actualOutput) != len(expectedOutput) {
		t.Errorf("Invalid length")
	}

	for i := 0; i < len(actualOutput); i++ {
		if actualOutput[i] != expectedOutput[i] {
			t.Errorf("Expected %d obtained %d", expectedOutput[i], actualOutput[i])
		}
	}
}
