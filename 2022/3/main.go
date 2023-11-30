package main

import (
	"advent-of-code/common/aoc"
	"fmt"
	"strings"
)

const priority = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const groupSize = 3

func main() {
	task := aoc.Init(2022, 3)
	input := task.AsStringSlice()

	var p1 int
	for _, row := range input {
		for i := len(row) / 2; i < len(row); i++ {
			letter := rune(row[i])
			if strings.ContainsRune(row[:len(row)/2], letter) {
				p1 += strings.IndexRune(priority, letter) + 1
				break
			}
		}
	}

	fmt.Println("Puzzle 1:", p1)

	var p2 int
	groups := make([][groupSize]string, len(input)/groupSize)
	for n := 0; n < len(input)/groupSize; n++ {
		for m := 0; m < groupSize; m++ {
			groups[n][m] = input[n*groupSize+m]
		}
	}

	for _, group := range groups {
		for _, r := range group[0] {
			if strings.ContainsRune(group[1], r) && strings.ContainsRune(group[2], r) {
				p2 += strings.IndexRune(priority, r) + 1
				break
			}
		}
	}

	fmt.Println("Puzzle 2:", p2)
}
