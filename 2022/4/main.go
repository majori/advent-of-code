package main

import (
	"advent-of-code/common/aoc"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	task := aoc.Init(2022, 4)
	input := task.AsStringSlice()

	reg := regexp.MustCompile(`(\d+)-(\d+),(\d+)-(\d+)`)

	var p1, p2 int
	for _, r := range input {
		line := reg.FindStringSubmatch(r)
		x1, _ := strconv.Atoi(line[1])
		x2, _ := strconv.Atoi(line[2])
		y1, _ := strconv.Atoi(line[3])
		y2, _ := strconv.Atoi(line[4])

		if (x1 <= y1 && x2 >= y2) || (y1 <= x1 && y2 >= x2) {
			p1++
		}

		if !((x1 < y1 && x2 < y1) || (y1 < x1 && y2 < x1)) {
			p2++
		}

	}

	fmt.Println("Puzzle 1:", p1)
	fmt.Println("Puzzle 2:", p2)
}
