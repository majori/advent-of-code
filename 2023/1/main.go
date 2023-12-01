package main

import (
	"advent-of-code/common/aoc"
	"fmt"
	"strings"
)

var digitMap map[string]int = map[string]int{
	"0":     0,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func main() {
	task := aoc.Init(2023, 1)
	input := task.AsStringSlice()
	sum1, sum2 := 0, 0

	for _, row := range input {
		first, second := -1, -1
		for _, c := range row {
			digit, found := digitMap[string(c)]
			if !found {
				continue
			}

			if first == -1 {
				first = digit
			}
			second = digit
		}

		if first >= 0 {
			sum1 += first*10 + second
		}
	}

	fmt.Println("Puzzle 1: ", sum1)

	for _, row := range input {
		first, second := -1, -1
		for i := range row {
			for k, digit := range digitMap {
				if !strings.HasPrefix(row[i:], k) {
					continue
				}
				if first == -1 {
					first = digit
				}
				second = digit
			}
		}

		if first >= 0 {
			sum2 += first*10 + second
		}
	}

	fmt.Println("Puzzle 2: ", sum2)
}
