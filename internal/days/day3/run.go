package day3

import (
	"fmt"
	"github.com/voziv/advent-of-code-2023/internal/util"
	"strconv"
)

type result struct {
	partOne int
	partTwo int
}

func Run() {

	util.AssertResult("example.txt", run("./internal/days/day3/example.txt"), result{
		partOne: 4361,
		partTwo: 467835,
	})
	util.AssertResult("input.txt", run("./internal/days/day3/input.txt"), result{
		partOne: 540025,
		partTwo: 84584891,
	})
}

func run(inputFileName string) result {
	lines := util.GetFileContents(inputFileName)

	result := result{
		partOne: 0,
		partTwo: 0,
	}

	// Part One
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
					result.partOne += castToNumber(currentNumber)
				}
				currentNumber = ""
				isAdjacent = false
			}

		}
	}

	//
	// Part Two
	//
	var gears = map[int]map[int][]int{}
	currentNumber = ""
	var adjacentGears = map[int]map[int]bool{}
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			currentCharacter := lines[y][x : x+1]
			if isDigit(currentCharacter) {
				currentNumber += currentCharacter

				for _, position := range getAdjacentGears(y, x, lines) {
					if _, ok := adjacentGears[position.y]; !ok {
						adjacentGears[position.y] = map[int]bool{}
					}
					adjacentGears[position.y][position.x] = true
				}

			} else {
				if len(adjacentGears) > 0 && currentNumber != "" {
					number := castToNumber(currentNumber)

					for gearY, exes := range adjacentGears {
						for gearX, _ := range exes {
							if _, ok := gears[gearY]; !ok {
								gears[gearY] = map[int][]int{}
							}

							//fmt.Printf("Character y:%d, x:%d, Adding %d to gearY:%d gearX:%d\n", x, y, number, gearY, gearX)
							gears[gearY][gearX] = append(gears[gearY][gearX], number)
						}
					}
				}

				currentNumber = ""
				adjacentGears = map[int]map[int]bool{}
			}
		}
	}

	for _, gearMap := range gears {
		for _, gearRatios := range gearMap {
			//fmt.Printf("%d %d %s\n", y, x, gearRatios)
			if len(gearRatios) == 2 {
				result.partTwo += gearRatios[0] * gearRatios[1]
			}
		}
	}

	return result
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

type Position struct {
	x int
	y int
}

func getAdjacentGears(y int, x int, lines []string) []*Position {
	var positionsToCheck = []*Position{
		// Character before
		{x - 1, y},
		// Character after
		{x + 1, y},
		// Above
		{x - 1, y - 1},
		{x, y - 1},
		{x + 1, y - 1},
		// Below
		{x - 1, y + 1},
		{x, y + 1},
		{x + 1, y + 1},
	}
	var positions = []*Position{}

	for _, position := range positionsToCheck {
		if positionIsGear(position, lines) {
			positions = append(positions, position)
		}

	}

	return positions
}

func positionIsGear(position *Position, lines []string) bool {
	if position.y < 0 || position.y >= len(lines) {
		return false
	}

	if position.x < 0 || position.x >= len(lines[position.y]) {
		return false
	}

	character := lines[position.y][position.x : position.x+1]

	if character == "*" {
		return true
	}

	return false
}
