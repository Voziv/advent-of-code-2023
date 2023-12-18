package day_7

import (
	"github.com/voziv/advent-of-code-2023/internal/util"
)

func Run() {
	util.AssertResult("Part 1 example.txt", runPartOne("./internal/day_7/example.txt"), "6440")
	util.AssertResult("Part 1 input.txt", runPartOne("./internal/day_7/input.txt"), "0")

	util.AssertResult("Part 2 example.txt", runPartTwo("./internal/day_7/example.txt"), "0")
	util.AssertResult("Part 2 input.txt", runPartTwo("./internal/day_7/input.txt"), "0")
}
