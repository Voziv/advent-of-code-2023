package day_3

import (
	"github.com/voziv/advent-of-code-2023/internal/util"
)

func Run() {
	util.AssertResult("Part 1 example.txt", runPartOne("./internal/day_3/example.txt"), "4361")
	util.AssertResult("Part 1 input.txt", runPartOne("./internal/day_3/input.txt"), "540025")

	util.AssertResult("Part 2 example.txt", runPartTwo("./internal/day_2/example.txt"), "467835")
	util.AssertResult("Part 2 input.txt", runPartTwo("./internal/day_2/input.txt"), "")

}
