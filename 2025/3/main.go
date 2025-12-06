package main

import (
	"advent-of-code/common/aoc"
	"fmt"
	"math"
)

func main() {
	task := aoc.Init(2025, 3)
	input := task.AsStringSlice()
	banks := make([][]int, len(input))

	for i := range input {
		banks[i] = make([]int, len(input[i]))
		for j := range input[i] {
			n := int(input[i][j] - '0')
			banks[i][j] = n
		}
	}

	sum1, sum2 := 0, 0

	for _, bank := range banks {
		sum1 += calc(bank, 2)
		sum2 += calc(bank, 12)
	}

	fmt.Println("Puzzle 1:", sum1)
	fmt.Println("Puzzle 2:", sum2)
}

func calc(bank []int, n int) int {
	if n == 0 {
		return 0
	}

	x := 0
	for i := 1; i < len(bank)-(n-1); i++ {
		if bank[i] > bank[x] {
			x = i
		}
	}
	return bank[x]*int(math.Pow(10, float64(n-1))) + calc(bank[x+1:], n-1)
}
