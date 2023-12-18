package day_5

import (
	"fmt"
	"github.com/voziv/advent-of-code-2023/internal/util"
	"strconv"
)

func runPartTwo(inputFileName string) string {
	lines := util.GetFileContents(inputFileName)

	seedRanges := NewSeedRangesFromInput(lines[0])

	fmt.Printf("\nSeed Ranges\n")
	printRanges(seedRanges)

	rangeConverters := NewRangeConverterFromInput(lines[1:])

	fmt.Printf("\nConverting to soil ranges\n")
	fmt.Println("SeedToSoil")

	var lowestLocation = 0

	for _, seedRange := range seedRanges {
		ranges := []*SeedRange{seedRange}
		ranges = convertRange(ranges, rangeConverters[SeedToSoil])
		fmt.Println("SoilToFertilizer")
		ranges = convertRange(ranges, rangeConverters[SoilToFertilizer])
		fmt.Println("FertilizerToWater")
		ranges = convertRange(ranges, rangeConverters[FertilizerToWater])
		fmt.Println("WaterToLight")
		ranges = convertRange(ranges, rangeConverters[WaterToLight])
		fmt.Println("LightToTemperature")
		ranges = convertRange(ranges, rangeConverters[LightToTemperature])
		fmt.Println("TemperatureToHumidity")
		ranges = convertRange(ranges, rangeConverters[TemperatureToHumidity])
		fmt.Println("HumidityToLocation")
		ranges = convertRange(ranges, rangeConverters[HumidityToLocation])

		for _, r := range ranges {
			if r.start < lowestLocation || lowestLocation == 0 {
				lowestLocation = r.start
			}
		}

	}

	// 12170500 was too low (12,170,500)
	return strconv.Itoa(lowestLocation)
	//return "46"
}

func printRanges(seedRanges []*SeedRange) {
	fmt.Printf("i: Start\tEnd\t\tDiff\n")
	for _, sr := range seedRanges {
		fmt.Printf("%d\t%d\n", sr.start, sr.end)
	}
}

func convertRange(ranges []*SeedRange, converters []*RangeConverter) []*SeedRange {
	fmt.Printf("Comparing %d ranges against %d converters\n", len(ranges), len(converters))

	var appliedRanges []*SeedRange
	for _, converter := range converters {
		var newRanges []*SeedRange

		for _, seedRange := range ranges {
			before := []int{seedRange.start, min(seedRange.end, converter.sourceStart)}
			middle := []int{max(seedRange.start, converter.sourceStart), min(converter.sourceEnd, seedRange.end)}
			after := []int{max(converter.sourceEnd, seedRange.start), seedRange.end}

			if before[1] > before[0] {
				newRanges = append(newRanges, NewSeedRange(before[0], before[1]))
			}

			if middle[1] > middle[0] {
				appliedRanges = append(appliedRanges, NewSeedRange(middle[0]+converter.modifier, middle[1]+converter.modifier))
			}

			if after[1] > after[0] {
				newRanges = append(newRanges, NewSeedRange(after[0], after[1]))
			}
		}

		ranges = newRanges
	}

	ranges = append(ranges, appliedRanges...)

	fmt.Printf("New Ranges Count: %d\n", len(ranges))

	return ranges
}
