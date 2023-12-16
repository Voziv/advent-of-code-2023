package day_5

import (
	"fmt"
	"github.com/voziv/advent-of-code-2023/internal/util"
	"strconv"
	"strings"
)

func runPartTwo(inputFileName string) string {
	lines := util.GetFileContents(inputFileName)

	var inputMaps = map[int][]*Converter{}
	var seedRanges []*SeedRange
	var currentCategory = -1

	for _, line := range lines {
		if strings.Contains(line, SeedsHeading) {
			seedRanges = NewSeedRangesFromInput(line)
			for i, seedRange := range seedRanges {
				fmt.Printf("#%d: %d to %d\n", i, seedRange.start, seedRange.end)
			}
			continue
		} else if strings.Contains(line, MapHeading) {
			for heading, category := range mapHeadingsToCategories {
				if strings.Contains(line, heading) {
					currentCategory = category
					break
				}
			}

			if currentCategory == -1 {
				panic("No heading found for line " + line)
			}

			continue
		} else if line == "" {
			currentCategory = -1
			continue
		}

		inputMaps[currentCategory] = append(inputMaps[currentCategory], NewConversionFromInput(line))
	}

	fmt.Println("SeedToSoil")
	ranges := convertRange(seedRanges, inputMaps[SeedToSoil])
	fmt.Println("SoilToFertilizer")
	ranges = convertRange(ranges, inputMaps[SoilToFertilizer])
	fmt.Println("FertilizerToWater")
	ranges = convertRange(ranges, inputMaps[FertilizerToWater])
	fmt.Println("WaterToLight")
	ranges = convertRange(ranges, inputMaps[WaterToLight])
	fmt.Println("LightToTemperature")
	ranges = convertRange(ranges, inputMaps[LightToTemperature])
	fmt.Println("TemperatureToHumidity")
	ranges = convertRange(ranges, inputMaps[TemperatureToHumidity])
	fmt.Println("HumidityToLocation")
	ranges = convertRange(ranges, inputMaps[HumidityToLocation])

	var lowestLocation = 0
	for _, r := range ranges {
		if r.start < lowestLocation || lowestLocation == 0 {
			lowestLocation = r.start
		}
	}

	return strconv.Itoa(lowestLocation)
}

func convertRange(ranges []*SeedRange, converters []*Converter) []*SeedRange {
	fmt.Printf("Comparing %d ranges against %d converters\n", len(ranges), len(converters))

	var newRanges []*SeedRange
	for convIndex, c := range converters {
		var rangesToIterate []*SeedRange
		copy(rangesToIterate, ranges)
		fmt.Printf("Ranges to iterate: %d\n", len(rangesToIterate))

		for i, r := range rangesToIterate {
			fmt.Printf("Range: %d to %d\t|\tConv#%d: %d to %d\t m:%d = \t", r.start, r.end, convIndex, c.sourceStart, c.sourceEnd, c.modifier)
			// Range does not match the converter
			if c.sourceStart > r.end || c.sourceEnd < r.start {
				fmt.Println("Range does not match converter")
				continue
			}

			if r.start >= c.sourceStart && r.end <= c.sourceEnd {
				// Range is entirely within the converter
				fmt.Println("Whole range fits the converter")
				newRanges = append(newRanges, NewSeedRange(r.start+c.modifier, r.end+c.modifier))
			} else if r.start >= c.sourceStart && r.end < c.sourceEnd {
				// Converter only overlaps start of range
				fmt.Println("Overlaps beginning of range")
				newRanges = append(newRanges, NewSeedRange(r.start+c.modifier, c.destinationEnd))
				ranges = append(ranges, NewSeedRange(c.sourceEnd+1, r.end))
			} else if r.start < c.sourceStart && r.end <= c.sourceEnd {
				// Converter only overlaps the end of range
				fmt.Println("Overlaps end of range")

				newRanges = append(newRanges, NewSeedRange(c.destinationStart, r.end+c.modifier))
				ranges = append(ranges, NewSeedRange(r.start, c.sourceStart-1))
			} else {
				// Converter is entirely within the range
				fmt.Println("Overlaps middle of range")
				ranges = append(ranges, NewSeedRange(r.start, c.sourceStart-1))
				newRanges = append(newRanges, NewSeedRange(c.destinationStart, c.destinationEnd))
				ranges = append(ranges, NewSeedRange(c.sourceEnd+1, r.end))
			}

			ranges = remove(ranges, i)
		}
	}

	for _, r := range ranges {
		fmt.Printf("Appending leftover range %d to %d\n", r.start, r.end)
		newRanges = append(newRanges, r)
	}

	fmt.Printf("New Ranges Count: %d\n", len(newRanges))

	return newRanges
}

func remove(slice []*SeedRange, s int) []*SeedRange {
	return append(slice[:s], slice[s+1:]...)
}
