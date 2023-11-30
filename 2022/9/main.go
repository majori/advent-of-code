package main

import (
	"advent-of-code/common/aoc"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

func main() {
	task := aoc.Init(2022, 9)
	task.SetExample(`R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`)
	input := task.AsStringSlice()

	// fmt.Println("Puzzle 1: ", len(simulate(input, make([]Point, 2))))
	fmt.Println("Puzzle 2: ", len(simulate(input, make([]Point, 10))))
}

func isPointClose(a, b Point) bool {
	return math.Abs(float64(a.X-b.X)) <= 1.0 && math.Abs(float64(a.Y-b.Y)) <= 1.0
}

func simulate(commands []string, points []Point) map[string]bool {
	tailPath := make(map[string]bool)
	tailPath["0,0"] = true

	for _, command := range commands {
		parts := strings.Split(command, " ")
		dir := parts[0]
		amount, _ := strconv.Atoi(parts[1])

		for n := amount; n > 0; n-- {
			old := make([]Point, len(points))
			copy(old, points)
			switch dir {
			case "U":
				points[0].Y++
			case "D":
				points[0].Y--
			case "R":
				points[0].X++
			case "L":
				points[0].X--
			}

			for i := 1; i < len(points); i++ {
				if !isPointClose(points[i-1], points[i]) {
					points[i] = old[i-1]
					if i == len(points)-1 {
						tailPath[fmt.Sprintf("%d,%d", points[i].X, points[i].Y)] = true
					}
				}
			}
		}
	}

	return tailPath
}
