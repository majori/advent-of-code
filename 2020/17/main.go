package main

import (
	"advent-of-code/common/aoc"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	task := aoc.Init(2020, 17)
	task.SetExample(`
.#.
..#
###
	`)
	input := task.AsStringSlice()

	fmt.Println(input)
	initialGrid := make(map[string]bool)

	for y, row := range input {
		for x, c := range row {
			initialGrid[hashKey(x, y, 0)] = c == '#'
		}
	}

	grid := initialGrid
	for i := 0; i < 5; i++ {
		grid = play(grid, 3)
	}
	
	var counter int
	for _, active := range grid {
		if active {
			counter++
		}
	}
	
	fmt.Println("Puzzle 1:", counter)
	
	// grid = initialGrid
	// for i := 0; i < 5; i++ {
	// 	grid = play(grid, 3)
	// }
}

func play(initial map[string]bool, dimensionCount int) map[string]bool {
	grid := initial
	distance := []int{-1, 0, 1}

	for coordinateHash, active := range grid {
		coordinates := extract(coordinateHash)
		fmt.Println(coordinates)
		fmt.Println(active)

		for a := 0; a < len(coordinates); a++ {
			diff := make([]int, len(coordinates))
			for _, d := range distance {	
				
			}
		}

		// 	}
		// 	for _, diffX := range distance {
		// 		checkX := x + diffX
		// 		for _, diffY := range distance {
		// 			checkY := y + diffY
		// 			for _, diffZ := range distance {
		// 				checkZ := z + diffZ
	
		// 				if diffX == 0 && diffY == 0 && diffZ == 0 {
		// 					continue
		// 				}
	
		// 				key := hashKey(checkX, checkY, checkZ, 0)
		// 				if _, ok := grid[key]; !ok && active {
		// 					grid[key] = false
		// 				}
		// 			}
		// 		}
		// 	}
		// }
	
		// for coord, active := range grid {
		// 	x, y, z, _ := extract(coord)
		// 	neighbours := make(map[bool]int)
		// 	for _, diffX := range distance {
		// 		checkX := x + diffX
		// 		for _, diffY := range distance {
		// 			checkY := y + diffY
		// 			for _, diffZ := range distance {
		// 				checkZ := z + diffZ
	
		// 				if diffX == 0 && diffY == 0 && diffZ == 0 {
		// 					continue
		// 				}
	
		// 				key := hashKey(checkX, checkY, checkZ, 0)
		// 				neighbours[grid[key]]++
		// 			}
		// 		}
		// 	}
	
		// 	if active {
		// 		grid[coord] = neighbours[true] == 2 || neighbours[true] == 3
		// 	} else {
		// 		if neighbours[true] == 0 {
		// 			continue
		// 		}
		// 		grid[coord] = neighbours[true] == 3
		// 	}
		// }
	}

	return grid
}

func getNeighbours(coordinateHash string) []string {
	coordinates := extract(coordinateHash)
	distance := []int{-1, 0, 1}
	neighbours := make([]string, 0)

	for i, c := range coordinates {
		temp := coordinates
		for _, d := range distance {
			temp[i] = c-d
			neighbours = append(neighbours, hashKey(c-d))
		}
	}
}

func hashKey(d... int) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(d)), ","), "[]")
}

func extract(hash string) []int {
	splitted := strings.Split(hash, ",")
	coordinates := make([]int, len(splitted))
	for i, axis := range splitted {
		a, _ := strconv.Atoi(axis)
		coordinates[i] = a 
	}
 	return coordinates
}