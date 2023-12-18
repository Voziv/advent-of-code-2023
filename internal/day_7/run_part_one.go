package day_7

import (
	"fmt"
	"github.com/voziv/advent-of-code-2023/internal/util"
	"sort"
	"strconv"
	"strings"
)

var cardValues = map[string]int{
	"2": 1,
	"3": 2,
	"4": 3,
	"5": 4,
	"6": 5,
	"7": 6,
	"8": 7,
	"9": 8,
	"T": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
	"A": 14,
}

func runPartOne(inputFileName string) string {
	lines := util.GetFileContents(inputFileName)
	var hands []*Hand
	for _, line := range lines {
		tokens := strings.Split(line, " ")
		bid, err := strconv.Atoi(tokens[1])
		if err != nil {
			panic(err)
		}

		hands = append(hands, &Hand{strings.Split(tokens[0], ""), bid})
	}

	sort.Slice(hands, func(i, j int) bool {
		return compareHands(hands[i], hands[j])
	})

	totalWinnings := 0
	for i, hand := range hands {
		winnings := hand.bid * (i + 1)
		totalWinnings += winnings
		fmt.Printf("%+v\t$%d\n", hand, winnings)
	}
	return strconv.Itoa(totalWinnings)
}

func compareHands(i *Hand, j *Hand) bool {
	handI := i.getType()
	handJ := j.getType()

	if handI != handJ {
		return handI < handJ
	}

	for c := 0; c < 5; c++ {
		if i.cards[c] == j.cards[c] {
			continue
		}

		return cardValues[i.cards[c]] < cardValues[j.cards[c]]
	}

	panic("Hands are identical :/")
}
