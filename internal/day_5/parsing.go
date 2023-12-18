package day_5

import (
	"strconv"
	"strings"
)

const (
	SeedToSoil            = 1
	SoilToFertilizer      = 2
	FertilizerToWater     = 3
	WaterToLight          = 4
	LightToTemperature    = 5
	TemperatureToHumidity = 6
	HumidityToLocation    = 7
)

const MapHeading = "map:"

var mapHeadingsToCategories = map[string]int{
	"seed-to-soil":            SeedToSoil,
	"soil-to-fertilizer":      SoilToFertilizer,
	"fertilizer-to-water":     FertilizerToWater,
	"water-to-light":          WaterToLight,
	"light-to-temperature":    LightToTemperature,
	"temperature-to-humidity": TemperatureToHumidity,
	"humidity-to-location":    HumidityToLocation,
}

func NewSeedRangesFromInput(input string) []*SeedRange {
	var seedRanges []*SeedRange
	numbers := parseNumbersFromInput(input)
	for i := 0; i < len(numbers); i += 2 {
		sr := NewSeedRange(numbers[i], numbers[i]+numbers[i+1])
		seedRanges = append(seedRanges, sr)
	}

	return seedRanges
}

func NewRangeConverterFromInput(inputs []string) map[int][]*RangeConverter {
	var rangeConverters = map[int][]*RangeConverter{}
	var converterType = -1
	for _, input := range inputs {
		if strings.Contains(input, MapHeading) {
			for heading, category := range mapHeadingsToCategories {
				if strings.Contains(input, heading) {
					converterType = category
					break
				}
			}

			if converterType == -1 {
				panic("No heading found for line " + input)
			}

			continue
		} else if input == "" {
			converterType = -1
			continue
		}

		numbers := parseNumbersFromInput(input)
		if len(numbers) != 3 {
			panic("Input was not a valid conversion map: " + input)
		}

		rangeConverters[converterType] = append(rangeConverters[converterType], NewRangeConverter(numbers[0], numbers[1], numbers[2]))
	}

	return rangeConverters
}

func parseNumbersFromInput(input string) []int {
	var numbers []int
	tokens := strings.Split(input, " ")
	for _, token := range tokens {
		number, err := strconv.Atoi(token)
		if err != nil {
			continue
		}

		numbers = append(numbers, number)
	}

	return numbers
}
