package day_5

import (
	"github.com/voziv/advent-of-code-2023/internal/util"
	"sort"
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

const SeedsHeading = "seeds:"
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

func runPartOne(inputFileName string) string {
	lines := util.GetFileContents(inputFileName)

	var inputMaps = map[int][]*Converter{}
	var seeds []int
	var currentCategory = -1

	for _, line := range lines {
		if strings.Contains(line, SeedsHeading) {
			seeds = parseNumbersFromInput(line)
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

	var locationNumbers []int

	for _, seed := range seeds {
		soilNumber := convertUsingMap(seed, inputMaps[SeedToSoil])
		fertilizerNumber := convertUsingMap(soilNumber, inputMaps[SoilToFertilizer])
		waterNumber := convertUsingMap(fertilizerNumber, inputMaps[FertilizerToWater])
		lightNumber := convertUsingMap(waterNumber, inputMaps[WaterToLight])
		temperatureNumber := convertUsingMap(lightNumber, inputMaps[LightToTemperature])
		humidityNumber := convertUsingMap(temperatureNumber, inputMaps[TemperatureToHumidity])
		locationNumber := convertUsingMap(humidityNumber, inputMaps[HumidityToLocation])
		locationNumbers = append(locationNumbers, locationNumber)
	}

	sort.Slice(locationNumbers, func(i, j int) bool {
		return locationNumbers[i] < locationNumbers[j]
	})

	return strconv.Itoa(locationNumbers[0])
}

func convertUsingMap(number int, conversions []*Converter) int {
	for _, conversion := range conversions {
		if conversion.IsInRange(number) {
			return number + conversion.Modifier()
		}
	}

	return number
}
