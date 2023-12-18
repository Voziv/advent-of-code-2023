package day5

import "fmt"

type SeedRange struct {
	key   string
	start int
	end   int
}

func NewSeedRange(start int, end int) *SeedRange {
	return &SeedRange{
		key:   fmt.Sprintf("%d-%d", start, end),
		start: start,
		end:   end,
	}
}
