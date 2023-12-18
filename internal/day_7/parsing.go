package day_7

import (
	"github.com/voziv/advent-of-code-2023/internal/util"
	"strconv"
	"strings"
)

func parseFileIntoHands(inputFileName string, partTwo bool) []*Hand {
	var hands []*Hand

	for _, line := range util.GetFileContents(inputFileName) {
		tokens := strings.Split(line, " ")
		bid, err := strconv.Atoi(tokens[1])
		if err != nil {
			panic(err)
		}

		cards := strings.Split(tokens[0], "")

		// Not a fan of bool switches, may refactor someday.
		if partTwo {
			cards = convertJacksToJokers(cards)
		}

		hands = append(hands, &Hand{cards, bid})
	}

	return hands
}

func convertJacksToJokers(tokens []string) []string {
	for i := 0; i < len(tokens); i++ {
		if tokens[i] == Jack {
			tokens[i] = Joker
		}
	}

	return tokens
}
