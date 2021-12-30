package main

import (
	"advent-of-code/common/aoc"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	task := aoc.Init(2021, 2)
	input := task.AsStringSlice()

	reg := regexp.MustCompile(`(forward|down|up) (\d)`)

	var pos, depth int

	for _, row := range input {
		command := reg.FindStringSubmatch(row)
		dir := command[1]
		steps, _ := strconv.Atoi(command[2])

		switch dir {
		case "forward":
			pos += steps
		case "down":
			depth += steps
		case "up":
			depth -= steps
		}
	}

	fmt.Println("Puzzle 1: ", pos*depth)

	pos, depth = 0, 0
	aim := 0

	for _, row := range input {
		command := reg.FindStringSubmatch(row)
		dir := command[1]
		steps, _ := strconv.Atoi(command[2])

		switch dir {
		case "forward":
			pos += steps
			depth += aim*steps
		case "down":
			aim += steps
		case "up":
			aim -= steps
		}
	}

	fmt.Println("Puzzle 2: ", pos*depth)
}
