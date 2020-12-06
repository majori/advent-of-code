package main

import (
	"advent-of-code/common/aoc"
	"fmt"
)

func main() {
	in := aoc.GetInputRowsAsInt(2020, 1)

	for i, elem := range in {
		compare := in[i:]
		for _, y := range compare {
			if elem+y == 2020 {
				fmt.Println("Puzzle 1: ", elem*y)
			}
		}
	}

	x := in[:]

	for i, a := range x {
		y := x[i:]
		for j, b := range y {
			z := y[j:]
			for _, c := range z {
				if a+b+c == 2020 {
					fmt.Println("Puzzle 2: ", a*b*c)
				}
			}
		}
	}
}
