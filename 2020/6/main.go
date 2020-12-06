package main

import (
	"advent-of-code/common/aoc"
	"fmt"
)

func main() {
	input := aoc.GetInputRows(2020, 6)
	input = append(input, "")

	var sum1, sum2, ppl int

	groupAnswers := make(map[rune]int)

	for _, row := range input {
		if row == "" {
			for _, value := range groupAnswers {
				if value == ppl {
					sum2++
				}
			}
			ppl = 0
			sum1 += len(groupAnswers)
			groupAnswers = make(map[rune]int)
			continue
		}

		for _, q := range row {
			groupAnswers[q]++
		}
		ppl++
	}
	fmt.Println("Puzzle 1: ", sum1)
	fmt.Println("Puzzle 2: ", sum2)
}
