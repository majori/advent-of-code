package main

import (
	"advent-of-code/common/aoc"
	"fmt"
)

func main() {
	task := aoc.Init(2025, 10)
	input := task.AsStringSlice()

	fmt.Println(input)
	result1, result2 := 0, 0

	fmt.Println("Puzzle 1:", result1)
	fmt.Println("Puzzle 2:", result2)
}
