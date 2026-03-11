package main

import (
	"advent-of-code/common/aoc"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

type Node []int

type Line struct {
	d  int
	n1 *Node
	n2 *Node
}

func main() {
	task := aoc.Init(2025, 8)
	task.SetExample(`1,1,1
1,2,3
1,1,2
10,10,10
10,11,12
10,10,11`)
	input := task.AsStringSlice()
	result1, result2 := 0, 0

	nodes := make([]Node, len(input))
	distances := make(map[*Node]map[*Node]int)
	lines := make([]Line, 0)
	graphs := make([]map[*Node]*Node, 0)

	for i, row := range input {
		node := Node{}
		s := strings.Split(row, ",")
		for _, dimension := range s {
			d, _ := strconv.Atoi(dimension)
			node = append(node, d)
		}

		nodes[i] = node
	}

	for i := range nodes {
		distances[&nodes[i]] = make(map[*Node]int)
		for j := range nodes {
			if i == j {
				continue
			}
			d := dist(nodes[i], nodes[j])
			if distances[&nodes[j]] == nil || distances[&nodes[j]][&nodes[i]] == 0 {
				lines = append(lines, Line{d, &nodes[i], &nodes[j]})
			}

			distances[&nodes[i]][&nodes[j]] = d
		}
	}

	slices.SortFunc(lines, func(a, b Line) int {
		if a.d > b.d {
			return 1
		} else if a.d < b.d {
			return -1
		} else {
			return 0
		}
	})

	for _, line := range lines {
		found := false
		for i := range graphs {
			if graphs[i][line.n1] != nil || graphs[i][line.n2] != nil  {
				found = true
				if graphs[i][line.n1] == nil {
					graphs[i][line.n1] = line.n2
				} else {
					graphs[i][line.n2] = line.n1
				}
				break
			}
		}
		if !found {
			newGraph := make(map[*Node]*Node)
			newGraph[line.n1] = line.n2
			newGraph[line.n2] = line.n1
			graphs = append(graphs, newGraph)
		}
	}

	fmt.Println(graphs)

	fmt.Println("Puzzle 1:", result1)
	fmt.Println("Puzzle 2:", result2)
}

func dist(n1, n2 Node) (d int) {
	for i := range n1 {
		d += int(math.Abs(float64(n1[i] - n2[i])))
	}
	return
}
