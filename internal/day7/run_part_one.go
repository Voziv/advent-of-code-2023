package day7

import (
	"sort"
	"strconv"
)

func runPartOne(inputFileName string) string {
	hands := parseFileIntoHands(inputFileName, false)

	sort.Slice(hands, func(i, j int) bool {
		return hands[i].compareForSorting(hands[j])
	})

	totalWinnings := 0
	for i, hand := range hands {
		winnings := hand.bid * (i + 1)
		totalWinnings += winnings
		//fmt.Printf("Hand: %+v\tBid: %+v \t$%d\n", hand.cards, hand.bid, winnings)
	}
	return strconv.Itoa(totalWinnings)
}
