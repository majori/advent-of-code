package main

import (
	"advent-of-code/common/aoc"
	"fmt"
	"strconv"
)

const free = -1

func main() {
	task := aoc.Init(2024, 9)
	input := task.AsString()

	diskMap := make([]int, len(input))
	for i, r := range input {
		diskMap[i], _ = strconv.Atoi(string(r))
	}

	idx := 0
	disk := make([]int, 0)
	for i := range diskMap {
		if i%2 == 0 {
			for j := 0; j < diskMap[i]; j++ {
				disk = append(disk, idx)
			}
			idx++
		} else {
			for j := 0; j < diskMap[i]; j++ {
				disk = append(disk, free)
			}
		}
	}

	// print(disk)
	defrag1(&disk)

	fmt.Println("Puzzle 1:", checksum(disk))
	fmt.Println("Puzzle 2:", "")
}

func defrag1(disk *[]int) {
	for i := len(*disk) - 1; i >= 0; i-- {
		for j := 0; j < len(*disk); j++ {
			if (*disk)[j] == free {
				(*disk)[j] = (*disk)[i]
				(*disk) = (*disk)[:i]
				break
			}

			// No more free space
			if j == len(*disk)-1 {
				return
			}
		}
	}
}

func checksum(disk []int) int {
	checksum := 0
	for i, d := range disk {
		if d == free {
			continue
		}
		checksum += i * d
	}
	return checksum
}

func print(disk []int) {
	for _, d := range disk {
		if d == free {
			fmt.Print(".")
		} else {
			fmt.Print(d)
		}
	}
	fmt.Println()
}
