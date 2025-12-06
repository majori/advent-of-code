package main

import (
	"advent-of-code/common/aoc"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	task := aoc.Init(2023, 2)
	input := task.AsStringSlice()
	sum1, sum2 := 0, 0
	for i, game := range input {
		str := strings.TrimPrefix(game, fmt.Sprintf("Game %d: ", i+1))
		max := make(map[string]int, 3)
		sets := strings.Split(str, "; ")
		for _, set := range sets {
			cubes := strings.Split(set, ", ")
			for _, cube := range cubes {
				cubeParts := strings.Split(cube, " ")
				n, _ := strconv.Atoi(cubeParts[0])
				color := cubeParts[1]

				if max[color] < n {
					max[color] = n
				}
			}
		}

		if max["red"] <= 12 && max["green"] <= 13 && max["blue"] <= 14 {
			sum1 += i + 1
		}

		sum2 += max["red"] * max["green"] * max["blue"]
	}

	fmt.Println("Puzzle 1:", sum1)
	fmt.Println("Puzzle 2:", sum2)
}
