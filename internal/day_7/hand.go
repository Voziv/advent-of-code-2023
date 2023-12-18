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

const (
	Joker = "*"
	Two   = "2"
	Three = "3"
	Four  = "4"
	Five  = "5"
	Six   = "6"
	Seven = "7"
	Eight = "8"
	Nine  = "9"
	Ten   = "T"
	Jack  = "J"
	Queen = "Q"
	King  = "K"
	Ace   = "A"
)

var cardValues = map[string]int{
	Joker: 0, // Represent Joker for part 2
	Two:   1,
	Three: 2,
	Four:  3,
	Five:  4,
	Six:   5,
	Seven: 6,
	Eight: 7,
	Nine:  8,
	Ten:   10,
	Jack:  11,
	Queen: 12,
	King:  13,
	Ace:   14,
}

type Hand struct {
	cards []string
	bid   int
}

func (h *Hand) compareForSorting(h2 *Hand) bool {
	handI := h.TypeScore()
	handJ := h2.TypeScore()

	if handI != handJ {
		return handI < handJ
	}

	for c := 0; c < 5; c++ {
		if h.cards[c] == h2.cards[c] {
			continue
		}

		return cardValues[h.cards[c]] < cardValues[h2.cards[c]]
	}

	panic("Hands are identical :/")
}

func (h *Hand) TypeScore() int {
	matches := map[string]int{}
	for _, card := range h.cards {
		if _, ok := matches[card]; !ok {
			matches[card] = 1
		} else {
			matches[card]++
		}
	}

	jokers := 0
	pairs := 0
	trips := 0
	quads := 0

	for card, count := range matches {
		if card == Joker {
			jokers = count
			continue
		}

		if count == 5 {
			return FiveOfAKind
		} else if count == 4 {
			quads++
		} else if count == 3 {
			trips++
		} else if count == 2 {
			pairs++
		}
	}

	if isFiveOfAKind(jokers, pairs, trips, quads) {
		return FiveOfAKind
	}

	if isFourOfAKind(jokers, pairs, trips, quads) {
		return FourOfAKind
	}

	if isFullHouse(jokers, pairs, trips, quads) {
		return FullHouse
	}

	if isThreeOfAKind(jokers, pairs, trips, quads) {
		return ThreeOfAKind
	}

	if isTwoPairs(jokers, pairs, trips, quads) {
		return TwoPair
	}

	if isOnePair(jokers, pairs, trips, quads) {
		return OnePair
	}

	return HighCard
}

func (h *Hand) Type() string {
	switch h.TypeScore() {
	case FiveOfAKind:
		return "#7 - 5 of a kind"
	case FourOfAKind:
		return "#6 - 4 of a kind"
	case FullHouse:
		return "#5 - Full Haus"
	case ThreeOfAKind:
		return "#4 - 3 of a kind"
	case TwoPair:
		return "#3 - 2 pairs"
	case OnePair:
		return "#2 - 1 pair"
	case HighCard:
		return "#1 - High card"
	default:
		return "Error - Unknown Hand"
	}
}

func isFiveOfAKind(jokers int, pairs int, trips int, quads int) bool {
	// *****
	// 1****
	// 11***
	// 111**
	// 1111*
	// 11111 Is handled by the card counting loop
	return (jokers == 5) || (jokers == 4) || (jokers == 3 && pairs == 1) || (jokers == 2 && trips == 1) || (jokers == 1 && quads == 1)
}

func isFourOfAKind(jokers int, pairs int, trips int, quads int) bool {
	// 11112
	// 111*2
	// 11**2
	// 1***2
	// ****2 doesn't count because we'd score it as a five of a kind instead
	return (quads == 1 && jokers == 0) || (trips == 1 && jokers == 1) || (pairs == 1 && jokers == 2) || (pairs == 0 && jokers == 3)
}
func isFullHouse(jokers int, pairs int, trips int, quads int) bool {
	// 11222
	// 1122*
	// 112** doesn't count because we'd score it as a four of a kind instead
	return trips == 1 && pairs == 1 || pairs == 2 && jokers == 1
}
func isThreeOfAKind(jokers int, pairs int, trips int, quads int) bool {
	// 11123
	// 11*23
	// 123**
	return (trips == 1 && pairs == 0 && jokers == 0) || (pairs == 1 && jokers == 1) || (trips == 0 && pairs == 0 && jokers == 2)
}

func isTwoPairs(jokers int, pairs int, trips int, quads int) bool {
	// 11223
	// 112*3 doesn't count because we'd score it as a three of a kind
	return pairs == 2 && jokers == 0
}

func isOnePair(jokers int, pairs int, trips int, quads int) bool {
	// 11234
	// 1234*
	return trips == 0 && quads == 0 && (pairs == 1 && jokers == 0 || pairs == 0 && jokers == 1)
}
