package main

import (
	"advent-of-code/common/aoc"
	"fmt"
)

func main() {
	task := aoc.Init(2025, 4)
	input := task.AsRuneGrid()

	sum1, sum2 := 0, 0

	for input.Iterate() {
		if input.Get() == '@' {
			count := 0
			for _, n := range input.GetNeighbours() {
				if n == '@' {
					count++
				}
			}
			if count < 4 {
				sum1++
			}
		}
	}

	fmt.Println("Puzzle 1:", sum1)

	removed := true
	for removed {
		removed = false
		for input.Iterate() {
			if input.Get() == '@' {
				count := 0
				for _, n := range input.GetNeighbours() {
					if n == '@' || n == 'x' {
						count++
					}
				}
				if count < 4 {
					removed = true
					input.Set('x')
					sum2++
				}
			}
		}

		for input.Iterate() {
			if input.Get() == 'x' {
				input.Set('.')
			}
		}
	}

	fmt.Println("Puzzle 2:", sum2)
}
