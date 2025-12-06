package main

import (
	"advent-of-code/common/aoc"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	task := aoc.Init(2023, 3)
	input := task.AsStringSlice()

	sum1 := 0

	r := regexp.MustCompile(`\d+`)
	rn := regexp.MustCompile(`[^\.\d]`)
	for i, row := range input {
		matches := r.FindAllStringIndex(row, -1)
		for _, m := range matches {
			start, end := m[0], m[1]
			lastIndex := len(row) - 1

			zone := ""
			if i > 0 {
				zone += input[i-1][max(start-1, 0):min(end+1, lastIndex)]
			}

			if start > 0 {
				zone += string(row[start-1])
			}

			if end < lastIndex {
				zone += string(row[end])
			}

			if i < len(input)-1 {
				zone += input[i+1][max(start-1, 0):min(end+1, lastIndex)]
			}

			if rn.MatchString(zone) {
				id, _ := strconv.Atoi(row[start:end])
				sum1 += id
			}

		}
	}

	fmt.Println("Puzzle 1:", sum1)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
