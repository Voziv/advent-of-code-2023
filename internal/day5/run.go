package day5

import (
	"github.com/voziv/advent-of-code-2023/internal/util"
)

func Run() {
	util.AssertResult("Part 1 example.txt", runPartOne("./internal/day5/example.txt"), "35")
	util.AssertResult("Part 1 input.txt", runPartOne("./internal/day5/input.txt"), "251346198")

	util.AssertResult("Part 2 example.txt", runPartTwo("./internal/day5/example.txt"), "46")
	util.AssertResult("Part 2 input.txt", runPartTwo("./internal/day5/input.txt"), "72263011")
}