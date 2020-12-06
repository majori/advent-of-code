package main

import (
	"advent-of-code/common/aoc"
	"fmt"
	"math"
	"sort"
)

func main() {
	input := aoc.GetInputRows(2020, 5)

	seats := make([]int, 0, len(input))
	max := 0
	for _, pass := range input {
		id := GetSeatID(pass)
		seats = append(seats, id)
		if id > max {
			max = id
		}
	}

	fmt.Println("Puzzle 1: ", max)

	var sorted sort.IntSlice
	sorted = seats
	sorted.Sort()

	for i := 1; i < len(sorted); i++ {
		if sorted[i]-sorted[i-1] == 2 {
			fmt.Println("Puzzle 2: ", sorted[i]-1)
		}
	}

}

func GetSeatID(pass string) int {
	low := 0
	high := 127

	for _, dir := range pass[:7] {
		low, high = binarySearch(dir, low, high)
	}

	if low != high {
		panic("Binary search failed")
	}

	row := low

	low = 0
	high = 7
	for _, dir := range pass[7:] {
		low, high = binarySearch(dir, low, high)
	}

	if low != high {
		panic("Binary search failed")
	}

	column := low

	return row*8 + column
}

func binarySearch(dir rune, low, high int) (int, int) {
	center := int(math.Ceil(float64(high-low) / 2.0))
	switch dir {
	case 'F', 'L':
		high -= center
	case 'B', 'R':
		low += center
	}

	return low, high
}
