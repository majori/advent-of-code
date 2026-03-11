package aoc

type Grid struct {
	grid [][]rune
	x    int
	y    int
	it   bool
}

func (g Grid) String() (s string) {
	for _, row := range g.grid {
		s += string(row) + "\n"
	}
	return
}

func (g Grid) validate(x, y int) bool {
	if x < 0 || y < 0 || x >= len(g.grid[0]) || y >= len(g.grid) {
		return false
	}

	return true
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

func (g *Grid) get(x, y int) (rune, bool) {
	if !g.validate(x, y) {
		return rune(0), false
	}
	return g.grid[y][x], true
}

func (g *Grid) Get() rune {
	r, _ := g.get(g.x, g.y)
	return r
}
func (g *Grid) GetLeft() (rune, bool) {
	return g.get(g.x-1, g.y)
}
func (g *Grid) GetRight() (rune, bool) {
	return g.get(g.x+1, g.y)
}
func (g *Grid) GetTop() (rune, bool) {
	return g.get(g.x, g.y-1)
}
func (g *Grid) GetDown() (rune, bool) {
	return g.get(g.x, g.y+1)
}

func (g *Grid) set(x, y int, r rune) bool {
	if _, ok := g.get(x, y); !ok {
		return false
	}
	g.grid[y][x] = r
	return true
}

func (g *Grid) Set(r rune) {
	g.set(g.x, g.y, r)
}

func (g *Grid) SetLeft(r rune) bool {
	return g.set(g.x-1, g.y, r)
}

func (g *Grid) SetRight(r rune) bool {
	return g.set(g.x+1, g.y, r)
}

func (g *Grid) SetUp(r rune) bool {
	return g.set(g.x, g.y-1, r)
}

func (g *Grid) SetDown(r rune) bool {
	return g.set(g.x, g.y+1, r)
}

func (g *Grid) Cursor() (int, int) {
	return g.x, g.y
}

func (g *Grid) SetCursor(x, y int) bool {
	if !g.validate(x, y) {
		return false
	}

	g.x = x
	g.y = y
	return true
}

func (g *Grid) Move(x, y int) bool {
	if !g.validate(g.x+x, g.y+y) {
		return false
	}

	g.x += x
	g.y += y
	return true
}

func (g *Grid) MoveLeft() bool {
	return g.Move(-1, 0)
}
func (g *Grid) MoveRight() bool {
	return g.Move(1, 0)
}
func (g *Grid) MoveUp() bool {
	return g.Move(0, -1)
}
func (g *Grid) MoveDown() bool {
	return g.Move(0, 1)
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

func (g *Grid) IsInLastRow() bool {
	return g.y == len(g.grid)-1
}
