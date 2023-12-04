package day_2

import (
	"github.com/voziv/advent-of-code-2023/internal/util"
	"strconv"
	"strings"
)

type Game struct {
	Id     int
	Rounds []Round
}

type Round struct {
	Cubes map[string]int
}

var bagOfCubes = map[string]int{
	"blue":  14,
	"green": 13,
	"red":   12,
}

func runPartOne(inputFileName string) string {

	lines := util.GetFileContents(inputFileName)
	sum := 0

	for _, line := range lines {
		game := parseGameFromString(line)

		isValidGame := validateGame(game)
		if isValidGame {
			sum += game.Id
		}
	}

	return strconv.Itoa(sum)
}

func validateGame(game *Game) bool {
	isValid := true

	for _, round := range game.Rounds {
		for color, count := range round.Cubes {
			if count > bagOfCubes[color] {
				isValid = false
				break
			}
		}
	}

	return isValid
}

func parseGameFromString(input string) *Game {
	gamePart, afterGamePart, found := strings.Cut(input, ": ")
	if !found {
		panic("Ahhhh no : found on line of input")
	}
	_, gameNumberString, _ := strings.Cut(gamePart, " ")

	gameNumber, err := strconv.Atoi(gameNumberString)
	if err != nil {
		panic(err)
	}

	roundStrings := strings.Split(afterGamePart, "; ")

	var rounds []Round

	for _, roundString := range roundStrings {
		cubes := map[string]int{
			"blue":  0,
			"green": 0,
			"red":   0,
		}

		cubeStrings := strings.Split(roundString, ", ")
		for _, cubeString := range cubeStrings {
			numberString, color, _ := strings.Cut(cubeString, " ")
			numberOfCubes, _ := strconv.Atoi(numberString)
			cubes[color] = numberOfCubes
		}

		rounds = append(rounds, Round{
			Cubes: cubes,
		})
	}

	return &Game{
		Id:     gameNumber,
		Rounds: rounds,
	}
}
