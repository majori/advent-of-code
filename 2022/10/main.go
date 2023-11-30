package main

import (
	"advent-of-code/common/aoc"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	task := aoc.Init(2022, 10)
	input := task.AsStringSlice()

	reg := regexp.MustCompile(`(\w+)(?: (-?\d+))?`)

	row := 0
	register, cycles := 1, 0
	signalStrengths := make([]int, 0)
	display := [6][40]rune{}

	tick := func() {
		cycles++
		if cycles >= 240 {
			return
		}
		if cycles == 20 || (20+cycles)%40 == 0 {
			signalStrengths = append(signalStrengths, cycles*register)
		}
		if cycles%40 == 0 {
			row++
		}
		cursor := cycles % 40
		if cursor >= register-1 && cursor <= register+1 {
			display[row][cursor] = '#'
		} else {
			display[row][cursor] = '.'
		}
	}

	for _, instruction := range input {
		parts := reg.FindStringSubmatch(instruction)
		switch parts[1] {
		case "addx":
			value, _ := strconv.Atoi(parts[2])
			tick()
			register += value
			tick()
		case "noop":
			tick()
		}
	}

	var p1 int
	for _, strength := range signalStrengths {
		p1 += strength
	}

	fmt.Println("Puzzle 1:", p1)

	fmt.Println("Puzzle 2:")
	for _, line := range display {
		for _, pixel := range line {
			fmt.Print(string(pixel))
		}
		fmt.Println()
	}
}
