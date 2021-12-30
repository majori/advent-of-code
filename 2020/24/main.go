package main

import (
	"advent-of-code/common/aoc"
	"fmt"
	"regexp"
)

func main() {
	input := aoc.GetInputRows(2020, 24)

	regex := regexp.MustCompile(`(se|sw|ne|nw|e|w)`)

	tiles := make(map[string]bool)

	for _, tile := range input {
		var col, row int
		for _, dir := range regex.FindAllString(tile, -1) {
			switch dir {
			case "e":
				col++
			case "se":
				if row%2 != 0 {
					col++
				}
				row++
			case "sw":
				if row%2 == 0 {
					col--
				}
				row++
			case "w":
				col--
			case "ne":
				if row%2 != 0 {
					col++
				}
				row--
			case "nw":
				if row%2 == 0 {
					col--
				}
				row--
			}
		}
		key := fmt.Sprintf("%d,%d", col, row)
		if _, ok := tiles[key]; !ok {
			tiles[key] = true
		} else {
			delete(tiles, key)
		}
	}

	fmt.Println("Puzzle 1:", len(tiles))
}
