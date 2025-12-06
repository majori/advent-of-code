package main

import (
	"advent-of-code/common/aoc"
	"fmt"
)

func main() {
	task := aoc.Init(2024, 4)
	input := task.AsRuneGrid()

	foo := func(y, x int, s string) bool {
		if rune(s[0]) != input[y][x] {
			return false
		} else if len(s) == 1 {
			return true
		}

		if y > 0 {

		} else if y < len(input)-1 {

		}

		if x > 0 {

		} else if x < len(input[y])-1 {

		}

		return false
	}

	count := 0
	for y := range input {
		for x := range input[y] {
			if foo(y, x, "XMAS") {
				count++
			}
		}
	}

	fmt.Printf("%v\n", input)

	// fmt.Println("Puzzle 1:", input)
	// fmt.Println("Puzzle 2:", "")
}
