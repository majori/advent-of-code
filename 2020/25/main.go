package main

import (
	"advent-of-code/common/aoc"
	"fmt"
)

func main() {
	task := aoc.Init(2020, 25)
	input := task.AsIntSlice()
	fmt.Println(input)

	publicCard := input[0]
	publicDoor := input[1]
	loopsCard := crack(publicCard, 7)
	loopsDoor := crack(publicDoor, 7)

	task.Submit1(loop(publicDoor, loopsCard))
	task.Submit2(loop(publicCard, loopsDoor))
}

func loop(s int, n int) (v int) {
	v = 1
	for i := 0; i < n; i++ {
		v *= s
		v = v % 20201227
	}
	return
}

func crack(k int, s int) (l int) {
	v := 1
	for {
		l++
		v *= s
		v = v % 20201227

		if v == k {
			return
		}
	}
}
