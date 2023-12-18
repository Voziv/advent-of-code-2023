package main

import (
	"fmt"
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
}

func main() {
	if os.Args[1] == "all" {
		fmt.Println("Running All Days")
		for dayNumber, runner := range DayFunctions {
			fmt.Printf("\n########################################\n")
			fmt.Printf("Running Day %d\n", dayNumber)
			fmt.Printf("########################################\n")
			runner()
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
