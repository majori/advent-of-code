package main

import (
	"advent-of-code/common/aoc"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	task := aoc.Init(2020, 23)
	task.SetExample("389125467")
	input := strings.Split(task.AsString(), "")
	
	cards := make([]int, len(input))
	for i, c := range input {
		cards[i], _ = strconv.Atoi(c)
	}

	
	index := 0
	for n := 0; n < 10; n++ {
		fmt.Println(cards, index)
		cards, index = move(cards, index)
	}
	
	fmt.Println(cards, index)
}

func move(cards []int, currentIndex int) ([]int, int) {
	current := cards[currentIndex]
	n := 3
	pickup := make([]int, 3)
	pickupIndex := (currentIndex + 1) % len(cards) 
	
	if len(cards) - pickupIndex >= n {
		copy(pickup, cards[pickupIndex:pickupIndex+n])
		cards = append(cards[:pickupIndex], cards[pickupIndex+3:]...)
	} else {
		cardsFromStart := n - (len(cards) - pickupIndex)
		copy(pickup, append(cards[pickupIndex:], cards[:cardsFromStart]...))
		cards = cards[cardsFromStart:pickupIndex]
	}

	fmt.Println(pickup)

	var destinationCardIndex int
	destinationCard := current
	found := false

	for !found {
		destinationCard--
		if destinationCard <= 0 {
			destinationCard = len(cards)+len(pickup)
		}

		for i, c := range cards {
			if c == destinationCard {
				destinationCardIndex = i
				found = true
				break
			}
		}
	}

	output := append(cards[:destinationCardIndex+1], append(pickup, cards[destinationCardIndex+1:]...)...)
	if destinationCardIndex < currentIndex {
		output = append(output[3:], output[:3]...)
	}

	return output, (currentIndex+1) % len(output)
	
}
