package main

import (
	"advent-of-code/common/aoc"
	"fmt"
	"strings"
)

func main() {
	task := aoc.Init(2022, 6)
	input := task.AsString()

	fmt.Println("Puzzle 1:", GetFirstUniqueWindowIndex(input, 4))
	fmt.Println("Puzzle 2:", GetFirstUniqueWindowIndex(input, 14))
}

func GetFirstUniqueWindowIndex(buffer string, windowSize int) int {
	for i := 0; i < len(buffer)-windowSize; i++ {
		window := buffer[i : i+windowSize]

		uniq := true
		for j := 0; j < windowSize; j++ {
			if strings.Count(window, string(window[j])) > 1 {
				uniq = false
				break
			}
		}

		if uniq {
			return i + windowSize
		}
	}

	return 0
}
