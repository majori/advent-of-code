package main

import (
	"advent-of-code/common/aoc"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	task := aoc.Init(2021, 5)
	input := task.AsStringSlice()

	reg := regexp.MustCompile(`(\d+),(\d+) -> (\d+),(\d+)`)

	chart := make(map[string]int)

	for _, row := range input {
		line := reg.FindStringSubmatch(row)
		x1, _ := strconv.Atoi(line[1])
		y1, _ := strconv.Atoi(line[2])
		x2, _ := strconv.Atoi(line[3])
		y2, _ := strconv.Atoi(line[4])

		xMul, yMul := 1, 1
		if x1 > x2 {
			xMul = -1
		}
		if y1 > y2 {
			yMul = -1
		}

		if y1 == y2 {
			y := y1
			for x := x1; x != x2; x+=1*xMul {
				chart[fmt.Sprintf("%d,%d", x, y)]++
			}
			chart[fmt.Sprintf("%d,%d", x2, y)]++
			} else if x1 == x2 {
				x := x1
				for y := y1; y != y2; y+=1*yMul {
					chart[fmt.Sprintf("%d,%d", x, y)]++
				}
				chart[fmt.Sprintf("%d,%d", x, y2)]++
		}
	}

	overlaps := 0
	for _, v := range chart {
		if v >= 2 {
			overlaps++
		}
	}

	fmt.Println("Puzzle 1:", overlaps)

	chart = make(map[string]int)

	for _, row := range input {
		line := reg.FindStringSubmatch(row)
		x1, _ := strconv.Atoi(line[1])
		y1, _ := strconv.Atoi(line[2])
		x2, _ := strconv.Atoi(line[3])
		y2, _ := strconv.Atoi(line[4])

		xMul, yMul := 1, 1

		if x1 > x2 {
			xMul = -1
		}
		if y1 > y2 {
			yMul = -1
		}

		x := x1
		y := y1
		for x != x2 && y != y2 {
			chart[fmt.Sprintf("%d,%d", x, y)]++
			x+=xMul
			y+=yMul
		}
		chart[fmt.Sprintf("%d,%d", x2, y2)]++
	}

	overlaps = 0
	for _, v := range chart {
		if v >= 2 {
			overlaps++
		}
	}

	fmt.Println("Puzzle 2:", overlaps)

}