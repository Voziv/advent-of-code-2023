package day8

import (
	"github.com/voziv/advent-of-code-2023/internal/math"
	"github.com/voziv/advent-of-code-2023/internal/util"
	"strings"
)

type result struct {
	steps int
}

func Run() {
	util.AssertResult("P1: example.txt", run("./internal/day8/example.txt"), result{steps: 2})
	util.AssertResult("P1: example2.txt", run("./internal/day8/example2.txt"), result{steps: 6})
	util.AssertResult("P1: input.txt", run("./internal/day8/input.txt"), result{steps: 19637})

	util.AssertResult("P2: example3.txt", solveLikeAGhost("./internal/day8/example3.txt"), result{steps: 6})
	util.AssertResult("P2: input.txt", solveLikeAGhost("./internal/day8/input.txt"), result{steps: 8811050362409})
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

	//fmt.Println(travellingNodes)

	countZ := 0
	greatestZ := 0
	zStepsByIndex := make([]int, len(travellingNodes))
	for i := 0; i < len(zStepsByIndex); i++ {
		zStepsByIndex[i] = 0
	}
	for countZ != len(travellingNodes) {
		for _, direction := range directions {
			result.steps++
			//if result.steps > 1000000000 {
			//	panic("Steps exceeded 1,000,000,000")
			//}

			countZ = 0
			for i, node := range travellingNodes {
				if direction == "L" {
					travellingNodes[i] = nodes[node].left
				} else {
					travellingNodes[i] = nodes[node].right
				}

				if travellingNodes[i][2:] == "Z" {
					if zStepsByIndex[i] == 0 {
						zStepsByIndex[i] = result.steps
						//fmt.Printf("%v after %d steps\n", travellingNodes[i], result.steps)
						hasZeros := false
						totalSteps := 1
						for i := 0; i < len(zStepsByIndex); i++ {
							if zStepsByIndex[i] == 0 {
								hasZeros = true
								break
							}
							totalSteps *= zStepsByIndex[i]
						}

						if !hasZeros {
							//fmt.Println("All nodes have found a minimum number of steps to get Z")
							//fmt.Println(totalSteps)

							numSteps := math.Lcm(zStepsByIndex[0], zStepsByIndex[1:])

							//fmt.Println("LCM")
							//fmt.Println(numSteps)

							result.steps = numSteps
							return result
						}
					}
					countZ++
				}
			}

			if countZ > greatestZ {
				greatestZ = countZ
				//fmt.Printf("%v - %d nodes ending in Z after %d steps\n", travellingNodes, countZ, result.steps)
			}

			if countZ == len(travellingNodes) {
				break
			}
		}
	}

	return result
}
