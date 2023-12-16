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

type Conversion struct {
	destinationStart int
	sourceStart      int
	length           int
	modifier         int
}

func (c *Conversion) Length() int {
	return c.length
}

func NewConversion(destinationStart int, sourceStart int, length int) *Conversion {
	return &Conversion{
		destinationStart: destinationStart,
		sourceStart:      sourceStart,
		length:           length,
		modifier:         destinationStart - sourceStart,
	}
}

func NewConversionFromInput(input string) *Conversion {
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

func (c *Conversion) IsInRange(value int) bool {
	if value >= c.sourceStart && value <= c.sourceStart+c.length {
		return true
	}

	return false
}

func (c *Conversion) Modifier() int {
	return c.modifier
}
