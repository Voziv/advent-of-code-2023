package day_5

import (
	"strconv"
	"strings"
)

// Map: 50 98 2
// Dest = 50
// Source = 98
// Length = 2
// Means 98-99 will be -48 in value

type Converter struct {
	destinationStart int
	destinationEnd   int
	sourceStart      int
	sourceEnd        int
	length           int
	modifier         int
}

func (c *Converter) Length() int {
	return c.length
}

func NewConversion(destinationStart int, sourceStart int, length int) *Converter {
	return &Converter{
		destinationStart: destinationStart,
		destinationEnd:   destinationStart + length,
		sourceStart:      sourceStart,
		sourceEnd:        sourceStart + length,
		length:           length,
		modifier:         destinationStart - sourceStart,
	}
}

func NewConversionFromInput(input string) *Converter {
	var numbers []int
	tokens := strings.Split(input, " ")
	for _, token := range tokens {
		number, err := strconv.Atoi(token)
		if err != nil {
			continue
		}

		numbers = append(numbers, number)
	}

	if len(numbers) != 3 {
		panic("Input was not a valid conversion map: " + input)
	}

	return NewConversion(numbers[0], numbers[1], numbers[2])
}

func (c *Converter) IsInRange(value int) bool {
	if value >= c.sourceStart && value <= c.sourceStart+c.length {
		return true
	}

	return false
}

func (c *Converter) Modifier() int {
	return c.modifier
}
