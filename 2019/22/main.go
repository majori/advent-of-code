package main

import (
	"advent-of-code/common/aoc"
	"fmt"
	"strconv"
	"strings"
)

type Deck struct {
	Cards []int
}

func main() {
	in := aoc.GetInputRows(2019, 22)

	cards := make([]int, 10007)
	for i := range cards {
		cards[i] = i
	}

	// in := []string{
	// 	"deal with increment 7",
	// 	"deal into new stack",
	// 	"deal into new stack"}

	// cards := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	deck := Deck{cards}

	for _, instruction := range in {
		splitted := strings.Split(instruction, " ")

		if strings.HasPrefix(instruction, "deal into new stack") {
			deck.DealIntoNewStack()
			continue
		}

		param, _ := strconv.Atoi(splitted[len(splitted)-1])

		if strings.HasPrefix(instruction, "deal with increment") {
			deck.DealWithIncrementN(param)
			continue
		}

		if strings.HasPrefix(instruction, "cut") {
			deck.CutNCards(param)
			continue
		}
	}

	for i, card := range deck.Cards {
		if card == 2019 {
			fmt.Printf("Answer 1: %d\n", i)
			break
		}
	}

	for i := 0; i < 101741582076661; i++ {
		if i%10000 == 0 {
			fmt.Println(i)
		}
	}

	fmt.Println("done")

}

func (d *Deck) DealIntoNewStack() {
	size := len(d.Cards)
	temp := make([]int, size)
	for i, card := range d.Cards {
		temp[size-1-i] = card
	}

	d.Cards = temp
}

func (d *Deck) CutNCards(n int) {
	var index int
	if n >= 0 {
		index = n
	} else {
		index = len(d.Cards) + n
	}
	d.Cards = append(d.Cards[index:], d.Cards[:index]...)
}

func (d *Deck) DealWithIncrementN(n int) {
	size := len(d.Cards)
	temp := make([]int, size)
	for i := range temp {
		temp[i] = -1
	}
	index := 0

	for _, card := range d.Cards {
		temp[index] = card
		index = (index + n) % size
	}

	d.Cards = temp
}
