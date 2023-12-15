package day_5

import (
	"fmt"
	"github.com/voziv/advent-of-code-2023/internal/util"
	"sort"
	"strconv"
	"strings"
)

func runPartTwo(inputFileName string) string {
	lines := util.GetFileContents(inputFileName)

	var inputMaps = map[string][][]int{}
	var seeds []int
	var currentHeader = ""
	var foundHeading = false

	for _, line := range lines {
		foundHeading = false
		if strings.Contains(line, SEEDS_HEADER) {
			seeds = expandSeeds(parseNumbersFromInput(line))
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

func expandSeeds(seeds []int) []int {
	expandedSeeds := []int{}
	fmt.Println(seeds)
	for i := 0; i < len(seeds); i += 2 {
		fmt.Printf("Seed Start: %d, Seed Range: %d\n", seeds[i], seeds[i+1])
		for j := 0; j < seeds[i+1]; j++ {
			expandedSeeds = append(expandedSeeds, seeds[i]+j)
		}
	}

	return expandedSeeds
}
