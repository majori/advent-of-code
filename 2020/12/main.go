package main

import (
	"advent-of-code/common/aoc"
	"fmt"
	"math"
	"strconv"
)

func main() {
	input := aoc.GetInputRows(2020, 12)

	directions := []rune{'N', 'E', 'S', 'W'}

	var x, y, i int
	i = 1

	travelToDirection := func(dir rune, param int) {
		switch dir {
		case 'N':
			y += param
		case 'S':
			y -= param
		case 'E':
			x += param
		case 'W':
			x -= param
		}
	}

	rotate := func(degrees int) {
		i = (i + degrees/90) % len(directions)
		if i < 0 {
			i += 4
		}
	}

	for _, value := range input {
		command := rune(value[0])
		param, _ := strconv.Atoi(value[1:])

		switch command {
		case 'L':
			rotate(-param)
		case 'R':
			rotate(param)
		case 'F':
			travelToDirection(directions[i], param)
		default:
			travelToDirection(command, param)
		}
	}

	fmt.Println("Puzzle 1:", int(math.Abs(float64(y)))+int(math.Abs(float64(x))))

	x, y = 10, 1
	xS, yS := 0, 0

	rotate = func(degrees int) {
		radians := float64(degrees) * math.Pi / 180.0
		px := math.Cos(radians)*float64(x) + math.Sin(radians)*float64(y)
		py := -math.Sin(radians)*float64(x) + math.Cos(radians)*float64(y)

		x = int(math.Round(px))
		y = int(math.Round(py))
	}

	for _, value := range input {
		command := rune(value[0])
		param, _ := strconv.Atoi(value[1:])

		switch command {
		case 'L':
			rotate(-param)
		case 'R':
			rotate(param)
		case 'F':
			xS += x * param
			yS += y * param
		default:
			travelToDirection(command, param)
		}
	}

	fmt.Println("Puzzle 2:", int(math.Abs(float64(yS)))+int(math.Abs(float64(xS))))
}
