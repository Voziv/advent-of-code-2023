package day_2

import (
	"github.com/voziv/advent-of-code-2023/internal/util"
)

func Run() {
	util.AssertResult("Part 1 example.txt", runPartOne("./internal/day_2/example.txt"), "8")
	util.AssertResult("Part 1 input.txt", runPartOne("./internal/day_2/input.txt"), "2149")

	util.AssertResult("Part 2 example.txt", runPartTwo("./internal/day_2/example.txt"), "2286")
	util.AssertResult("Part 2 input.txt", runPartTwo("./internal/day_2/input.txt"), "2286")

}
