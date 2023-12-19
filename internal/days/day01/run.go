package day01

import (
	"errors"
	"github.com/voziv/advent-of-code-2023/internal/util"
	"strconv"
)

type result struct {
	partOne int
	partTwo int
}

var wordsAsNumbers = []string{
	"one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
}

var wordReplacements = map[string]string{
	"one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9",
}

func Run() {
	util.AssertResult("example.txt", run("./internal/days/day01/example.txt"), result{
		partOne: 142,
		partTwo: 142,
	})

	util.AssertResult("example2.txt", run("./internal/days/day01/example2.txt"), result{
		partOne: 209,
		partTwo: 281,
	})

	util.AssertResult("input.txt", run("./internal/days/day01/input.txt"), result{
		partOne: 54953,
		partTwo: 53868,
	})
}

func run(inputFileName string) result {
	lines := util.GetFileContents(inputFileName)

	result := result{
		partOne: 0,
		partTwo: 0,
	}

	for _, line := range lines {
		partOneValue := parseCalibrationDigits(line)
		result.partOne += partOneValue

		partTwoValue := parseCalibrationNumber(line)
		result.partTwo += partTwoValue

		//fmt.Printf("%-20s\tP1: %d\t P2: %d\n", line, partOneValue, partTwoValue)
	}

	return result
}

func parseCalibrationDigits(input string) int {
	firstNumber, err := findFirstDigit(input)
	if err != nil {
		return 0
	}

	secondNumber, err := findLastDigit(input)
	if err != nil {
		return 0
	}

	calibrationDigits, err := strconv.Atoi(firstNumber + secondNumber)
	if err != nil {
		panic("Calibration number could not be cast to an int")
	}

	return calibrationDigits
}

func findFirstDigit(input string) (string, error) {
	for i := 0; i < len(input); i++ {
		_, err := strconv.Atoi(input[i : i+1])
		if err == nil {
			return input[i : i+1], nil
		}
	}

	return "", errors.New("could not find a number in the string")
}

func findLastDigit(input string) (string, error) {
	for i := len(input) - 1; i >= 0; i-- {
		_, err := strconv.Atoi(input[i : i+1])
		if err == nil {
			return input[i : i+1], nil
		}
	}

	return "", errors.New("could not find a number in the string")
}

func parseCalibrationNumber(input string) int {
	firstNumber, err := findFirstNumber(input)
	if err != nil {
		return 0
	}

	secondNumber, err := findLastNumber(input)
	if err != nil {
		return 0
	}

	calibrationNumber, err := strconv.Atoi(firstNumber + secondNumber)
	if err != nil {
		panic("Calibration number could not be cast to an int")
	}

	return calibrationNumber
}

func findFirstNumber(input string) (string, error) {
	for i := 0; i < len(input); i++ {
		_, err := strconv.Atoi(input[i : i+1])
		if err == nil {
			return input[i : i+1], nil
		}

		for _, word := range wordsAsNumbers {
			// Ensure we have enough characters left to fit the word
			if len(word) > len(input)-i {
				continue
			}
			matches := true
			for j := 0; j < len(word); j++ {
				if input[i+j] != word[j] {
					matches = false
					break
				}
			}

			if !matches {
				continue
			}

			return wordReplacements[word], nil
		}
	}

	return "", errors.New("could not find a number in the string")
}

func findLastNumber(input string) (string, error) {
	for i := len(input) - 1; i >= 0; i-- {
		_, err := strconv.Atoi(input[i : i+1])
		if err == nil {
			return input[i : i+1], nil
		}

		for _, word := range wordsAsNumbers {
			// Ensure we have enough characters left to fit the word
			if len(word) > len(input)-i {
				continue
			}
			matches := true
			for j := 0; j < len(word); j++ {
				if input[i+j] != word[j] {
					matches = false
					break
				}
			}

			if !matches {
				continue
			}

			return wordReplacements[word], nil
		}
	}

	return "", errors.New("could not find a number in the string")
}
