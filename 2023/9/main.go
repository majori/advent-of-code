package main

import (
	"advent-of-code/common/aoc"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	task := aoc.Init(2023, 9)
	input := task.AsStringSlice()

	sum1 := 0
	sum2 := 0

	for _, row := range input {
		ds := strings.Split(row, " ")
		d := make([]int, len(ds))
		for i, s := range ds {
			d[i], _ = strconv.Atoi(s)
		}

		d2, d1 := Extrapolate(d)
		sum1 += d1
		sum2 += d2
	}

	fmt.Println("Puzzle 1:", sum1)
	fmt.Println("Puzzle 1:", sum2)
}

func Extrapolate(d []int) (int, int) {
	diffs := make([]int, len(d)-1)
	allZeros := true

	for i := 1; i < len(d); i++ {
		diffs[i-1] = d[i] - d[i-1]
		if diffs[i-1] != 0 {
			allZeros = false
		}
	}

	if allZeros {
		return d[0], d[0]
	}

	f, l := Extrapolate(diffs)

	return d[0] - f, d[len(d)-1] + l
}
