package main

import (
	"advent-of-code/common/aoc"
	"fmt"
	"strconv"
)

const size = 99

func main() {
	task := aoc.Init(2022, 8)
	input := task.AsStringSlice()

	grid := [size][size]int{}
	for y := range input {
		for x := range input[y] {
			value, _ := strconv.Atoi(string(input[y][x]))
			grid[y][x] = value
		}
	}

	var count1, maxScore2 int
	for y := range grid {
		for x := range grid[y] {
			if isCellVisible(x, y, grid) {
				count1++
			}
			if score := treeScore(x, y, grid); score > maxScore2 {
				maxScore2 = score
			}
		}
	}

	fmt.Println("Puzzle 1: ", count1)
	fmt.Println("Puzzle 2: ", maxScore2)
}

func isCellVisible(x, y int, grid [size][size]int) bool {
	if x == 0 || y == 0 || x == size-1 || y == size-1 {
		return true
	}

	visibleFromDir := 4
	cell := grid[y][x]

	// left to right
	for x0 := 0; x0 < x; x0++ {
		if grid[y][x0] >= cell {
			visibleFromDir--
			break
		}
	}

	// right to left
	for x0 := size - 1; x0 > x; x0-- {
		if grid[y][x0] >= cell {
			visibleFromDir--
			break
		}
	}

	// top to bottom
	for y0 := 0; y0 < y; y0++ {
		if grid[y0][x] >= cell {
			visibleFromDir--
			break
		}
	}

	// bottom to top
	for y0 := size - 1; y0 > y; y0-- {
		if grid[y0][x] >= cell {
			visibleFromDir--
			break
		}
	}

	return visibleFromDir > 0
}

func treeScore(x, y int, grid [size][size]int) int {
	distances := [4]int{}
	cell := grid[y][x]

	// right
	for x0 := x + 1; x0 < size; x0++ {
		distances[0]++
		if grid[y][x0] >= cell {
			break
		}
	}

	// left
	for x0 := x - 1; x0 >= 0; x0-- {
		distances[1]++
		if grid[y][x0] >= cell {
			break
		}
	}

	// bottom
	for y0 := y + 1; y0 < size; y0++ {
		distances[2]++
		if grid[y0][x] >= cell {
			break
		}
	}

	// top
	for y0 := y - 1; y0 >= 0; y0-- {
		distances[3]++
		if grid[y0][x] >= cell {
			break
		}
	}

	return distances[0] * distances[1] * distances[2] * distances[3]
}
