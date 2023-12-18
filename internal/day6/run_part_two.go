package day6

import (
	"fmt"
	"github.com/voziv/advent-of-code-2023/internal/util"
	"strconv"
)

func runPartTwo(inputFileName string) string {
	lines := util.GetFileContents(inputFileName)
	raceLength := combineDigitsOfInts(parseNumbersFromInput(lines[0]))
	distanceToBeat := combineDigitsOfInts(parseNumbersFromInput(lines[1]))

	fmt.Printf("Race: %dms, %dmm\n", raceLength, distanceToBeat)

	waysToBeat := 0

	for holdLength := 0; holdLength <= raceLength; holdLength++ {
		timeRemaining := raceLength - holdLength
		distance := holdLength * timeRemaining
		if distance > distanceToBeat {
			waysToBeat++
		}
	}

	return strconv.Itoa(waysToBeat)
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
