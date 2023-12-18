package main

import (
	"github.com/voziv/advent-of-code-2023/internal/day_1"
	"github.com/voziv/advent-of-code-2023/internal/day_2"
	"github.com/voziv/advent-of-code-2023/internal/day_3"
	"github.com/voziv/advent-of-code-2023/internal/day_4"
	"github.com/voziv/advent-of-code-2023/internal/day_5"
	"github.com/voziv/advent-of-code-2023/internal/day_6"
	"os"
)

var DAYS = map[string]func(){
	"1": day_1.Run,
	"2": day_2.Run,
	"3": day_3.Run,
	"4": day_4.Run,
	"5": day_5.Run,
	"6": day_6.Run,
}

func main() {
	runner, ok := DAYS[os.Args[1]]
	if !ok {
		panic("We don't have that day coded!")
	}

	runner()
}
