package main

import (
	"advent-of-code/common/aoc"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := aoc.GetInputRows(2020, 13)
	initialWaitTtime, _ := strconv.Atoi(input[0])
	lines := strings.Split(input[1], ",")

	min := -1
	chosenLine := 0
	for _, line := range lines {
		if line == "x" {
			continue
		}
		interval, _ := strconv.Atoi(line)

		waitTime := 0
		if initialWaitTtime%interval != 0 {
			waitTime = interval - initialWaitTtime%interval
		}

		if min == -1 || waitTime < min {
			min = waitTime
			chosenLine = interval
		}
	}

	fmt.Println("Puzzle 1:", min*chosenLine)

	var timestamp int64

}
