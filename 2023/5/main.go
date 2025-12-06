package main

import (
	"advent-of-code/common/aoc"
	"fmt"
	"strings"
)

func main() {
	task := aoc.Init(2023, 5)
	input := task.AsString()
	fmt.Println(strings.Split(input, "\n\n"))
}
