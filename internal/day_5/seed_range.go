package day_5

import "sort"

type SeedRange struct {
	start  int
	length int
	end    int
}

func NewSeedRange(start int, length int) *SeedRange {
	return &SeedRange{start: start, length: length, end: start + length}
}

func NewSeedRangesFromInput(input string) []*SeedRange {
	var seedRanges []*SeedRange
	numbers := parseNumbersFromInput(input)
	for i := 0; i < len(numbers); i += 2 {
		seedRanges = append(seedRanges, NewSeedRange(numbers[i], numbers[i+1]))
	}
	sort.Slice(seedRanges, func(i, j int) bool {
		return seedRanges[i].start < seedRanges[j].start
	})
	return seedRanges
}
