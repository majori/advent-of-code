package main

import (
	"advent-of-code/common/aoc"
	"fmt"
	"regexp"
	"strconv"
)

type Instruction struct {
	Command   string
	Parameter int
}

func main() {
	input := aoc.GetInput(2020, 8)
	reg := regexp.MustCompile(`(\w+) ((?:\+|-)\d+)`)

	rows := reg.FindAllStringSubmatch(input, -1)
	instructions := make([]Instruction, 0, len(rows))
	for _, row := range rows {
		command := row[1]
		parameter, _ := strconv.Atoi(row[2])

		instruction := Instruction{command, parameter}
		instructions = append(instructions, instruction)
	}

	acc, _ := run(instructions)
	fmt.Println("Puzzle 1:", acc)

	for i := range instructions {
		if instructions[i].Command == "acc" {
			continue
		}

		modified := make([]Instruction, len(instructions))
		copy(modified, instructions)

		switch instructions[i].Command {
		case "nop":
			modified[i].Command = "jmp"
		case "jmp":
			modified[i].Command = "nop"
		}

		acc, isLoop := run(modified)
		if !isLoop {
			fmt.Println(fmt.Sprintf("Puzzle 2: %d (instruction index: %d)", acc, i))
			break
		}
	}
}

func run(instructions []Instruction) (acc int, isLoop bool) {
	history := make(map[int]bool)
	i := 0

	for {
		if i >= len(instructions) {
			return
		}

		instruction := instructions[i]

		switch instruction.Command {
		case "nop":
			break
		case "acc":
			acc += instruction.Parameter
		case "jmp":
			i += instruction.Parameter - 1
		}

		i++

		if _, ok := history[i]; ok {
			isLoop = true
			return
		} else {
			history[i] = true
		}
	}
}
