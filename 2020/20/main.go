package main

import (
	"advent-of-code/common/aoc"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input := aoc.GetInput(2020, 20)

	blocks := make(map[int][][]bool)
	regex := regexp.MustCompile(`Tile (\d+):\n((?:.+\n)+)`)
	matches := regex.FindAllStringSubmatch(input, -1)
	for _, match := range matches {
		id, _ := strconv.Atoi(match[1])
		rows := make([][]bool, 0)
		for _, row := range strings.Split(match[2], "\n") {
			if len(row) == 0 {
				continue
			}

			pixels := make([]bool, len(row))
			for i, c := range row {
				pixels[i] = c == '#'
			}
			rows = append(rows, pixels)
		}
		blocks[id] = rows
	}

	x := make(map[int][]int)

	for id, block := range blocks {
		sides := [4][]bool{}
		sides[0] = block[0]
		sides[2] = block[len(block)-1]
		left := make([]bool, 0, len(block[0]))
		right := make([]bool, 0, len(block[0]))
		for _, row := range block {
			left = append(left, row[0])
			right = append(right, row[len(row)-1])
		}
		sides[1] = right
		sides[3] = left

		for _, side := range sides {
			hash := boolSliceToInt(side)
			x[hash] = append(x[hash], id)
			reversed := make([]bool, len(side))
			for i := len(side) - 1; i >= 0; i-- {
				reversed[len(side)-1-i] = side[i]
			}

			hash = boolSliceToInt(reversed)
			x[hash] = append(x[hash], id)
		}

	}

	z := make(map[int]int)

	for _, y := range x {
		if len(y) != 1 {
			continue
		}
		z[y[0]]++
	}

	mul := 1
	for id, n := range z {
		if n == 4 {
			mul *= id
		}
	}

	fmt.Println("Puzzle 1:", mul)
}

func boolSliceToInt(s []bool) (b int) {
	for i, pixel := range s {
		var bit int
		if pixel {
			bit = 1
		}
		b |= (bit << i)
	}
	return
}
