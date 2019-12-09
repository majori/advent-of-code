package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	dat, _ := ioutil.ReadFile("./input.txt")
	elements := strings.Split(string(dat), ",")

	original := []int{}
	for _, line := range elements {
		value, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		original = append(original, value)
	}

	codes := round(original, 12, 2)

	fmt.Printf("Part 1: %d\n", codes[0])

	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			codes = round(original, i, j)

			if codes[0] == 19690720 {
				fmt.Printf("Part 2: %d\n", 100*codes[1]+codes[2])
			}
		}
	}
}

func round(arr []int, p1 int, p2 int) []int {
	index := 0
	codes := make([]int, len(arr))
	copy(codes, arr)

	codes[1] = p1
	codes[2] = p2

	for codes[index] != 99 {
		switch codes[index] {
		case 1:
			point1, point2, point3 := codes[index+1], codes[index+2], codes[index+3]
			codes[point3] = codes[point1] + codes[point2]
			break
		case 2:
			point1, point2, point3 := codes[index+1], codes[index+2], codes[index+3]
			codes[point3] = codes[point1] * codes[point2]
			break
		}

		index += 4
	}

	return codes
}
