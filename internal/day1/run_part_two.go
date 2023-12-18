package day1

import (
	"errors"
	"github.com/voziv/advent-of-code-2023/internal/util"
	"strconv"
)

var wordsAsNumbers = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

var wordReplacements = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func runPartTwo(inputFileName string) string {
	lines := util.GetFileContents(inputFileName)

	sum := 0

	for _, line := range lines {

		value, err := parseCalibrationNumber(line)
		if err != nil {
			panic(err)
		}

		sum += value
	}

	return strconv.Itoa(sum)
}

func parseCalibrationNumber(input string) (int, error) {
	firstNumber, err := findFirstNumber(input)
	if err != nil {
		return 0, errors.Join(errors.New("couldn't find first number"), err)
	}

	secondNumber, err := findLastNumber(input)
	if err != nil {
		return 0, errors.Join(errors.New("couldn't find second number"), err)
	}

	calibrationNumber, err := strconv.Atoi(firstNumber + secondNumber)
	if err != nil {
		panic("Calibration number could not be cast to an int")
	}

	return calibrationNumber, nil
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
