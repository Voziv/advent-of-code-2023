package day7

import (
	"github.com/voziv/advent-of-code-2023/internal/util"
	"sort"
	"strconv"
	"strings"
)

type result struct {
	partOne int
	partTwo int
}

func Run() {
	util.AssertResult("example.txt", run("./internal/day7/example.txt"), result{
		partOne: 6440,
		partTwo: 5905,
	})

	util.AssertResult("input.txt", run("./internal/day7/input.txt"), result{
		partOne: 253910319,
		partTwo: 254083736,
	})
}

func run(inputFileName string) result {
	result := result{
		partOne: 0,
		partTwo: 0,
	}

	hands := parseFileIntoHands(inputFileName)
	result.partOne = calculateWinnings(hands)

	//
	// Part Two - Convert jacks to jokers
	//
	convertJacksToJokers(hands)
	result.partTwo = calculateWinnings(hands)

	return result
}

func calculateWinnings(hands []*Hand) int {
	sort.Slice(hands, func(i, j int) bool {
		return hands[i].compareForSorting(hands[j])
	})

	sum := 0

	for i, hand := range hands {
		sum += hand.bid * (i + 1)
		//fmt.Printf("Hand: %+v\t%-20s\n", hand.cards, hand.Type())
	}

	return sum
}

func convertJacksToJokers(hands []*Hand) {
	for _, hand := range hands {
		for i := 0; i < len(hand.cards); i++ {
			if hand.cards[i] == Jack {
				hand.cards[i] = Joker
			}
		}
	}
}

func parseFileIntoHands(inputFileName string) []*Hand {
	var hands []*Hand

	for _, line := range util.GetFileContents(inputFileName) {
		tokens := strings.Split(line, " ")
		bid, err := strconv.Atoi(tokens[1])
		if err != nil {
			panic(err)
		}

		cards := strings.Split(tokens[0], "")

		hands = append(hands, &Hand{cards, bid})
	}

	return hands
}
