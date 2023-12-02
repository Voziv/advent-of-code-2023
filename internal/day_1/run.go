package day_1

import (
	"github.com/voziv/advent-of-code-2023/internal/util"
)

func Run() {
	util.AssertResult("Part 1 example.txt", runPartOne("./internal/day_1/example.txt"), "142")
	util.AssertResult("Part 1 input.txt", runPartOne("./internal/day_1/input.txt"), "54953")

	util.AssertResult("Part 2 example2.txt", runPartTwo("./internal/day_1/example2.txt"), "281")

	// 51175 is too low
	// 53868 is just right
	// 53885 is too high
	util.AssertResult("Part 2 input.txt", runPartTwo("./internal/day_1/input.txt"), "53868")
}
