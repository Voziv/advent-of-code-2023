package day10

import (
	"github.com/voziv/advent-of-code-2023/internal/util"
	"strings"
)

type result struct {
	partOne int
	partTwo int
}

func Run() {
	util.AssertResult("example.txt", run("./internal/days/day10/example.txt"), result{
		partOne: 4,
		partTwo: 1,
	})

	util.AssertResult("example2.txt", run("./internal/days/day10/example2.txt"), result{
		partOne: 8,
		partTwo: 1,
	})

	util.AssertResult("example3.txt", run("./internal/days/day10/example3.txt"), result{
		partOne: 23,
		partTwo: 4,
	})

	util.AssertResult("example4.txt", run("./internal/days/day10/example4.txt"), result{
		partOne: 22,
		partTwo: 4,
	})

	util.AssertResult("example5.txt", run("./internal/days/day10/example5.txt"), result{
		partOne: 70,
		partTwo: 8,
	})

	util.AssertResult("input.txt", run("./internal/days/day10/input.txt"), result{
		partOne: 6754,
		partTwo: 567,
	})
}

func run(inputFileName string) result {
	result := result{
		partOne: 0,
		partTwo: 0,
	}

	lines := util.GetFileContents(inputFileName)

	var data []string
	for _, line := range lines {
		data = append(data, strings.Split(line, "")...)
	}

	for i := 0; i < len(data); i++ {
		data[i] = asciiToUnicodePipeMap[data[i]]
	}

	maze := NewMaze(data, len(lines[0]), len(lines))

	result.partOne = maze.DistanceToFurthestPointFromStart()
	result.partTwo = maze.FindNestsUsingScanLine(true)
	maze.Print()

	return result
}
