package main

import (
	"advent-of-code/common/aoc"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	task := aoc.Init(2021, 6)
	input := task.AsString()
	split := strings.Split(strings.TrimSpace(input), ",")

	fishes := make([]int8, len(split))
	for i, l := range split {
		v, err := strconv.ParseInt(l, 10, 8)
		if err != nil {
			log.Fatal(err, i)
		}
		fishes[i] = int8(v)
	}

	
	for n := 0; n < 80; n++ {
		fishes = step(fishes)
	}
	
	fmt.Println("Puzzle 1:", len(fishes))
	
	fishes2 := [8]int64{}

	for i, l := range split {
		v, err := strconv.Atoi(l)
		if err != nil {
			log.Fatal(err, i)
		}
		fishes2[v]++
	}

	for n := 0; n < 256; n++ {
		temp := [8]int64{}
		for i := range fishes2 {
			if i == 0 {
				temp[7] = fishes2[0]
				temp[5] = fishes2[0]
			} else {
				temp[i-1] += fishes2[i]
			}
		}
		fishes2 = temp
	}

	fmt.Println("Puzzle 2:", fishes2)
}

func step(fishes []int8) []int8 {
	temp := make([]int8, len(fishes))
	copy(temp, fishes)
	for i, f := range fishes {
		if f == 0 {
			temp = append(temp, 8)
			temp[i] = 6
		} else {
			temp[i]--
		}
	}
		
	return temp
}