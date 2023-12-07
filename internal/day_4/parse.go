package day_4

import (
	"strconv"
	"strings"
)

func parseLineIntoCard(input string) *Card {
	cardText, results, _ := strings.Cut(input, ": ")
	cardNumber := getCardNumber(cardText)

	winningNumberPart, cardNumberPart, _ := strings.Cut(results, "|")
	winningNumbers := parseIntoNumbers(winningNumberPart)
	cardNumbers := parseIntoNumbers(cardNumberPart)

	return &Card{
		cardNumber,
		winningNumbers,
		cardNumbers,
	}
}

func getCardNumber(input string) int {
	cardDetails := strings.Split(input, " ")
	for _, detail := range cardDetails {
		cardNumber, err := strconv.Atoi(detail)
		if err == nil {
			return cardNumber
		}
	}

	panic("Couldn't find card number")
}
