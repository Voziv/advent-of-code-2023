package day_3

import (
	"fmt"
	"github.com/voziv/advent-of-code-2023/internal/util"
	"strconv"
)

func isPositionAdjacentToSymbol(i int, j int, lines []string) bool {
	// Character before
	if positionIsSymbol(i, j-1, lines) {
		return true
	}
	// Character after
	if positionIsSymbol(i, j+1, lines) {
		return true
	}
	// Character Top Left
	if positionIsSymbol(i-1, j-1, lines) {
		return true
	}
	// Character Top Middle
	if positionIsSymbol(i-1, j, lines) {
		return true
	}
	// Character Top Right
	if positionIsSymbol(i-1, j+1, lines) {
		return true
	}
	// Character Bottom Left
	if positionIsSymbol(i+1, j-1, lines) {
		return true
	}
	// Character Bottom Middle
	if positionIsSymbol(i+1, j, lines) {
		return true
	}
	// Character Bottom Right
	if positionIsSymbol(i+1, j+1, lines) {
		return true
	}

	return false
}
func positionIsSymbol(i int, j int, lines []string) bool {
	if i < 0 || i >= len(lines) {
		return false
	}

	if j < 0 || j >= len(lines[i]) {
		return false
	}

	character := lines[i][j : j+1]

	if character == "." {
		return false
	}

	if isDigit(character) {
		return false
	}

	return true
}

func runPartOne(inputFileName string) string {
	lines := util.GetFileContents(inputFileName)

	sum := 0
	currentNumber := ""
	isAdjacent := false
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			currentCharacter := lines[i][j : j+1]
			if isDigit(currentCharacter) {
				currentNumber += currentCharacter
				if !isAdjacent {
					isAdjacent = isPositionAdjacentToSymbol(i, j, lines)
				}
			} else {
				if isAdjacent && currentNumber != "" {
					sum += castToNumber(currentNumber)
				}
				currentNumber = ""
				isAdjacent = false
			}

		}
	}

	return strconv.Itoa(sum)
}

func isDigit(input string) bool {
	_, err := strconv.Atoi(input)
	return err == nil
}

func castToNumber(input string) int {
	number, err := strconv.Atoi(input)
	if err != nil {
		fmt.Printf("ERROR: %s is NaN\n", input)
		panic(err)
	}
	return number
}
