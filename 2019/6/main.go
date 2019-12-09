package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Node struct {
	name     string
	children []*Node
	parent   *Node
}

func main() {
	rawInput, _ := ioutil.ReadFile("./input.txt")
	input := strings.Split(string(rawInput), "\n")

	lookup := make(map[string]*Node)
	for _, orbit := range input {
		planets := strings.Split(orbit, ")")
		parent, child := planets[0], planets[1]

		if _, ok := lookup[child]; !ok {
			lookup[child] = &Node{child, []*Node{}, nil}
		}

		if _, ok := lookup[parent]; !ok {
			lookup[parent] = &Node{parent, []*Node{lookup[child]}, nil}
		} else {
			lookup[parent].children = append(lookup[parent].children, lookup[child])
		}

		lookup[child].parent = lookup[parent]
	}

	fmt.Printf("Part 1: %d\n", calculateDepth(lookup["COM"], 0))

	santa := pathFromRoot(lookup["SAN"])
	you := pathFromRoot(lookup["YOU"])

	for i := 0; i < len(you); i++ {
		if you[i] == santa[i] {
			continue
		}

		fmt.Printf("Part 2: %d", (len(santa)-i-1)+(len(you)-i-1))
		break
	}
}

func calculateDepth(node *Node, currentDepth int) int {
	depth := currentDepth

	for _, child := range node.children {
		depth += calculateDepth(child, currentDepth+1)
	}

	return depth
}

func pathFromRoot(node *Node) []*Node {
	path := []*Node{node}

	currentNode := node
	for currentNode.parent != nil {
		path = append(path, currentNode.parent)
		currentNode = currentNode.parent
	}

	// Reverse path so root comes first
	for i := len(path)/2 - 1; i >= 0; i-- {
		opp := len(path) - 1 - i
		path[i], path[opp] = path[opp], path[i]
	}

	return path
}
