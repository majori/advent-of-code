package main

import (
	"advent-of-code/common/aoc"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	task := aoc.Init(2022, 5)
	input := task.AsStringSlice()

	stacks := make([][]rune, 9)

	for _, row := range input[:8] {
		for i := 0; i < len(stacks); i++ {
			char := rune(row[i*4+1])
			if char != ' ' {
				stacks[i] = append([]rune{char}, stacks[i]...)
			}
		}
	}

	stacks1 := make([][]rune, 9)
	stacks2 := make([][]rune, 9)
	copy(stacks1, stacks)
	copy(stacks2, stacks)

	reg := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	for _, row := range input[10:] {
		line := reg.FindStringSubmatch(row)
		amount, _ := strconv.Atoi(line[1])
		from, _ := strconv.Atoi(line[2])
		to, _ := strconv.Atoi(line[3])

		for n := 0; n < amount; n++ {
			moveOne(stacks1, from-1, to-1)
		}

		moveMany(stacks2, amount, from-1, to-1)
	}

	fmt.Println("Puzzle 1:", printStack(stacks1))
	fmt.Println("Puzzle 2:", printStack(stacks2))
}

func moveOne(stacks [][]rune, from, to int) {
	stacks[to] = append(stacks[to], stacks[from][len(stacks[from])-1])
	stacks[from] = stacks[from][:len(stacks[from])-1]
}

func moveMany(stacks [][]rune, amount, from, to int) {
	stacks[to] = append(stacks[to], stacks[from][len(stacks[from])-amount:]...)
	stacks[from] = stacks[from][:len(stacks[from])-amount]
}

func printStack(stacks [][]rune) (output string) {
	for _, stack := range stacks {
		output += string(stack[len(stack)-1])
	}
	return
}
