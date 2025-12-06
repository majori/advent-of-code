package main

import (
	"advent-of-code/common/aoc"
	"fmt"
	"regexp"
)

type Node struct {
	Left  string
	Right string
}

func main() {
	task := aoc.Init(2023, 8)
	input := task.AsStringSlice()
	re := regexp.MustCompile(`(\w{3}) = \((\w{3}), (\w{3})\)`)
	nodeMap := make(map[string]Node)
	for i := 2; i < len(input); i++ {
		parts := re.FindStringSubmatch(input[i])
		nodeMap[parts[1]] = Node{
			Left:  parts[2],
			Right: parts[3],
		}
	}

	instructions := input[0]
	node := "AAA"
	steps := 0
	for {
		instruction := instructions[steps%len(instructions)]
		switch instruction {
		case 'L':
			node = nodeMap[node].Left
		case 'R':
			node = nodeMap[node].Right
		}
		steps++

		if node == "ZZZ" {
			fmt.Println("Puzzle 1:", steps)
			break
		}
	}

	nodes := make([]string, 0)
	for k := range nodeMap {
		if k[2] == 'A' {
			nodes = append(nodes, k)
		}
	}

	for {
		instruction := instructions[steps%len(instructions)]
		for i, node := range nodes {
			switch instruction {
			case 'L':
				nodes[i] = nodeMap[node].Left
			case 'R':
				nodes[i] = nodeMap[node].Right
			}
		}
		steps++

		all := true
		for _, node := range nodes {
			if node[2] != 'Z' {
				all = false
			}
		}

		if all {
			fmt.Println("Puzzle 2:", steps)
			break
		}
	}
}
