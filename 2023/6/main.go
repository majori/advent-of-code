package main

import "fmt"

func main() {
	// task := aoc.Init(2023, 6)
	// input := task.AsString()

	// Too lazy to parse the input
	input := [][]int{
		{
			45, 295,
		},
		{
			98, 1734,
		},
		{
			83, 1278,
		},
		{
			73, 1210,
		},
	}

	p1 := 1
	for _, v := range input {
		p1 *= CalculateAmountOfWinningStrategies(v[0], v[1])
	}

	fmt.Println("Puzzle 1:", p1)
	fmt.Println("Puzzle 2:", CalculateAmountOfWinningStrategies(45988373, 295173412781210))
}

func CalculateAmountOfWinningStrategies(t, d int) (n int) {
	for i := 1; i < t; i++ {
		if (i * (t - i)) >= d {
			n++
		}
	}
	return
}
