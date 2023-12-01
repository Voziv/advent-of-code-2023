package day_1

import (
	"github.com/voziv/advent-of-code-2023/internal/util"
)

func Run() {
	util.AssertResult("Part 1 example.txt", runPartOne("./inputs/1/example.txt"), "0")
	util.AssertResult("Part 1 intput.txt", runPartOne("./inputs/1/input.txt"), "0")
	util.AssertResult("Part 2 example.txt", runPartTwo("./inputs/1/example.txt"), "0")
	util.AssertResult("Part 2 input.txt", runPartTwo("./inputs/1/input.txt"), "0")
}
