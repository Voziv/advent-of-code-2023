package day_4

import (
	"github.com/voziv/advent-of-code-2023/internal/util"
	"strconv"
	"strings"
)

func runPartOne(inputFileName string) string {
	lines := util.GetFileContents(inputFileName)

	sum := 0
	cards := map[int]*Card{}
	for _, line := range lines {
		card := parseLineIntoCard(line)
		cards[card.Id()] = card

		matches := card.Matches()

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
