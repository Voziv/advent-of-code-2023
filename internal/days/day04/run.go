package day04

import (
	"github.com/voziv/advent-of-code-2023/internal/util"
	"strconv"
	"strings"
)

type result struct {
	partOne int
	partTwo int
}

func Run() {
	util.AssertResult("example.txt", run("./internal/days/day04/example.txt"), result{
		partOne: 13,
		partTwo: 30,
	})

	util.AssertResult("input.txt", run("./internal/days/day04/input.txt"), result{
		partOne: 18653,
		partTwo: 5921508,
	})
}

func run(inputFileName string) result {
	lines := util.GetFileContents(inputFileName)

	result := result{
		partOne: 0,
		partTwo: 0,
	}
	cards := map[int]*Card{}
	var cardStack []*Card
	for _, line := range lines {
		card := parseLineIntoCard(line)
		cards[card.Id()] = card
		cardStack = append(cardStack, card)

		matches := card.Matches()

		total := 0
		if matches > 0 {
			total = 1
			for i := 0; i < matches-1; i++ {
				total = total * 2
			}

		}

		result.partOne += total
	}

	// Part Two
	for i := 0; i < len(cardStack); i++ {
		matches := cardStack[i].Matches()
		for j := 1; j <= matches; j++ {
			card, ok := cards[cardStack[i].Id()+j]
			if !ok {
				break
			}
			cardStack = append(cardStack, card)
		}
	}

	result.partTwo = len(cardStack)

	return result
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
