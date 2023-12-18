package main

import (
	"github.com/voziv/advent-of-code-2023/internal/day0"
	"github.com/voziv/advent-of-code-2023/internal/day1"
	"github.com/voziv/advent-of-code-2023/internal/day2"
	"github.com/voziv/advent-of-code-2023/internal/day3"
	"github.com/voziv/advent-of-code-2023/internal/day4"
	"github.com/voziv/advent-of-code-2023/internal/day5"
	"github.com/voziv/advent-of-code-2023/internal/day6"
	"github.com/voziv/advent-of-code-2023/internal/day7"
	"github.com/voziv/advent-of-code-2023/internal/day8"
	"os"
)

var DayFunctions = map[string]func(){
	"0": day0.Run,
	"1": day1.Run,
	"2": day2.Run,
	"3": day3.Run,
	"4": day4.Run,
	"5": day5.Run,
	"6": day6.Run,
	"7": day7.Run,
	"8": day8.Run,
}

func main() {
	runner, ok := DayFunctions[os.Args[1]]
	if !ok {
		panic("We don't have that day coded!")
	}

	runner()
}
