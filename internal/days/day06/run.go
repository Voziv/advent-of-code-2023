package day06

import (
	"github.com/voziv/advent-of-code-2023/internal/util"
	"strconv"
	"strings"
)

type result struct {
	partOne int
	partTwo int
}

func Run() {
	util.AssertResult("example.txt", run("./internal/days/day06/example.txt"), result{
		partOne: 288,
		partTwo: 71503,
	})

	util.AssertResult("input.txt", run("./internal/days/day06/input.txt"), result{
		partOne: 4568778,
		partTwo: 28973936,
	})
}

func run(inputFileName string) result {
	lines := util.GetFileContents(inputFileName)
	times := parseNumbersFromInput(lines[0])
	distances := parseNumbersFromInput(lines[1])

	result := result{
		partOne: 1,
		partTwo: 0,
	}

	for i := 0; i < len(times); i++ {
		raceLength := times[i]
		distanceToBeat := distances[i]

		//fmt.Printf("Race #%d. %dms, %dmm\n", i, raceLength, distanceToBeat)

		waysToBeat := 0

		for holdLength := 0; holdLength <= raceLength; holdLength++ {
			timeRemaining := raceLength - holdLength
			distance := holdLength * timeRemaining
			//fmt.Printf("Hold for %dms, remaining time %dms, distance travelled %dmm.", holdLength, timeRemaining, distance)
			if distance > distanceToBeat {
				//fmt.Println("We Win!")
				waysToBeat++
			} else {
				//fmt.Println("We Lose!")
			}
		}

		result.partOne *= waysToBeat

	}

	// Part Two
	raceLength := combineDigitsOfInts(parseNumbersFromInput(lines[0]))
	distanceToBeat := combineDigitsOfInts(parseNumbersFromInput(lines[1]))

	//fmt.Printf("Race: %dms, %dmm\n", raceLength, distanceToBeat)

	for holdLength := 0; holdLength <= raceLength; holdLength++ {
		timeRemaining := raceLength - holdLength
		distance := holdLength * timeRemaining
		if distance > distanceToBeat {
			result.partTwo++
		}
	}

	return result
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

func combineDigitsOfInts(ints []int) int {
	stringInt := ""
	for _, value := range ints {
		stringInt += strconv.Itoa(value)
	}
	number, err := strconv.Atoi(stringInt)
	if err != nil {
		panic("Couldn't cast " + stringInt + " to integer")
	}
	return number
}
