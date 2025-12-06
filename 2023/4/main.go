package main

import (
	"advent-of-code/common/aoc"
	"fmt"
	"math"
	"strings"
)

func main() {
	task := aoc.Init(2023, 4)
	input := task.AsStringSlice()

	matches := make([]int, len(input))
	for i, row := range input {
		row = strings.TrimPrefix(row, fmt.Sprintf("Card %3d: ", i+1))
		card := strings.Split(row, " | ")
		win := TrimEmpty(strings.Split(card[0], " "))
		scratches := TrimEmpty(strings.Split(card[1], " "))

		winMap := make(map[string]bool)
		for _, w := range win {
			winMap[w] = true
		}

		for _, r := range scratches {
			if _, ok := winMap[r]; ok {
				matches[i]++
			}
		}
	}

	scoreSum := 0
	copiesSum := 0
	for i := range matches {
		if matches[i] > 0 {
			scoreSum += int(math.Pow(2, float64(matches[i]-1)))
		}
		copiesSum++
		copiesSum += X(matches, i)
	}

	fmt.Println("Puzzle 1:", scoreSum)
	fmt.Println("Puzzle 2:", copiesSum)
}

var cache map[int]int = make(map[int]int, 193)

func X(matches []int, i int) int {
	if _, found := cache[i]; found {
		return cache[i]
	}

	copies := matches[i]
	for j := 1; j <= matches[i]; j++ {
		copies += X(matches, i+j)
	}

	cache[i] = copies
	return copies
}

func TrimEmpty(s []string) []string {
	for i, v := range s {
		if v == "" {
			s = append(s[:i], s[i+1:]...)
		}
	}

	return s
}
