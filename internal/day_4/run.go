package day_4

import (
	"github.com/voziv/advent-of-code-2023/internal/util"
)

func Run() {
	util.AssertResult("Part 1 example.txt", runPartOne("./internal/day_4/example.txt"), "13")
	util.AssertResult("Part 1 input.txt", runPartOne("./internal/day_4/input.txt"), "18653")

	util.AssertResult("Part 2 example.txt", runPartTwo("./internal/day_4/example.txt"), "30")
	util.AssertResult("Part 2 input.txt", runPartTwo("./internal/day_4/input.txt"), "5921508")
}
