package day_7

const (
	FiveOfAKind  = 7
	FourOfAKind  = 6
	FullHouse    = 5
	ThreeOfAKind = 4
	TwoPair      = 3
	OnePair      = 2
	HighCard     = 1
)

type Hand struct {
	cards []string
	bid   int
}

func (h *Hand) getType() int {
	matches := map[string]int{}
	for _, card := range h.cards {
		if _, ok := matches[card]; !ok {
			matches[card] = 1
		} else {
			matches[card]++
		}
	}

	pairs := 0
	trips := 0

	for _, count := range matches {
		if count == 5 {
			return FiveOfAKind
		} else if count == 4 {
			return FourOfAKind
		} else if count == 3 {
			trips++
		} else if count == 2 {
			pairs++
		}
	}

	if pairs == 2 {
		return TwoPair
	} else if pairs == 1 && trips == 1 {
		return FullHouse
	} else if trips == 1 {
		return ThreeOfAKind
	} else if pairs == 1 {
		return OnePair
	}

	return HighCard
}
