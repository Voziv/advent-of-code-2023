package day20

import (
	"github.com/voziv/advent-of-code-2023/internal/util"
)

type result struct {
	partOne int
}

func Run() {
	util.AssertResult("example.txt", run("./internal/days/day20/example.txt"), result{
		partOne: 0,
	})

	util.AssertResult("input.txt", run("./internal/days/day20/input.txt"), result{
		partOne: 0,
	})
}

func run(inputFileName string) result {
	result := result{
		partOne: 0,
	}

	panic("Not implemented")

	return result
}
