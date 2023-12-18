package day4

type Card struct {
	id             int
	winningNumbers []int
	cardNumbers    []int
}

func NewCard(id int, winningNumbers []int, cardNumbers []int) *Card {
	return &Card{id: id, winningNumbers: winningNumbers, cardNumbers: cardNumbers}
}

func (c Card) Id() int {
	return c.id
}

func (c Card) WinningNumbers() []int {
	return c.winningNumbers
}

func (c Card) CardNumbers() []int {
	return c.cardNumbers
}

func (c Card) Matches() int {
	matches := 0
	for _, cardNumber := range c.cardNumbers {
		for _, winningNumber := range c.winningNumbers {
			if cardNumber == winningNumber {
				matches++
				break
			}
		}
	}
	return matches
}
