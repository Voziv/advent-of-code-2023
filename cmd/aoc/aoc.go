package main

import (
	"github.com/voziv/advent-of-code-2023/internal/day_1"
	"os"
)

var DAYS = map[string]func(){
	"1": day_1.Run,
}

func main() {
	runner, ok := DAYS[os.Args[1]]
	if !ok {
		panic("We don't have that day coded!")
	}

	runner()
}
