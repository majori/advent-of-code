package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("./input.txt")

	boundaries := strings.Split(string(input), "-")
	lower, _ := strconv.Atoi(boundaries[0])
	upper, _ := strconv.Atoi(boundaries[1])

	possible1 := 0
	possible2 := 0
	for candidate := lower; candidate <= upper; candidate++ {
		digits := []int{
			(candidate / 100000) % 10,
			(candidate / 10000) % 10,
			(candidate / 1000) % 10,
			(candidate / 100) % 10,
			(candidate / 10) % 10,
			candidate % 10,
		}

		adj, inc := false, true
		adjCount := make(map[int]int)
		for i, digit := range digits {
			if i == len(digits)-1 {
				break
			}

			if digit == digits[i+1] {
				adj = true
				if adjCount[digit] == 0 {
					adjCount[digit] = 2
				} else {
					adjCount[digit]++
				}
			}

			if digit > digits[i+1] {
				inc = false
			}
		}

		if inc {
			if adj {
				possible1++
			}

			for _, value := range adjCount {
				if value == 2 {
					possible2++
					break
				}
			}
		}
	}

	fmt.Printf("Part 1: %d\n", possible1)
	fmt.Printf("Part 2: %d", possible2)
}
