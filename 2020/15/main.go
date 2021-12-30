package main

import (
	"advent-of-code/common/aoc"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := aoc.GetInput(2020, 15)
	initial := strings.Split(strings.TrimSpace(input), ",")
	mem := make(map[int][]int)
	var turn, lastSpoken int

	for i, v := range initial {
		d, _ := strconv.Atoi(v)
		mem[d] = append(mem[d], i)
		lastSpoken = d
		turn = i
	}

	for {
		turn++

		if len(mem[lastSpoken]) <= 1 {
			lastSpoken = 0
		} else {
			occurences := mem[lastSpoken]
			lastSpoken = occurences[len(occurences)-1] - occurences[len(occurences)-2]
		}

		mem[lastSpoken] = append(mem[lastSpoken], turn)

		if turn == 2020-1 {
			fmt.Println("Puzzle 1:", lastSpoken)
		}
		if turn == 30_000_000-1 {
			fmt.Println("Puzzle 2:", lastSpoken)
			return
		}
	}

}
