package main

import (
	"advent-of-code/common/aoc"
	"fmt"
	"sync"
)

func main() {
	input := aoc.GetInputRows(2020, 11)

	layout := make([][]rune, 0, len(input))
	for _, row := range input {
		layout = append(layout, []rune(row))
	}

	layout = runUntilMatched(layout, checkSeat1)

	fmt.Println("Puzzle 1:", countOccupied(layout))

	for i := range input {
		layout[i] = []rune(input[i])
	}

	layout = runUntilMatched(layout, checkSeat2)

	fmt.Println("Puzzle 2:", countOccupied(layout))
}

func runUntilMatched(layout [][]rune, checkFunc func(int, int, [][]rune, *rune, *sync.WaitGroup)) [][]rune {
	for {
		next := runRound(layout, checkFunc)
		match := compareSlices(layout, next)
		layout = next

		if match {
			break
		}
	}

	return layout
}

func runRound(layout [][]rune, checkFunc func(int, int, [][]rune, *rune, *sync.WaitGroup)) [][]rune {
	next := make([][]rune, len(layout))
	for i := range next {
		next[i] = make([]rune, len(layout[i]))
		copy(next[i], layout[i])
	}
	var wg sync.WaitGroup
	for i := range layout {
		for j := range layout[i] {
			wg.Add(1)
			go checkFunc(i, j, layout, &next[i][j], &wg)
		}
	}
	wg.Wait()

	return next
}

func checkSeat1(y, x int, layout [][]rune, write *rune, wg *sync.WaitGroup) {
	defer wg.Done()
	seat := layout[y][x]

	if seat == '.' {
		return
	}

	rows := []int{y - 1, y, y + 1}
	if y == 0 {
		rows = rows[1:]
	}
	if y == len(layout)-1 {
		rows = rows[:2]
	}

	columns := []int{x - 1, x, x + 1}
	if x == 0 {
		columns = columns[1:]
	}
	if x == len(layout[y])-1 {
		columns = columns[:2]
	}

	adjacent := make(map[rune]int)
	for _, row := range rows {
		for _, column := range columns {
			if row == y && column == x {
				continue
			}

			current := layout[row][column]
			adjacent[current]++
		}
	}

	switch seat {
	case 'L':
		if adjacent['#'] == 0 {
			*write = '#'
		}

	case '#':
		if adjacent['#'] >= 4 {
			*write = 'L'
		}
	}
}

func checkSeat2(y, x int, layout [][]rune, write *rune, wg *sync.WaitGroup) {
	defer wg.Done()
	seat := layout[y][x]
	adjacent := make(map[rune]int)

	if seat == '.' {
		return
	}

	for i := x - 1; i >= 0; i-- {
		if layout[y][i] != '.' {
			adjacent[layout[y][i]]++
			break
		}
	}
	for i := x + 1; i < len(layout[y]); i++ {
		if layout[y][i] != '.' {
			adjacent[layout[y][i]]++
			break
		}
	}
	for i := y - 1; i >= 0; i-- {
		if layout[i][x] != '.' {
			adjacent[layout[i][x]]++
			break
		}
	}
	for i := y + 1; i < len(layout); i++ {
		if layout[i][x] != '.' {
			adjacent[layout[i][x]]++
			break
		}
	}

	i := 1
	for {
		if x+i >= len(layout[y]) || y+i >= len(layout) {
			break
		}

		if layout[y+i][x+i] != '.' {
			adjacent[layout[y+i][x+i]]++
			break
		}
		i++
	}

	i = 1
	for {
		if x-i < 0 || y-i < 0 {
			break
		}

		if layout[y-i][x-i] != '.' {
			adjacent[layout[y-i][x-i]]++
			break
		}
		i++
	}

	i = 1
	for {
		if x+i >= len(layout[y]) || y-i < 0 {
			break
		}

		if layout[y-i][x+i] != '.' {
			adjacent[layout[y-i][x+i]]++
			break
		}
		i++
	}

	i = 1
	for {
		if x-i < 0 || y+i >= len(layout) {
			break
		}

		if layout[y+i][x-i] != '.' {
			adjacent[layout[y+i][x-i]]++
			break
		}
		i++
	}

	switch seat {
	case 'L':
		if adjacent['#'] == 0 {
			*write = '#'
		}

	case '#':
		if adjacent['#'] >= 5 {
			*write = 'L'
		}
	}
}

func compareSlices(a, b [][]rune) bool {
	match := true
	for i := range a {
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				match = false
				break
			}
		}
		if !match {
			break
		}
	}

	return match
}

func countOccupied(layout [][]rune) (counter int) {
	for i := range layout {
		for j := range layout[i] {
			if layout[i][j] == '#' {
				counter++
			}
		}
	}
	return
}
