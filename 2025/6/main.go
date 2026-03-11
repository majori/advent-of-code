package main

import (
	"advent-of-code/common/aoc"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	task := aoc.Init(2025, 6)
	input := task.AsStringSlice()

	rs := make([][]int, 0)
	fs := make([]func(a, b int) int, 0)

	for i, row := range input {
		cells := strings.Split(row, " ")
		n := 0
		for _, c := range cells {
			if c != "" {
				cells[n] = c
				n++
			}
		}
		cells = cells[:n]

		if i != len(input)-1 {
			r := make([]int, 0)
			for _, cell := range cells {
				c, _ := strconv.Atoi(cell)
				r = append(r, c)
			}
			rs = append(rs, r)
		} else {
			var f func(a, b int) int
			for _, cell := range cells {
				if cell == "+" {
					f = func(a, b int) int { return a + b }
				}
				if cell == "*" {
					f = func(a, b int) int { return a * b }
				}
				fs = append(fs, f)
			}
		}
	}

	result1 := 0
	result2 := 0

	for x := range rs[0] {
		f := fs[x]

		// Puzzle 1
		var columnSum int
		for y := 0; y < len(rs); y++ {
			if y == 0 {
				columnSum = rs[y][x]
				continue
			}
			columnSum = f(rs[y][x], columnSum)
		}
		result1 += columnSum

		// Puzzle 2

	}

	fmt.Println("Puzzle 1:", result1)
	fmt.Println("Puzzle 2:", result2)
}
