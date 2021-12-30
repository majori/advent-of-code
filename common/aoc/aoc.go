package aoc

import (
	"strconv"
	"strings"
)

type AoC struct {
	day int
	year int
	input string
	example string
}

func Init(year, day int) AoC {
	return AoC{
		year: year,
		day: day,
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

func (a AoC) Submit1(answer interface{}) {
	submit(a.year, a.day, 1, answer)
}

func (a AoC) Submit2(answer interface{}) {
	submit(a.year, a.day, 2, answer)
}