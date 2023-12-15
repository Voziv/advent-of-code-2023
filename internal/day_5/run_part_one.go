package day_5

import (
	"github.com/voziv/advent-of-code-2023/internal/util"
	"sort"
	"strconv"
	"strings"
)

const (
	SEEDS_HEADER                   = "seeds:"
	SEED_TO_SOIL_HEADER            = "seed-to-soil map:"
	SOIL_TO_FERTILIZER_HEADER      = "soil-to-fertilizer map:"
	FERTILIZER_TO_WATER_HEADER     = "fertilizer-to-water map:"
	WATER_TO_LIGHT_HEADER          = "water-to-light map:"
	LIGHT_TO_TEMPERATURE_HEADER    = "light-to-temperature map:"
	TEMPERATURE_TO_HUMIDITY_HEADER = "temperature-to-humidity map:"
	HUMIDITY_TO_LOCATION_HEADER    = "humidity-to-location map:"
)

var headings = []string{
	SEED_TO_SOIL_HEADER,
	SOIL_TO_FERTILIZER_HEADER,
	FERTILIZER_TO_WATER_HEADER,
	WATER_TO_LIGHT_HEADER,
	LIGHT_TO_TEMPERATURE_HEADER,
	TEMPERATURE_TO_HUMIDITY_HEADER,
	HUMIDITY_TO_LOCATION_HEADER,
}

func runPartOne(inputFileName string) string {
	lines := util.GetFileContents(inputFileName)

	var inputMaps = map[string][][]int{}
	var seeds []int
	var currentHeader = ""
	var foundHeading = false

	for _, line := range lines {
		foundHeading = false
		if strings.Contains(line, SEEDS_HEADER) {
			seeds = parseNumbersFromInput(line)
			continue
		}

		if line == "" {
			currentHeader = ""
			continue
		}

		for _, heading := range headings {
			if strings.Contains(line, heading) {
				currentHeader = heading
				foundHeading = true
				break
			}
		}

		if foundHeading {
			continue
		}

		inputMaps[currentHeader] = append(inputMaps[currentHeader], parseNumbersFromInput(line))
	}

	var locationNumbers []int

	for _, seed := range seeds {
		soilNumber := convertUsingMap(seed, inputMaps[SEED_TO_SOIL_HEADER])
		fertilizerNumber := convertUsingMap(soilNumber, inputMaps[SOIL_TO_FERTILIZER_HEADER])
		waterNumber := convertUsingMap(fertilizerNumber, inputMaps[FERTILIZER_TO_WATER_HEADER])
		lightNumber := convertUsingMap(waterNumber, inputMaps[WATER_TO_LIGHT_HEADER])
		temperatureNumber := convertUsingMap(lightNumber, inputMaps[LIGHT_TO_TEMPERATURE_HEADER])
		humidityNumber := convertUsingMap(temperatureNumber, inputMaps[TEMPERATURE_TO_HUMIDITY_HEADER])
		locationNumber := convertUsingMap(humidityNumber, inputMaps[HUMIDITY_TO_LOCATION_HEADER])
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

func convertUsingMap(number int, conversionRanges [][]int) int {
	for _, conversionRange := range conversionRanges {
		length := conversionRange[2]
		destinationRangeStart := conversionRange[0]
		// destinationRangeEnd := destinationRangeStart + length
		sourceRangeStart := conversionRange[1]
		sourceRangeEnd := sourceRangeStart + length

		if number >= sourceRangeStart && number <= sourceRangeEnd {
			return number - sourceRangeStart + destinationRangeStart
		}
	}

	return number
}
