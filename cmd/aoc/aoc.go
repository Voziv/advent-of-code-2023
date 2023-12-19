package main

import (
	"fmt"
	"github.com/voziv/advent-of-code-2023/internal/days/day0"
	"github.com/voziv/advent-of-code-2023/internal/days/day1"
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
	"github.com/voziv/advent-of-code-2023/internal/days/day2"
	"github.com/voziv/advent-of-code-2023/internal/days/day20"
	"github.com/voziv/advent-of-code-2023/internal/days/day21"
	"github.com/voziv/advent-of-code-2023/internal/days/day22"
	"github.com/voziv/advent-of-code-2023/internal/days/day23"
	"github.com/voziv/advent-of-code-2023/internal/days/day24"
	"github.com/voziv/advent-of-code-2023/internal/days/day25"
	"github.com/voziv/advent-of-code-2023/internal/days/day3"
	"github.com/voziv/advent-of-code-2023/internal/days/day4"
	"github.com/voziv/advent-of-code-2023/internal/days/day5"
	"github.com/voziv/advent-of-code-2023/internal/days/day6"
	"github.com/voziv/advent-of-code-2023/internal/days/day7"
	"github.com/voziv/advent-of-code-2023/internal/days/day8"
	"github.com/voziv/advent-of-code-2023/internal/days/day9"
	"os"
	"strconv"
)

var DayFunctions = []func(){
	day0.Run,
	day1.Run,
	day2.Run,
	day3.Run,
	day4.Run,
	day5.Run,
	day6.Run,
	day7.Run,
	day8.Run,
	day9.Run,
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
