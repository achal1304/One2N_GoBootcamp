package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	// "github.com/achal1304/One2N_GoBootcamp/basics/story1"
	// "github.com/achal1304/One2N_GoBootcamp/basics/story2"
	"github.com/achal1304/One2N_GoBootcamp/basics/story3"
)

func main() {

	// // Story1 : Even Numbers
	// fmt.Println(story1.EvenNumbers(takeInput()))
	// // Story2 : Odd Numbers
	// fmt.Println(story2.OddNumbers(takeInput()))
	// Story3 : Prime Numbers
	fmt.Println(story3.PrimeNumbers(takeInput()))
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
