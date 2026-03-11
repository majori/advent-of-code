package aoc

import (
	"strconv"
	"strings"
)

type AoC struct {
	day     int
	year    int
	input   string
	example string
}

func Init(year, day int) AoC {
	return AoC{
		year:  year,
		day:   day,
		input: getInput(year, day),
	}
}

func (a *AoC) SetExample(example string) {
	a.example = example
}

func (a AoC) AsString() string {
	if a.example != "" {
		return a.example
	}
	return a.input
}

func (a AoC) AsStringSlice() []string {
	input := strings.TrimSpace(a.AsString())
	return strings.Split(input, "\n")
}

func (a AoC) AsIntSlice() []int {
	rows := a.AsStringSlice()
	intRows := make([]int, len(rows))
	for i, s := range rows {
		x, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		intRows[i] = x
	}
	return intRows
}

func (a AoC) AsRuneGrid() Grid {
	rows := a.AsStringSlice()
	grid := make([][]rune, len(rows))
	for i, s := range rows {
		grid[i] = []rune(s)
	}
	return Grid{
		x:    0,
		y:    0,
		grid: grid,
		it:   false,
	}
}

func (a AoC) Submit1(answer any) {
	submit(a.year, a.day, 1, answer)
}

func (a AoC) Submit2(answer any) {
	submit(a.year, a.day, 2, answer)
}
