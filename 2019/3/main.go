package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	dat, _ := ioutil.ReadFile("./input.txt")
	paths := strings.Split(string(dat), "\n")
	path1 := strings.Split(paths[0], ",")
	path2 := strings.Split(paths[1], ",")

	type Coordinate struct {
		x int
		y int
	}

	type Visit struct {
		count    int
		length   int
		pathMask int
	}

	visits := make(map[Coordinate]*Visit)

	for p, path := range [][]string{path1, path2} {
		currentX, currentY := 0, 0
		totalLength := 0

		for _, step := range path {
			length, _ := strconv.Atoi(step[1:])

			for i := 1; i <= length; i++ {
				totalLength++

				switch step[0] {
				case 'U':
					currentY++
					break
				case 'D':
					currentY--
					break
				case 'R':
					currentX++
					break
				case 'L':
					currentX--
					break
				}

				coord := Coordinate{currentX, currentY}
				if visits[coord] == nil {
					visits[coord] = &Visit{1, totalLength, p}
				} else if visits[coord].pathMask != p {
					visits[coord].count++
					visits[coord].pathMask += p
					visits[coord].length += totalLength
				}
			}
		}
	}

	minCount := 1000000
	minLength := 1000000
	for key, value := range visits {
		if value.count > 1 {
			distance := int(math.Abs(float64(key.x))) + int(math.Abs(float64(key.y)))
			if distance < minCount {
				minCount = distance
			}
			if value.length < minLength {
				minLength = value.length
			}
		}
	}
	fmt.Printf("Part 1: %d\n", minCount)
	fmt.Printf("Part 2: %d\n", minLength)
}
