package day4

import (
	"github.com/voziv/advent-of-code-2023/internal/util"
	"strconv"
)

func runPartTwo(inputFileName string) string {
	lines := util.GetFileContents(inputFileName)

	cards := map[int]*Card{}
	var cardStack []*Card

	for _, line := range lines {
		card := parseLineIntoCard(line)
		cards[card.Id()] = card
		cardStack = append(cardStack, card)
	}

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

	return strconv.Itoa(len(cardStack))
}
