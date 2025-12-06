package main

import (
	"advent-of-code/common/aoc"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	task := aoc.Init(2025, 5)
	// 	task.SetExample(`3-5
	// 10-14
	// 16-20
	// 12-18

	// 1
	// 5
	// 8
	// 11
	// 17
	// 32`)
	input := task.AsStringSlice()

	ranges := make([][2]int, 0)
	r := true
	sum1, sum2 := 0, 0

	for _, row := range input {
		if row == "" {
			r = false
			ranges = mergeRanges(ranges)
			continue
		}

		if r {
			split := strings.Split(row, "-")
			min, _ := strconv.Atoi(split[0])
			max, _ := strconv.Atoi(split[1])
			ranges = append(ranges, [2]int{min, max})
		} else {
			id, _ := strconv.Atoi(row)
			for _, r := range ranges {
				if id >= r[0] && id <= r[1] {
					sum1++
					break
				}
			}
		}

	}

	fmt.Println("Puzzle 1:", sum1)
	for _, r := range ranges {
		if r[0] > 0 {
			sum2 += r[1] - r[0] + 1
		}
	}

	fmt.Println("Puzzle 2:", sum2)
}

func mergeRanges(rs [][2]int) [][2]int {
	sort.Slice(rs, func(i, j int) bool {
		return rs[i][0] < rs[j][0]
	})

	for i := 1; i < len(rs); i++ {
		for j := 0; j < len(rs); j++ {
			if i == j {
				continue
			}

			if rs[i][0] >= rs[j][0] && rs[i][0] <= rs[j][1] {
				if rs[i][1] > rs[j][1] {
					rs[j][1] = rs[i][1]
				}
				rs[i] = [2]int{}
			}
		}
	}

	n := 0
	for _, r := range rs {
		if r[0] > 0 {
			rs[n] = r
			n++
		}
	}

	return rs[:n]
}
