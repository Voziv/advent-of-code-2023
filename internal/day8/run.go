package day8

import (
	"github.com/voziv/advent-of-code-2023/internal/util"
	"strings"
)

type result struct {
	partOne int
	partTwo int
}

func Run() {
	util.AssertResult("example.txt", run("./internal/day8/example.txt"), result{
		partOne: 2,
		partTwo: 0,
	})

	util.AssertResult("example2.txt", run("./internal/day8/example2.txt"), result{
		partOne: 6,
		partTwo: 0,
	})

	util.AssertResult("input.txt", run("./internal/day8/input.txt"), result{
		partOne: 19637,
		partTwo: 0,
	})
}

func run(inputFileName string) result {
	result := result{
		partOne: 0,
		partTwo: 0,
	}
	lines := util.GetFileContents(inputFileName)

	directions := strings.Split(lines[0], "")
	nodeLines := lines[2:]

	var nodes = map[string]struct {
		left  string
		right string
	}{}

	for _, line := range nodeLines {
		// AAA = (BBB, CCC)
		id := line[:3]
		left := line[7:10]
		right := line[12:15]

		nodes[id] = struct {
			left  string
			right string
		}{left: left, right: right}
	}

	foundZZZ := false
	currentPosition := "AAA"

	for foundZZZ == false {
		for _, direction := range directions {
			result.partOne++
			if result.partOne > 1000000 {
				panic("Steps exceeded 1,000,000")
			}

			if direction == "L" {
				currentPosition = nodes[currentPosition].left
			} else {
				currentPosition = nodes[currentPosition].right
			}

			if currentPosition == "ZZZ" {
				foundZZZ = true
				break
			}
		}
	}

	return result
}
