package day_4

import (
	"github.com/voziv/advent-of-code-2023/internal/util"
	"strconv"
	"strings"
)

func runPartOne(inputFileName string) string {
	lines := util.GetFileContents(inputFileName)

	sum := 0
	for _, line := range lines {
		_, results, _ := strings.Cut(line, ": ")
		winningNumberPart, cardNumberPart, _ := strings.Cut(results, "|")
		winningNumbers := parseIntoNumbers(winningNumberPart)
		cardNumbers := parseIntoNumbers(cardNumberPart)

		matches := 0
		for _, cardNumber := range cardNumbers {
			for _, winningNumber := range winningNumbers {
				if cardNumber == winningNumber {
					matches++
				}
			}
		}

		total := 0
		if matches > 0 {
			total = 1
			for i := 0; i < matches-1; i++ {
				total = total * 2
			}

		}

		sum += total

	}
	return strconv.Itoa(sum)
}

func parseIntoNumbers(input string) []int {
	var numbers []int

	numberStrings := strings.Split(input, " ")
	for _, item := range numberStrings {
		number, err := strconv.Atoi(item)
		if err != nil {
			continue
		}

		numbers = append(numbers, number)
	}

	return numbers
}
