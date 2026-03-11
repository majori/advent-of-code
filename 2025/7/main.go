package main

import (
	"advent-of-code/common/aoc"
	"fmt"
)

func main() {
	task := aoc.Init(2025, 7)
	input := task.AsRuneGrid()

	result1, result2 := 0, 0

	for input.Iterate() {
		if input.Get() == 'S' {
			x, y := input.Cursor()
			result2 = traverse(input, x, y)
			break
		}
	}

	fmt.Println("Puzzle 1:", result1)
	fmt.Println("Puzzle 2:", result2)
}

func traverse(input aoc.Grid, x, y int) int {
	ok := input.SetCursor(x, y+1)
	if !ok {
		return 1
	}
	for input.Get() == '.' {
		return traverse(input, x, y)
	}

	x, y = input.Cursor()
	return traverse(input, x-1, y) + traverse(input, x+1, y)
}
