package main

import (
	"advent-of-code/common/aoc"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	in := aoc.GetInput(2020, 2)
	var valid1, valid2 int

	reg := regexp.MustCompile(`(\d+)-(\d+) ([a-z]): (\w+)`)
	entries := reg.FindAllStringSubmatch(in, -1)

	for _, entry := range entries {
		lower, _ := strconv.Atoi(entry[1])
		upper, _ := strconv.Atoi(entry[2])
		char := []rune(entry[3])[0]
		password := entry[4]

		dict := make(map[rune]int)
		for _, x := range password {
			dict[x]++
		}

		if lower <= dict[char] && dict[char] <= upper {
			valid1++
		}

		if (rune(password[lower-1]) == char && rune(password[upper-1]) != char) || (rune(password[lower-1]) != char && rune(password[upper-1]) == char) {
			valid2++
		}
	}

	fmt.Println("Puzzle 1: ", valid1)
	fmt.Println("Puzzle 2: ", valid2)
}
