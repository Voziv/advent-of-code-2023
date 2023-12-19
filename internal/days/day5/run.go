package day5

import (
	"fmt"
	"github.com/voziv/advent-of-code-2023/internal/util"
	"sort"
)

type result struct {
	partOne int
	partTwo int
}

func Run() {
	util.AssertResult("example.txt", run("./internal/days/day5/example.txt"), result{
		partOne: 35,
		partTwo: 46,
	})

	util.AssertResult("input.txt", run("./internal/days/day5/input.txt"), result{
		partOne: 251346198,
		partTwo: 72263011,
	})
}

func run(inputFileName string) result {
	lines := util.GetFileContents(inputFileName)

	result := result{
		partOne: 0,
		partTwo: 0,
	}

	seeds := parseNumbersFromInput(lines[0])
	rangeConverters := NewRangeConverterFromInput(lines[1:])

	var locationNumbers []int

	for _, seed := range seeds {
		soilNumber := convertUsingMap(seed, rangeConverters[SeedToSoil])
		fertilizerNumber := convertUsingMap(soilNumber, rangeConverters[SoilToFertilizer])
		waterNumber := convertUsingMap(fertilizerNumber, rangeConverters[FertilizerToWater])
		lightNumber := convertUsingMap(waterNumber, rangeConverters[WaterToLight])
		temperatureNumber := convertUsingMap(lightNumber, rangeConverters[LightToTemperature])
		humidityNumber := convertUsingMap(temperatureNumber, rangeConverters[TemperatureToHumidity])
		locationNumber := convertUsingMap(humidityNumber, rangeConverters[HumidityToLocation])
		locationNumbers = append(locationNumbers, locationNumber)
	}

	sort.Slice(locationNumbers, func(i, j int) bool {
		return locationNumbers[i] < locationNumbers[j]
	})

	result.partOne = locationNumbers[0]

	// Part Two
	seedRanges := NewSeedRangesFromInput(lines[0])

	//fmt.Printf("\nSeed Ranges\n")
	//printRanges(seedRanges)

	for _, seedRange := range seedRanges {
		ranges := []*SeedRange{seedRange}
		//fmt.Println("SeedToSoil")
		ranges = convertRange(ranges, rangeConverters[SeedToSoil])
		//fmt.Println("SoilToFertilizer")
		ranges = convertRange(ranges, rangeConverters[SoilToFertilizer])
		//fmt.Println("FertilizerToWater")
		ranges = convertRange(ranges, rangeConverters[FertilizerToWater])
		//fmt.Println("WaterToLight")
		ranges = convertRange(ranges, rangeConverters[WaterToLight])
		//fmt.Println("LightToTemperature")
		ranges = convertRange(ranges, rangeConverters[LightToTemperature])
		//fmt.Println("TemperatureToHumidity")
		ranges = convertRange(ranges, rangeConverters[TemperatureToHumidity])
		//fmt.Println("HumidityToLocation")
		ranges = convertRange(ranges, rangeConverters[HumidityToLocation])

		for _, r := range ranges {
			if r.start < result.partTwo || result.partTwo == 0 {
				result.partTwo = r.start
			}
		}
	}

	return result
}

func convertUsingMap(number int, conversions []*RangeConverter) int {
	for _, conversion := range conversions {
		if conversion.IsInRange(number) {
			return number + conversion.Modifier()
		}
	}

	return number
}

func printRanges(seedRanges []*SeedRange) {
	fmt.Printf("i: Start\tEnd\t\tDiff\n")
	for _, sr := range seedRanges {
		fmt.Printf("%d\t%d\n", sr.start, sr.end)
	}
}

func convertRange(ranges []*SeedRange, converters []*RangeConverter) []*SeedRange {
	//fmt.Printf("Comparing %d ranges against %d converters\n", len(ranges), len(converters))

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

	//fmt.Printf("New Ranges Count: %d\n", len(ranges))

	return ranges
}
