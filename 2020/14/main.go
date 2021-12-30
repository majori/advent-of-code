package main

import (
	"advent-of-code/common/aoc"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const SIZE = 36

func main() {
	input := aoc.GetInput(2020, 14)
	regex := regexp.MustCompile(`(?:mask|mem\[(\d+)\]) = (.*)`)

	matches := regex.FindAllStringSubmatch(input, -1)
	memory := make(map[uint64]uint64)
	var mask []rune
	for _, row := range matches {
		switch {
		case strings.HasPrefix(row[0], "mask"):
			mask = []rune(row[2])

		case strings.HasPrefix(row[0], "mem"):
			address, _ := strconv.ParseUint(row[1], 10, SIZE)
			value, _ := strconv.ParseUint(row[2], 10, SIZE)

			for i, x := range mask {
				switch x {
				case '1':
					value |= (1 << (len(mask) - 1 - i))
				case '0':
					value &= ^(1 << (len(mask) - 1 - i))
				}
			}

			memory[address] = value
		}
	}

	var sum uint64
	for _, value := range memory {
		sum += value
	}

	fmt.Println("Puzzle 1:", sum)

	memory = make(map[uint64]uint64)
	mask = []rune{}

	for _, row := range matches {
		switch {
		case strings.HasPrefix(row[0], "mask"):
			mask = []rune(row[2])

		case strings.HasPrefix(row[0], "mem"):
			address, _ := strconv.ParseUint(row[1], 10, SIZE)
			value, _ := strconv.ParseUint(row[2], 10, SIZE)

			for i, c := range mask {
				if c == '0' {
					bit := address >> (SIZE - 1 - i) & 1
					if bit == 1 {
						mask[i] = '1'
					} else {
						mask[i] = '0'
					}
				}
			}

			var perm func(str []rune) []uint64
			perm = func(str []rune) []uint64 {
				for i, c := range str {
					if c == 'X' {
						perm(str[i+1:])

					} else {
						continue
					}
				}
			}

			addressPermutations := perm(mask)

			// TODO

			for _, address := range addressPermutations {
				memory[address] = value
			}
		}
	}
}
