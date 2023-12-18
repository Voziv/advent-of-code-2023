package day_5

// Map: 50 98 2
// Dest = 50
// Source = 98
// Length = 2
// Means 98-99 will be -48 in value

type RangeConverter struct {
	destinationStart int
	destinationEnd   int
	sourceStart      int
	sourceEnd        int
	length           int
	modifier         int
}

func NewRangeConverter(destinationStart int, sourceStart int, length int) *RangeConverter {
	return &RangeConverter{
		destinationStart: destinationStart,
		destinationEnd:   destinationStart + length,
		sourceStart:      sourceStart,
		sourceEnd:        sourceStart + length,
		length:           length,
		modifier:         destinationStart - sourceStart,
	}
}

func (c *RangeConverter) IsInRange(value int) bool {
	if value >= c.sourceStart && value <= c.sourceStart+c.length {
		return true
	}

	return false
}

func (c *RangeConverter) Modifier() int {
	return c.modifier
}
