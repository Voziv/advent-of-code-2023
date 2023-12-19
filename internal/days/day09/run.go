package day09

import (
	"fmt"
	"github.com/voziv/advent-of-code-2023/internal/util"
)

type result struct {
	partOne int
}

func Run() {
	util.AssertResult("example.txt", run("./internal/days/day09/example.txt"), result{
		partOne: 114,
	})

	util.AssertResult("input.txt", run("./internal/days/day09/input.txt"), result{
		partOne: 2174807968,
	})
}

func run(inputFileName string) result {
	result := result{
		partOne: 0,
	}

	lines := util.GetFileContents(inputFileName)

	var trees [][][]int

	for _, line := range lines {
		numbers := parseNumbersFromInput(line)
		var tree [][]int
		tree = append(tree, numbers)

		isAllZeroes := false
		for !isAllZeroes {
			newRow := make([]int, len(numbers)-1)
			for i := 0; i < len(numbers)-1; i++ {
				newRow[i] = numbers[i+1] - numbers[i]
			}
			tree = append(tree, newRow)

			isAllZeroes := true
			for i := 0; i < len(newRow); i++ {
				if newRow[i] != 0 {
					isAllZeroes = false
					break
				}
			}

			if isAllZeroes {
				break
			}

			numbers = newRow
		}

		trees = append(trees, tree)

	}

	for _, tree := range trees {
		//printTree(tree)
		extrapolateTree(tree)
		//printTree(tree)
		result.partOne += tree[0][len(tree[0])-1]
	}

	return result
}

func extrapolateTree(tree [][]int) {
	for i := len(tree) - 1; i >= 0; i-- {
		tree[i] = append(tree[i], 0)
	}

	for row := len(tree) - 1; row >= 1; row-- {
		cell := len(tree[row]) - 1
		parentRow := row - 1
		leftParent := cell
		rightParent := len(tree[row])
		tree[parentRow][rightParent] = tree[parentRow][leftParent] + tree[row][cell]
	}
}

func printTree(tree [][]int) {
	for i := 0; i < len(tree); i++ {
		for j := 0; j < i; j++ {
			fmt.Printf("  ")
		}
		for j := 0; j < len(tree[i]); j++ {
			fmt.Printf(" %2d ", tree[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}
