package main

import (
	"fmt"
	"github.com/voziv/advent-of-code-2023/internal/days/day00"
	"github.com/voziv/advent-of-code-2023/internal/days/day01"
	"github.com/voziv/advent-of-code-2023/internal/days/day02"
	"github.com/voziv/advent-of-code-2023/internal/days/day03"
	"github.com/voziv/advent-of-code-2023/internal/days/day04"
	"github.com/voziv/advent-of-code-2023/internal/days/day05"
	"github.com/voziv/advent-of-code-2023/internal/days/day06"
	"github.com/voziv/advent-of-code-2023/internal/days/day07"
	"github.com/voziv/advent-of-code-2023/internal/days/day08"
	"github.com/voziv/advent-of-code-2023/internal/days/day09"
	"github.com/voziv/advent-of-code-2023/internal/days/day10"
	"github.com/voziv/advent-of-code-2023/internal/days/day11"
	"github.com/voziv/advent-of-code-2023/internal/days/day12"
	"github.com/voziv/advent-of-code-2023/internal/days/day13"
	"github.com/voziv/advent-of-code-2023/internal/days/day14"
	"github.com/voziv/advent-of-code-2023/internal/days/day15"
	"github.com/voziv/advent-of-code-2023/internal/days/day16"
	"github.com/voziv/advent-of-code-2023/internal/days/day17"
	"github.com/voziv/advent-of-code-2023/internal/days/day18"
	"github.com/voziv/advent-of-code-2023/internal/days/day19"
	"github.com/voziv/advent-of-code-2023/internal/days/day20"
	"github.com/voziv/advent-of-code-2023/internal/days/day21"
	"github.com/voziv/advent-of-code-2023/internal/days/day22"
	"github.com/voziv/advent-of-code-2023/internal/days/day23"
	"github.com/voziv/advent-of-code-2023/internal/days/day24"
	"github.com/voziv/advent-of-code-2023/internal/days/day25"
	"os"
	"strconv"
)

var DayFunctions = []func(){
	day00.Run,
	day01.Run,
	day02.Run,
	day03.Run,
	day04.Run,
	day05.Run,
	day06.Run,
	day07.Run,
	day08.Run,
	day09.Run,
	day10.Run,
	day11.Run,
	day12.Run,
	day13.Run,
	day14.Run,
	day15.Run,
	day16.Run,
	day17.Run,
	day18.Run,
	day19.Run,
	day20.Run,
	day21.Run,
	day22.Run,
	day23.Run,
	day24.Run,
	day25.Run,
}

func main() {
	if os.Args[1] == "all" {
		fmt.Println("Running All Days")
		for i := 1; i < len(DayFunctions); i++ {
			fmt.Printf("\n########################################\n")
			fmt.Printf("Running Day %d\n", i)
			fmt.Printf("########################################\n")
			DayFunctions[i]()

		}
	} else {
		dayNumber, err := strconv.Atoi(os.Args[1])

		if err != nil {
			fmt.Println("You must provide a valid day number")
			os.Exit(1)
		}

		if dayNumber > len(DayFunctions) {
			fmt.Println("We don't have that day coded!")
			os.Exit(1)
		}

		fmt.Printf("\n########################################\n")
		fmt.Printf("Running Day %s\n", os.Args[1])
		fmt.Printf("########################################\n")
		DayFunctions[dayNumber]()
	}

}
