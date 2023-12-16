package day_5

import (
	"github.com/voziv/advent-of-code-2023/internal/util"
	"sort"
	"strconv"
	"strings"
)

const (
	SEED_TO_SOIL            = 1
	SOIL_TO_FERTILIZER      = 2
	FERTILIZER_TO_WATER     = 3
	WATER_TO_LIGHT          = 4
	LIGHT_TO_TEMPERATURE    = 5
	TEMPERATURE_TO_HUMIDITY = 6
	HUMIDITY_TO_LOCATION    = 7
)

const SeedsHeading = "seeds:"
const MapHeading = "map:"

var mapHeadingsToCategories = map[string]int{
	"seed-to-soil":            SEED_TO_SOIL,
	"soil-to-fertilizer":      SOIL_TO_FERTILIZER,
	"fertilizer-to-water":     FERTILIZER_TO_WATER,
	"water-to-light":          WATER_TO_LIGHT,
	"light-to-temperature":    LIGHT_TO_TEMPERATURE,
	"temperature-to-humidity": TEMPERATURE_TO_HUMIDITY,
	"humidity-to-location":    HUMIDITY_TO_LOCATION,
}

func runPartOne(inputFileName string) string {
	lines := util.GetFileContents(inputFileName)

	var inputMaps = map[int][]*Conversion{}
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
		soilNumber := convertUsingMap(seed, inputMaps[SEED_TO_SOIL])
		fertilizerNumber := convertUsingMap(soilNumber, inputMaps[SOIL_TO_FERTILIZER])
		waterNumber := convertUsingMap(fertilizerNumber, inputMaps[FERTILIZER_TO_WATER])
		lightNumber := convertUsingMap(waterNumber, inputMaps[WATER_TO_LIGHT])
		temperatureNumber := convertUsingMap(lightNumber, inputMaps[LIGHT_TO_TEMPERATURE])
		humidityNumber := convertUsingMap(temperatureNumber, inputMaps[TEMPERATURE_TO_HUMIDITY])
		locationNumber := convertUsingMap(humidityNumber, inputMaps[HUMIDITY_TO_LOCATION])
		locationNumbers = append(locationNumbers, locationNumber)
	}

	sort.Slice(locationNumbers, func(i, j int) bool {
		return locationNumbers[i] < locationNumbers[j]
	})

	return strconv.Itoa(locationNumbers[0])
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

func convertUsingMap(number int, conversions []*Conversion) int {
	for _, conversion := range conversions {
		if conversion.IsInRange(number) {
			return number + conversion.Modifier()
		}
	}

	return number
}
