package day0

import (
	"github.com/voziv/advent-of-code-2023/internal/util"
)

type day0Answer struct {
	partOne int
	partTwo int
}

func Run() {
	util.AssertResult("example.txt", run("./internal/day0/example.txt"), day0Answer{
		partOne: 0,
		partTwo: 0,
	})

	util.AssertResult("input.txt", run("./internal/day0/input.txt"), day0Answer{
		partOne: 0,
		partTwo: 0,
	})
}

func run(inputFileName string) day0Answer {
	return day0Answer{
		partOne: 0,
		partTwo: 0,
	}
}
