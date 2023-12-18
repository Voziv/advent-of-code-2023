package day3

import (
	"fmt"
	"github.com/voziv/advent-of-code-2023/internal/util"
	"strconv"
)

type Position struct {
	x int
	y int
}

func runPartTwo(inputFileName string) string {
	lines := util.GetFileContents(inputFileName)

	var gears = map[int]map[int][]int{}
	currentNumber := ""
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

							fmt.Printf("Character y:%d, x:%d, Adding %d to gearY:%d gearX:%d\n", x, y, number, gearY, gearX)
							gears[gearY][gearX] = append(gears[gearY][gearX], number)
						}
					}
				}

				currentNumber = ""
				adjacentGears = map[int]map[int]bool{}
			}
		}
	}

	sum := 0
	for y, gearMap := range gears {
		for x, gearRatios := range gearMap {
			fmt.Printf("%d %d %s\n", y, x, gearRatios)
			if len(gearRatios) == 2 {
				sum += gearRatios[0] * gearRatios[1]
			}
		}
	}

	return strconv.Itoa(sum)
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
