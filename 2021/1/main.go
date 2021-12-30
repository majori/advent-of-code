package main

import (
	"advent-of-code/common/aoc"
	"fmt"
)

func main() {
	task := aoc.Init(2021, 1)
	input := task.AsIntSlice()

	var inc, dec int
	
	for i := 1; i < len(input); i++ {
		if input[i] > input[i-1] {
			inc++
		}
		if input[i] < input[i-1] {
			dec++
		}
	}

	fmt.Println("Puzzle 1: ", inc)
	
	inc = 0
	dec = 0
	
	for i := 1; i < len(input)-2; i++ {
		sum1 := input[i-1] + input[i] + input[i+1]
		sum2 := input[i] + input[i+1] + input[i+2]
		if sum2 > sum1 {
			inc++
		}
		if sum2 < sum1 {
			dec++
		}
	}
	
	fmt.Println("Puzzle 2: ", inc)
}
