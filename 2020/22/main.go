package main

import (
	"advent-of-code/common/aoc"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input := aoc.GetInput(2020, 22)

	decks := make([][]int, 0)

	regex := regexp.MustCompile(`Player \d:\n((?:\d+\n)+)`)
	raw := regex.FindAllString(input, 2)

	for _, p := range raw {
		player := make([]int, 0)
		for _, card := range strings.Split(p, "\n")[1:] {
			c, err := strconv.Atoi(card)
			if err != nil {
				continue
			}
			player = append(player, c)
		}
		decks = append(decks, player)
	}

play:
	for {
		for _, deck := range decks {
			if len(deck) == 0 {

				break play
			}
		}
		round(decks)
	}

	sum := 0
	for i, card := range decks[0] {
		sum += (len(decks[0]) - i) * card
	}
	fmt.Println("Puzzle 1:", sum)
}

func round(decks [][]int) {
	tops := make([]int, len(decks))
	var max, winner int
	for i := 0; i < len(decks); i++ {
		top := decks[i][0]
		tops[i] = top
		if decks[i][0] > max {
			max = top
			winner = i
		}
	}

	decks[winner] = append(decks[winner][1:], decks[winner][0])
	for i := 0; i < len(decks); i++ {
		if i != winner {
			decks[winner] = append(decks[winner], decks[i][0])
			decks[i] = decks[i][1:]
		}
	}
}
