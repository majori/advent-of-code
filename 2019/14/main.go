package main

import (
	"advent-of-code/common/aoc"
	"fmt"
	"math"
	"regexp"
	"strconv"
)

type Reaction struct {
	out int
	in  map[string]int
}

func main() {
	in := aoc.GetInputRows(2019, 14)

	reactions := make(map[string]Reaction)

	regex := regexp.MustCompile(`(\d+) ([A-Z]+)`)
	for _, reaction := range in {
		in := make(map[string]int)
		matches := regex.FindAllStringSubmatch(reaction, -1)
		for i, match := range matches {
			amount, _ := strconv.Atoi(match[1])
			elem := match[2]

			if i != len(matches)-1 {
				in[elem] = amount
			} else {
				reactions[elem] = Reaction{out: amount, in: in}
			}
		}
	}

	inventory := make(map[string]int)

	convert := func(elem string) {
		if elem == "ORE" {
			return
		}

		amount := inventory[elem]
		reaction := reactions[elem]
		multiplier := float64(1)
		if reaction.out/amount == 0 {
			fmt.Println(1 / (float64(reaction.out) / float64(amount)))
			multiplier = math.Ceil(1 / (float64(reaction.out) / float64(amount)))
		}

		for k, v := range reaction.in {
			inventory[k] += int(multiplier * float64(v))
		}

		delete(inventory, elem)
	}

	inventory["FUEL"] = 1

	for {
		if _, ok := inventory["ORE"]; ok && len(inventory) == 1 {
			break
		}

		keys := make([]string, 0, len(inventory))
		for k := range inventory {
			keys = append(keys, k)
		}

		for _, k := range keys {
			convert(k)
		}
		fmt.Println(inventory)
	}
}
