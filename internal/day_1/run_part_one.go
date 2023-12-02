package day_1

import (
	"errors"
	"github.com/voziv/advent-of-code-2023/internal/util"
	"strconv"
)

func runPartOne(inputFileName string) string {
	lines := util.GetFileContents(inputFileName)

	sum := 0

	for _, line := range lines {

		value, err := parseCalibrationDigits(line)
		if err != nil {
			panic(err)
		}

		sum += value
	}

	return strconv.Itoa(sum)
}

func parseCalibrationDigits(input string) (int, error) {
	firstNumber, err := findFirstDigit(input)
	if err != nil {
		return 0, errors.Join(errors.New("couldn't find first number"), err)
	}

	secondNumber, err := findLastDigit(input)
	if err != nil {
		return 0, errors.Join(errors.New("couldn't find second number"), err)
	}

	calibrationDigits, err := strconv.Atoi(firstNumber + secondNumber)
	if err != nil {
		panic("Calibration number could not be cast to an int")
	}

	return calibrationDigits, nil
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
