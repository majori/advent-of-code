package main

import (
	"advent-of-code/common/aoc"
	"strconv"
	"strings"
)

func main() {
	task := aoc.Init(2021, 7)
	input := strings.Split(task.AsString(), ",")

	crabs := make([]int, len(input))

	for i, x := range input {
		x, _ := strconv.Atoi(x)
		crabs[i] = x
	}

	task.Submit1("asd")
}