package main

import (
	"advent-of-code/common/aoc"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	task := aoc.Init(2024, 3)
	input := fmt.Sprintf("do()%s", task.AsString())

	regex := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := regex.FindAllStringSubmatch(input, -1)

	sum1 := 0
	for _, match := range matches {
		a, _ := strconv.Atoi(match[1])
		b, _ := strconv.Atoi(match[2])
		sum1 += a * b
	}

	fmt.Println("Puzzle 1:", sum1)

	regex2 := regexp.MustCompile(`don't\(\).*?do\(\)`)
	new2 := regex2.ReplaceAllString(input, "")
	matches2 := regex.FindAllStringSubmatch(new2, -1)

	sum2 := 0
	for _, match := range matches2 {
		a, _ := strconv.Atoi(match[1])
		b, _ := strconv.Atoi(match[2])
		sum2 += a * b
	}

	fmt.Println("Puzzle 2:", sum2)
	// 103833649
}
