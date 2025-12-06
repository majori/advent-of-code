package main

import (
	"advent-of-code/common/aoc"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	task := aoc.Init(2024, 1)
	input := task.AsStringSlice()

	list1 := make([]int, len(input))
	list2 := make([]int, len(input))
	list2Counts := make(map[int]int)

	for i, line := range input {
		parts := strings.Split(line, "   ")
		list1[i], _ = strconv.Atoi(parts[0])
		list2[i], _ = strconv.Atoi(parts[1])
		list2Counts[list2[i]]++
	}

	slices.Sort(list1)
	slices.Sort(list2)

	sum1 := 0
	for i := range input {
		tmp := list1[i] - list2[i]
		if tmp < 0 {
			tmp = -tmp
		}
		sum1 += tmp
	}

	fmt.Println("Puzzle 1: ", sum1)

	sum2 := 0
	for _, v := range list1 {
		sum2 += v * list2Counts[v]
	}

	fmt.Println("Puzzle 2: ", sum2)
}
