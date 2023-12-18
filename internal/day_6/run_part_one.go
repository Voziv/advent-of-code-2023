package day_6

import (
	"fmt"
	"github.com/voziv/advent-of-code-2023/internal/util"
	"strconv"
	"strings"
)

func runPartOne(inputFileName string) string {
	lines := util.GetFileContents(inputFileName)
	times := parseNumbersFromInput(lines[0])
	distances := parseNumbersFromInput(lines[1])

	result := 1
	for i := 0; i < len(times); i++ {
		raceLength := times[i]
		distanceToBeat := distances[i]

		fmt.Printf("Race #%d. %dms, %dmm\n", i, raceLength, distanceToBeat)

		waysToBeat := 0

		for holdLength := 0; holdLength <= raceLength; holdLength++ {
			timeRemaining := raceLength - holdLength
			distance := holdLength * timeRemaining
			fmt.Printf("Hold for %dms, remaining time %dms, distance travelled %dmm.", holdLength, timeRemaining, distance)
			if distance > distanceToBeat {
				fmt.Println("We Win!")
				waysToBeat++
			} else {
				fmt.Println("We Lose!")
			}
		}

		result = result * waysToBeat

	}

	return strconv.Itoa(result)
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
