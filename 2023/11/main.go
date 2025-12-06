package main

import (
	"advent-of-code/common/aoc"
	"fmt"
	"strings"
)

type Node struct {
	Left  string
	Right string
}

func main() {
	task := aoc.Init(2023, 11)
	input := task.AsStringSlice()
	fmt.Println(input)
	matrix := make([][]rune, 0, len(input))

	for _, line := range input {
		matrix = append(matrix, []rune(line))
		if !strings.Contains("#", line) {
			matrix = append(matrix, []rune(line))
		}
	}

	for i := 0; i < len(matrix[0]); i++ {
		fmt.Println(string(matrix[i]))
	}
}
