package day_7

import (
	"fmt"
	"sort"
	"strconv"
)

func runPartTwo(inputFileName string) string {
	hands := parseFileIntoHands(inputFileName, true)

	sort.Slice(hands, func(i, j int) bool {
		return hands[i].compareForSorting(hands[j])
	})

	totalWinnings := 0
	for i, hand := range hands {
		winnings := hand.bid * (i + 1)
		totalWinnings += winnings
		fmt.Printf("Hand: %+v\t%-20s\t$%d\n", hand.cards, hand.Type(), winnings)
	}
	return strconv.Itoa(totalWinnings)
}
