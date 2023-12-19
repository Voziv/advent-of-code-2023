package day8

import (
	"github.com/voziv/advent-of-code-2023/internal/math"
	"github.com/voziv/advent-of-code-2023/internal/util"
	"strings"
	"sync"
)

type result struct {
	steps int
}

func Run() {
	util.AssertResult("P1: example.txt", run("./internal/days/day8/example.txt"), result{steps: 2})
	util.AssertResult("P1: example2.txt", run("./internal/days/day8/example2.txt"), result{steps: 6})
	util.AssertResult("P1: input.txt", run("./internal/days/day8/input.txt"), result{steps: 19637})

	util.AssertResult("P2: example3.txt", solveLikeAGhost("./internal/days/day8/example3.txt"), result{steps: 6})
	util.AssertResult("P2: input.txt", solveLikeAGhost("./internal/days/day8/input.txt"), result{steps: 8811050362409})
}

func run(inputFileName string) result {
	result := result{steps: 0}

	lines := util.GetFileContents(inputFileName)

	directions := strings.Split(lines[0], "")
	nodeLines := lines[2:]

	var nodes = map[string]struct {
		left  string
		right string
	}{}

	for _, line := range nodeLines {
		// AAA = (BBB, CCC)
		id := line[:3]
		left := line[7:10]
		right := line[12:15]

		nodes[id] = struct {
			left  string
			right string
		}{left: left, right: right}
	}

	foundZZZ := false
	currentPosition := "AAA"

	for foundZZZ == false {
		for _, direction := range directions {
			result.steps++
			if result.steps > 1000000 {
				panic("Steps exceeded 1,000,000")
			}

			if direction == "L" {
				currentPosition = nodes[currentPosition].left
			} else {
				currentPosition = nodes[currentPosition].right
			}

			if currentPosition == "ZZZ" {
				foundZZZ = true
				break
			}
		}
	}

	return result
}

// Using Lowest Common Denominator to solve. I only came to LCM after looking on reddit. It really bothers me that LCM
// only works for the input and doesn't work for the problem's description.
// I don't know how I'd recognize that I'd need LCM based on the problem's description.
func solveLikeAGhost(inputFileName string) result {
	result := result{steps: 0}
	lines := util.GetFileContents(inputFileName)

	directions := strings.Split(lines[0], "")
	nodeLines := lines[2:]

	var nodes = map[string]struct {
		left  string
		right string
	}{}

	var travellingNodes []string

	for _, line := range nodeLines {
		// AAA = (BBB, CCC)
		id := line[:3]
		left := line[7:10]
		right := line[12:15]

		if id[2:] == "A" {
			travellingNodes = append(travellingNodes, id)
		}

		nodes[id] = struct {
			left  string
			right string
		}{left: left, right: right}
	}

	stepsToFirstZ := make([]int, len(travellingNodes))

	var wg sync.WaitGroup

	for i := 0; i < len(travellingNodes); i++ {
		wg.Add(1)

		i := i
		go func() {
			defer wg.Done()

			foundZ := false

			for foundZ == false {
				for _, direction := range directions {
					stepsToFirstZ[i]++

					if direction == "L" {
						travellingNodes[i] = nodes[travellingNodes[i]].left
					} else {
						travellingNodes[i] = nodes[travellingNodes[i]].right
					}

					if travellingNodes[i][2:] == "Z" {
						foundZ = true
						break
					}
				}
			}
		}()

	}
	wg.Wait()

	numSteps := math.Lcm(stepsToFirstZ[0], stepsToFirstZ[1:])
	result.steps = numSteps
	return result

}
