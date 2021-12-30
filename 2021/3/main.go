package main

import (
	"advent-of-code/common/aoc"
	"fmt"
	"strconv"
)

func main() {
	task := aoc.Init(2021, 3)
	input := task.AsStringSlice()

	const bitSize int = 12
	numbers := make([]int, len(input))

	for i, row := range input {
		x, _ := strconv.ParseInt(row, 2, 0)
		numbers[i] = int(x)
	}
	
	mostCommonBits := make([]int, bitSize)
	
	for _, row := range numbers {	
		for i := 0; i < bitSize; i++ {
			if row & (1 << i) != 0 {
				mostCommonBits[i]++
			} else {
				mostCommonBits[i]--
			}
		}
	}

	gamma, epsilon := 0, 0
	for i, column := range mostCommonBits {
		if column > 0 {
			gamma = gamma | (1 << i)
			mostCommonBits[i] = 1
		} else {
			epsilon = epsilon | (1 << i)
			mostCommonBits[i] = 0
		}
	}

	fmt.Println("Puzzle 1:", gamma*epsilon)
	fmt.Println(mostCommonBits)
	
	oxygenCandidates := make([]int, len(numbers))
	copy(oxygenCandidates, numbers)
	
	for i := 0; len(oxygenCandidates) > 1; i++ {
		temp := make([]int, 0, len(oxygenCandidates))
		for _, candidate := range oxygenCandidates {
			if (candidate & (1 << (11-i)) > 0) == (mostCommonBits[11-i] == 1) {
				temp = append(temp, candidate)
			}
		}
		oxygenCandidates = temp
	}

	oxygenRating := oxygenCandidates[0]
	
	co2Candidates := make([]int, len(numbers))
	copy(co2Candidates, numbers)
	
	for i := 0; len(co2Candidates) > 1; i++ {
		temp := make([]int, 0, len(co2Candidates))
		for _, candidate := range co2Candidates {
			fmt.Printf("%012b\n", candidate)
			fmt.Printf("%012b\n", (1 << (11-i)))
			fmt.Printf("%v = %v", candidate & (1 << (11-i)) > 0, (mostCommonBits[11-i] == 1))
			if (candidate & (1 << (11-i)) > 0) != (mostCommonBits[11-i] == 1) {
				temp = append(temp, candidate)
			}
		}
		co2Candidates = temp
	}

	co2Rating := co2Candidates[0]

	fmt.Println("Puzzle 2: ", oxygenRating*co2Rating)
}
