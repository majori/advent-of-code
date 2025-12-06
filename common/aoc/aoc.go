package aoc

import (
	"errors"
	"strconv"
	"strings"
)

type AoC struct {
	day     int
	year    int
	input   string
	example string
}

type Grid struct {
	grid [][]rune
	x    int
	y    int
	it   bool
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

func (g *Grid) Iterate() bool {
	if !g.it {
		g.SetCursor(0, 0)
		g.it = true
		return true
	}
	if g.x == len(g.grid[g.y])-1 {
		if g.y == len(g.grid)-1 {
			g.SetCursor(0, 0)
			g.it = false
			return false
		}
		g.x = 0
		g.y++
	} else {
		g.x++
	}
	return true
}

func (g *Grid) Get() rune {
	return g.grid[g.y][g.x]
}

func (g *Grid) Set(r rune) {
	g.grid[g.y][g.x] = r
}

func (g *Grid) SetCursor(x, y int) error {
	if x < 0 || y < 0 || x >= len(g.grid[0]) || y >= len(g.grid) {
		return errors.New("invalid cursor")
	}

	g.x = x
	g.y = y
	return nil
}

func (g *Grid) GetNeighbours() []rune {
	n := make([]rune, 0)
	if g.x > 0 {
		n = append(n, g.grid[g.y][g.x-1])
	}

	if g.y > 0 {
		n = append(n, g.grid[g.y-1][g.x])

		if g.x > 0 {
			n = append(n, g.grid[g.y-1][g.x-1])
		}
		if g.x < len(g.grid[g.y])-1 {
			n = append(n, g.grid[g.y-1][g.x+1])
		}
	}

	if g.x < len(g.grid[g.y])-1 {
		n = append(n, g.grid[g.y][g.x+1])
	}

	if g.y < len(g.grid)-1 {
		n = append(n, g.grid[g.y+1][g.x])

		if g.x > 0 {
			n = append(n, g.grid[g.y+1][g.x-1])
		}
		if g.x < len(g.grid[g.y])-1 {
			n = append(n, g.grid[g.y+1][g.x+1])
		}
	}

	return n
}

func (a AoC) Submit1(answer any) {
	submit(a.year, a.day, 1, answer)
}

func (a AoC) Submit2(answer any) {
	submit(a.year, a.day, 2, answer)
}
