package util

import (
	"bufio"
	"os"
)

func GetFileContents(inputFileName string) []string {
	file, err := os.Open(inputFileName)
	if err != nil {
		panic(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)

	// All the problem wants is the highest total amount, as well as the top three totals summed.
	// We don't actually need to know which elf has what.
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// If any error has happened during our scanning, panic.
	// In theory scanner.Scan() has returned false by then?
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}
