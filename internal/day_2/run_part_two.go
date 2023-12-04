package day_2

import (
	"github.com/voziv/advent-of-code-2023/internal/util"
	"strconv"
)

func runPartTwo(inputFileName string) string {

	lines := util.GetFileContents(inputFileName)
	sum := 0

	for _, line := range lines {
		game := parseGameFromString(line)

		sum += getPowerOfGame(game)
	}

	return strconv.Itoa(sum)
}

func getPowerOfGame(game *Game) int {
	minimumNumberOfCubes := map[string]int{
		"blue":  0,
		"green": 0,
		"red":   0,
	}

	for _, round := range game.Rounds {
		for color, count := range round.Cubes {
			if count > minimumNumberOfCubes[color] {
				minimumNumberOfCubes[color] = count
			}
		}
	}

	return minimumNumberOfCubes["blue"] * minimumNumberOfCubes["green"] * minimumNumberOfCubes["red"]
}
