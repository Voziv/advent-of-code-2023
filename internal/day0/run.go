package day0

import (
	"github.com/voziv/advent-of-code-2023/internal/util"
)

type result struct {
	partOne int
	partTwo int
}

func Run() {
	util.AssertResult("example.txt", run("./internal/day0/example.txt"), result{
		partOne: 0,
		partTwo: 0,
	})

	util.AssertResult("input.txt", run("./internal/day0/input.txt"), result{
		partOne: 0,
		partTwo: 0,
	})
}

func run(inputFileName string) result {
	result := result{
		partOne: 0,
		partTwo: 0,
	}

	return result
}
