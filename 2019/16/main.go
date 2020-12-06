package main

import (
	"advent-of-code/common/aoc"
	"fmt"
	"strconv"
)

func main() {
	in := aoc.GetInput(2019, 16)
	original := make([]int, len(in))
	for i, c := range in {
		d, _ := strconv.Atoi(string(c))
		original[i] = d
	}
	pattern := []int{0, 1, 0, -1}

	signal := original
	for i := 0; i < 100; i++ {
		signal = runPhase(signal, pattern)
	}

	var answer1 string
	for _, d := range signal[:8] {
		answer1 += strconv.Itoa(d)
	}
	fmt.Printf("Answer 1: %s\n", answer1) // 30550349

	// signal = make([]int, 0, 10000*len(original))
	// for i := 0; i < 10000; i++ {
	// 	signal = append(signal, original...)
	// }

	// for i := 0; i < 100; i++ {
	// 	signal = runPhase(signal, pattern)
	// }

	// fmt.Println(signal)
}

func runPhase(signal []int, pattern []int) []int {
	out := make([]int, len(signal))

	for round := range signal {
		var sum, patternIndex, patternOffsetCounter int

		if round > 0 {
			patternIndex = 1
		}

		for _, value := range signal[round:] {
			if patternOffsetCounter >= round {
				patternIndex = (patternIndex + 1) % len(pattern)
				patternOffsetCounter = 0
			} else {
				patternOffsetCounter++
			}

			if pattern[patternIndex] == 0 {
				continue
			}

			sum += value * pattern[patternIndex]
		}

		if sum > 0 {
			out[round] = sum % 10
		} else {
			out[round] = -sum % 10
		}
	}

	return out
}
