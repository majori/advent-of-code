package main

import (
	"advent-of-code/common/aoc"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	task := aoc.Init(2025, 9)
	input := task.AsStringSlice()
	coordinates := make([][2]int, len(input))
	result1, result2 := 0, 0

	for i, c := range input {
		s := strings.Split(c, ",")
		x, _ := strconv.Atoi(s[0])
		y, _ := strconv.Atoi(s[1])
		coordinates[i] = [2]int{x, y}
	}

	for i := range coordinates {
		for j := range coordinates {
			if i == j {
				continue
			}

			area := int(math.Abs(float64(coordinates[i][0]-coordinates[j][0]))+1) * int(math.Abs(float64(coordinates[i][1]-coordinates[j][1]))+1)
			if area > result1 {
				result1 = area
			}
		}
	}

	fmt.Println("Puzzle 1:", result1)
	fmt.Println("Puzzle 2:", result2)
}
