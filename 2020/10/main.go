package main

import (
	"advent-of-code/common/aoc"
	"fmt"
	"sort"
)

func main() {
	input := aoc.GetInputRowsAsInt(2020, 10)

	adapters := sort.IntSlice(input)
	adapters.Sort()

	adapters = append(adapters, adapters[len(adapters)-1]+3) // Built-in
	adapters = append(sort.IntSlice{0}, adapters...)         // Source

	diffs := make(map[int]int)
	for i := 0; i < len(adapters)-1; i++ {
		diffs[adapters[i+1]-adapters[i]]++
	}

	fmt.Println("Puzzle 1:", diffs[1]*diffs[3])
	fmt.Println("Puzzle 2:", travel(adapters))
}

var cache map[string]int = make(map[string]int)

func travel(slice sort.IntSlice) int {
	if len(slice) == 1 {
		return 1
	}

	branches := 0
	for i := 1; i < len(slice); i++ {
		cacheKey := fmt.Sprintf("%d:%d", slice[0], slice[i])
		if slice[i]-slice[0] <= 3 {
			subBranches, ok := cache[cacheKey]
			if !ok {
				subBranches = travel(slice[i:])
				cache[cacheKey] = subBranches
			}
			branches += subBranches
		} else {
			break
		}
	}
	return branches
}
