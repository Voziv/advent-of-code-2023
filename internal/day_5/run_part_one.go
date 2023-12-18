package day_5

import (
	"github.com/voziv/advent-of-code-2023/internal/util"
	"sort"
	"strconv"
)

func runPartOne(inputFileName string) string {
	lines := util.GetFileContents(inputFileName)

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

	return strconv.Itoa(locationNumbers[0])
}

func convertUsingMap(number int, conversions []*RangeConverter) int {
	for _, conversion := range conversions {
		if conversion.IsInRange(number) {
			return number + conversion.Modifier()
		}
	}

	return number
}
