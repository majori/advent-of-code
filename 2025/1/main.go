package main

import (
	"advent-of-code/common/aoc"
	"fmt"
	"strconv"
)

const CLOCK = 100

func main() {
	task := aoc.Init(2025, 1)
	// 	task.SetExample(`L50
	// R100
	// L100
	// L10
	// R20
	// L10`)
	input := task.AsStringSlice()

	dial := 50
	stillZeroCount := 0
	movingZeroCount := 0

	for _, cmd := range input {
		direction := cmd[0]
		value, _ := strconv.Atoi(cmd[1:])

		rots := value / CLOCK
		movingZeroCount += rots

		if direction == 'L' {
			value = -value
		}

		newDial := (dial + value)
		if dial != 0 && (dial+(value%CLOCK) > CLOCK || dial+(value%CLOCK) < 0) {
			movingZeroCount++
		}

		dial = newDial % CLOCK

		if dial < 0 {
			dial += CLOCK
		}

		if dial == 0 {
			stillZeroCount++
		}
	}
	fmt.Println("Puzzle 1:", stillZeroCount)
	fmt.Println("Puzzle 2:", stillZeroCount+movingZeroCount)
}

// 6435 7142 6453 7042 6926 6677
