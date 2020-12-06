package main

import (
	"advent-of-code/common/aoc"
	"fmt"
)

type Slope struct {
	X int
	Y int
}

func main() {
	in := aoc.GetInputRows(2020, 3)

	getEncountersForSlope := func(slope Slope) (encounters int) {
		var x int
		for y := 0; y < len(in); y += slope.Y {
			if rune(in[y][x%len(in[y])]) == '#' {
				encounters++
			}
			x += slope.X
		}

		return
	}

	slope := Slope{3, 1}
	encounters := getEncountersForSlope(slope)
	fmt.Println("Puzzle 1: ", encounters)

	slopes := []Slope{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	mul := 1

	for _, slope := range slopes {
		mul *= getEncountersForSlope(slope)
	}

	fmt.Println("Puzzle 2: ", mul)
}
