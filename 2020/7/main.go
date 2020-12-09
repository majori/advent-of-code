package main

import (
	"advent-of-code/common/aoc"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	input := aoc.GetInput(2020, 7)

	reg1 := regexp.MustCompile(`((?:\w+ ?)+) bags contain((?:(?: \d+)? (?:(?:\w+ ?)+) bags?,?)+)\.`)
	reg2 := regexp.MustCompile(`(\d+) (\w+ \w+)`)

	entries := reg1.FindAllStringSubmatch(input, -1)

	nodes1 := make(map[string]map[string]int)
	nodes2 := make(map[string]map[string]int)

	for _, row := range entries {
		node := row[1]
		for _, match := range row[2:] {
			bags := reg2.FindAllStringSubmatch(match, -1)
			if len(bags) == 0 {
				continue
			}

			out := make(map[string]int)
			for _, match := range bags {
				amount, _ := strconv.Atoi(match[1])
				bag := match[2]
				out[bag] = amount

				if nodes1[bag] == nil {
					nodes1[bag] = make(map[string]int)
				}

				nodes1[bag][node] += amount
			}
			nodes2[node] = out
		}
	}

	checked := make(map[string]bool)

	var recursive func(node string) (count int)
	recursive = func(node string) (count int) {
		for key := range nodes1[node] {
			if _, ok := checked[key]; !ok {
				checked[key] = true
				count++
				count += recursive(key)
			}
		}
		return
	}

	fmt.Println("Puzzle 1: ", recursive("shiny gold"))

	recursive = func(node string) (count int) {
		for key, value := range nodes2[node] {
			count += value
			count += value * recursive(key)
		}
		return
	}

	fmt.Println("Puzzle 2: ", recursive("shiny gold"))
}
