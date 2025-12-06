package main

import (
	"advent-of-code/common/aoc"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type HandType int

type Hand struct {
	Cards string
	Type  HandType
	Bid   int
}

const (
	Unknown HandType = iota
	HighCard
	OnePair
	TwoPairs
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

const JokerRune rune = 'J'

var CardScores map[rune]int = map[rune]int{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

func main() {
	task := aoc.Init(2023, 7)
	input := task.AsStringSlice()
	hands1 := make([]Hand, len(input))
	hands2 := make([]Hand, len(input))

	for i, row := range input {
		parts := strings.Split(row, " ")

		bid, _ := strconv.Atoi(parts[1])

		hands1[i] = Hand{
			Cards: parts[0],
			Type:  DetermineHandType(parts[0], false),
			Bid:   bid,
		}

		hands2[i] = Hand{
			Cards: parts[0],
			Type:  DetermineHandType(parts[0], true),
			Bid:   bid,
		}
	}

	OrderHands(hands1, false)
	OrderHands(hands2, true)

	sum1, sum2 := 0, 0
	for i, hand := range hands1 {
		sum1 += hand.Bid * (i + 1)
	}

	for i, hand := range hands2 {
		sum2 += hand.Bid * (i + 1)
	}

	fmt.Println("Puzzle 1:", sum1)
	fmt.Println("Puzzle 2:", sum2)
}

func DetermineHandType(cards string, useJokers bool) HandType {
	cardAmounts := make(map[rune]int)
	for _, card := range cards {
		cardAmounts[card]++
	}

	orderedCardAmounts := make([]int, len(cardAmounts))
	i := 0
	for _, v := range cardAmounts {
		orderedCardAmounts[i] = v
		i++
	}
	sort.Sort(sort.Reverse(sort.IntSlice(orderedCardAmounts)))

	if jokerCount := cardAmounts[JokerRune]; useJokers && jokerCount > 0 {
		switch len(cardAmounts) {
		case 1, 2:
			return FiveOfAKind
		case 3:
			switch jokerCount {
			case 1:
				if orderedCardAmounts[0] == 3 {
					return FourOfAKind
				}
				return FullHouse
			case 2:
				return FourOfAKind
			case 3:
				return FourOfAKind
			}
		case 4:
			return ThreeOfAKind
		case 5:
			return OnePair
		default:
			return Unknown
		}
	}

	switch len(cardAmounts) {
	case 1:
		return FiveOfAKind
	case 2:
		if orderedCardAmounts[0] == 3 && orderedCardAmounts[1] == 2 {
			return FullHouse
		}
		return FourOfAKind
	case 3:
		if orderedCardAmounts[0] == 3 {
			return ThreeOfAKind
		}
		return TwoPairs
	case 4:
		return OnePair
	case 5:
		return HighCard
	default:
		return Unknown
	}
}

func OrderHands(hands []Hand, useJokers bool) {
	scores := CardScores
	if useJokers {
		scores[JokerRune] = 1
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].Type == hands[j].Type {
			for k := range hands[i].Cards {
				if CardScores[rune(hands[i].Cards[k])] == CardScores[rune(hands[j].Cards[k])] {
					continue
				}
				return CardScores[rune(hands[i].Cards[k])] < CardScores[rune(hands[j].Cards[k])]
			}
		}
		return hands[i].Type < hands[j].Type
	})

}
