package main

import (
	"advent-of-code/common/aoc"
	"fmt"
	"sort"
	"strconv"
)

func main() {
	task := aoc.Init(2022, 1)
	input := task.AsStringSlice()

	elves := make([]int, 1)
	var i, max int
	for _, l := range input {
		if l == "" {
			if elves[i] > max {
				max = elves[i]
			}
			elves = append(elves, 0)
			i++
			continue
		}

		c, _ := strconv.Atoi(l)
		elves[i] += c
	}

	fmt.Println("Puzzle 1: ", max)

	sort.Ints(elves)
	n := len(elves)
	fmt.Println("Puzzle 2: ", elves[n-1]+elves[n-2]+elves[n-3])
}
