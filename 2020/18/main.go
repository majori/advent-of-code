package main

import (
	"advent-of-code/common/aoc"
	"fmt"
	"strconv"
)

func main() {
	input := aoc.GetInputRows(2020, 18)

	sum := 0
	for _, row := range input {
		total, _ := calculate1(row)
		sum += total
	}

	fmt.Println("Puzzle 1:", sum)

	sum = 0
	for _, row := range input {
		total, _ := calculate2(row)
		sum += total
	}

	fmt.Println("Puzzle 2:", sum)
}

func calculate1(in string) (acc int, i int) {
	var operator func(int, int) int = sum

	for i = 0; i < len(in); i++ {
		c := in[i]

		switch c {
		case ' ':
			continue
		case '+':
			operator = sum
		case '*':
			operator = multiply
		case '(':
			part, j := calculate1(in[i+1:])
			i += j + 1
			acc = operator(acc, part)
		case ')':
			return acc, i
		default:
			d, _ := strconv.Atoi(string(c))
			acc = operator(acc, d)
		}
	}

	return acc, len(in) - 1
}

func calculate2(in string) (acc int, i int) {
	var operator func(int, int) int = sum

	for i = 0; i < len(in); i++ {
		c := in[i]

		switch c {
		case ' ':
			continue
		case '+':
			operator = sum
		case '*':
			operator = multiply
		case '(':
			part, j := calculate2(in[i+1:])
			i += j + 1
			acc = operator(acc, part)
		case ')':
			return acc, i
		default:
			d, _ := strconv.Atoi(string(c))
			acc = operator(acc, d)
		}
	}

	return acc, len(in) - 1
}

func sum(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}
