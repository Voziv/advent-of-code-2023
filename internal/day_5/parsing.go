package day_5

import (
	"strconv"
	"strings"
)

func parseNumbersFromInput(input string) []int {
	var numbers []int
	tokens := strings.Split(input, " ")
	for _, token := range tokens {
		number, err := strconv.Atoi(token)
		if err != nil {
			continue
		}

		numbers = append(numbers, number)
	}

	return numbers
}
