package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/achal1304/One2N_GoBootcamp/basics/story1"
	"github.com/achal1304/One2N_GoBootcamp/basics/story2"
	"github.com/achal1304/One2N_GoBootcamp/basics/story3"
	"github.com/achal1304/One2N_GoBootcamp/basics/story4"
	"github.com/achal1304/One2N_GoBootcamp/basics/story5"
	"github.com/achal1304/One2N_GoBootcamp/basics/story6"
	"github.com/achal1304/One2N_GoBootcamp/basics/story7"
	"github.com/achal1304/One2N_GoBootcamp/basics/utils"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter a number:")
	scanner.Scan()
	input := scanner.Text()
	storyNumber, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return
	}
	switch storyNumber {
	case 1:
		fmt.Println(story1.EvenNumbers(takeInput()))
	case 2:
		fmt.Println(story2.OddNumbers(takeInput()))
	case 3:
		fmt.Println(story3.PrimeNumbers(takeInput()))
	case 4:
		fmt.Println(story4.OddPrimeNumbers(takeInput()))
	case 5:
		fmt.Println(story5.EvenAndDivisibleByFive(takeInput()))
	case 6:
		fmt.Println(story6.OddMultiplesOfThree(takeInput()))
	case 7:
		conditions := []func(int) bool{utils.IsOdd, utils.IsGreaterThanTen, utils.IsPrime}
		input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 16, 17, 18, 19}
		result := story7.MultiConditionsCheck[int](input, utils.Any[int], conditions)
		fmt.Println(result)
	}
}

func takeInput() []int {
	fmt.Print("Sample Input : ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Failed to read input: %v\n", err)
		return []int{}
	}
	values := strings.Split(input, ",")
	for i := 0; i < len(values); i++ {
		values[i] = strings.TrimSpace(values[i])
	}

	// Convert string values to integers and store them in a slice
	numbers := make([]int, len(values))
	for i, v := range values {
		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Printf("Failed to convert value '%s' to an integer: %v\n", v, err)
			return []int{}
		}
		numbers[i] = num
	}
	return numbers
}
