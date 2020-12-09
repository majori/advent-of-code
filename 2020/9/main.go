package main

import (
	"advent-of-code/common/aoc"
	"fmt"
)

const SIZE = 25

func main() {
	input := aoc.GetInputRowsAsInt(2020, 9)

	buffer := [SIZE]int{}
	var corrupted int

	for i, code := range input {
		if buffer[(i%SIZE)] == 0 || checkIfValid(code, &buffer) {
			buffer[i%SIZE] = code
		} else {
			fmt.Println("Puzzle 1:", code)
			corrupted = code
			break
		}
	}

	for i := range input {
		var acc int
		for j, code := range input[i:] {
			acc += code
			if acc == corrupted {
				var min, max int
				for k, code := range input[i : i+j+1] {
					if k == 0 || code < min {
						min = code
					} else if code > max {
						max = code
					}
				}
				fmt.Println("Puzzle 2:", min+max)
				return
			} else if acc > corrupted {
				break
			}
		}
	}
}

func checkIfValid(new int, buffer *[SIZE]int) bool {
	for i, x := range *buffer {
		compare := buffer[i+1:]
		for _, y := range compare {
			if x+y == new {
				return true
			}
		}
	}
	return false
}
