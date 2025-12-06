package main

import (
	"advent-of-code/common/aoc"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	task := aoc.Init(2025, 2)
	task.SetExample("11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124")
	input := task.AsString()

	raw := strings.Split(strings.TrimSpace(input), ",")
	ranges := make([][2]int, len(raw))

	for i, r := range raw {
		rs := strings.Split(r, "-")
		ranges[i][0], _ = strconv.Atoi(rs[0])
		ranges[i][1], _ = strconv.Atoi(rs[1])
	}

	matches1 := 0
	matches2 := 0

	for _, r := range ranges {
		for nid := r[0]; nid <= r[1]; nid++ {
			id := strconv.Itoa(nid)

			for i := 1; i <= len(id)/2; i++ {
				if len(id)%i != 0 {
					continue
				}

				chunks := slices.Chunk(strings.Split(id, ""), i)

				v := ""
				match := true
				for chunk := range chunks {
					c := strings.Join(chunk, "")
					if v == "" {
						v = c
					}

					if v != c {
						match = false
						continue
					}
				}

				if match {
					if len(id)%2 == 0 && i == len(id)/2 {
						matches1 += nid
					}
					matches2 += nid
					break
				}
			}
		}
	}

	fmt.Println("Puzzle 1:", matches1)
	fmt.Println("Puzzle 2:", matches2)
}
