package day_5

import (
	"fmt"
	"github.com/voziv/advent-of-code-2023/internal/util"
	"sort"
	"strconv"
	"strings"
)

type Conversion struct {
	length           int
	destinationStart int
	sourceStart      int
}

type SeedRange struct {
	start  int
	length int
}

func runPartTwo(inputFileName string) string {
	lines := util.GetFileContents(inputFileName)

	var inputMaps = map[int][][]int{}
	var seeds []int
	var currentCategory = -1

	for _, line := range lines {
		if strings.Contains(line, SeedsHeading) {
			seeds = expandSeeds(parseNumbersFromInput(line))
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

		inputMaps[currentCategory] = append(inputMaps[currentCategory], parseNumbersFromInput(line))
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
