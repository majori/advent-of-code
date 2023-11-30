package main

import (
	"advent-of-code/common/aoc"
	"fmt"
)

var strategy1 = map[string]int{
	"A X": 4, // 1 + 3
	"A Y": 8, // 2 + 6
	"A Z": 3, // 3 + 0
	"B X": 1, // 1 + 0
	"B Y": 5, // 2 + 3
	"B Z": 9, // 3 + 6
	"C X": 7, // 1 + 6
	"C Y": 2, // 2 + 0
	"C Z": 6, // 3 + 3
}

var strategy2 = map[string]int{
	"A X": 3, // 3 + 0
	"A Y": 4, // 1 + 3
	"A Z": 8, // 2 + 6
	"B X": 1, // 1 + 0
	"B Y": 5, // 2 + 3
	"B Z": 9, // 3 + 6
	"C X": 2, // 2 + 0
	"C Y": 6, // 3 + 3
	"C Z": 7, // 1 + 6
}

func main() {
	task := aoc.Init(2022, 2)
	input := task.AsStringSlice()

	var p1, p2 int
	for _, r := range input {
		p1 += strategy1[r]
		p2 += strategy2[r]
	}

	fmt.Println("Puzzle 1:", p1)
	fmt.Println("Puzzle 2:", p2)
}
