package main

import (
	"advent-of-code/common/aoc"
	"fmt"
	"strconv"
	"strings"
)

type Cell struct {
	Value int
	Checked bool
}
type Board [5][5]Cell
type BingoBoard struct {
	Board Board
	SearchIndex map[int]*bool
}


func main() {
	task := aoc.Init(2021, 4)
	input := task.AsString()

	lines := strings.Split(input,"\n")

	rolls := make([]int, 0)
	for _, x := range strings.Split(lines[0],",") {
		roll, _ := strconv.Atoi(x)
		rolls = append(rolls, roll)
	}

	boards := make([]*BingoBoard, 0)
	
	for i := 2; i < len(lines); i+=6 {
		bingo := BingoBoard{}
		bingo.SearchIndex = make(map[int]*bool)
		for y := 0; y < 6; y++ {
			rawRow := strings.Split(lines[i+y], " ")
			x := 0
			for _, cell := range rawRow {
				if cell == "" {
					continue
				}
				a, _ := strconv.Atoi(cell)
				bingo.Board[y][x] = Cell{Value: a, Checked: false}
				bingo.SearchIndex[a] = &(bingo.Board[y][x].Checked)
				x++
			}
		}
		boards = append(boards, &bingo)
	}

	winningBoard, lastRoll := calculateWinningBoard(rolls, boards)
	fmt.Println("Puzzle 1: ", lastRoll*calculateUncheckedSum(winningBoard))

	resetBingoBoards(boards)
	
	loserBoard, lastRoll := calculateLoserBoard(rolls, boards)
	fmt.Println("Puzzle 2: ", lastRoll*calculateUncheckedSum(loserBoard))

}

func checkBingo(b *BingoBoard) bool {
	for y := 0; y < 5; y++ {
		bingo := true
		for x := 0; x < 5; x++ {
			if !b.Board[y][x].Checked {
				bingo = false
			}
		}
		if bingo {
			return true
		}
	}
	for x := 0; x < 5; x++ {
		bingo := true
		for y := 0; y < 5; y++ {
			if !b.Board[y][x].Checked {
				bingo = false
			}
		}
		if bingo {
			return true
		}
	}

	return false
}

func resetBingoBoards(boards []*BingoBoard) {
	for _, b := range boards {
		for _, x := range b.SearchIndex {
			*x = false
		}
	}
}

func calculateWinningBoard(rolls []int, boards []*BingoBoard) (*BingoBoard, int) {
	var winningBoard *BingoBoard
	var lastRoll int

	for _, r := range rolls {
		if winningBoard != nil {
			break
		}
		for _, b := range boards {
			if val, ok := b.SearchIndex[r]; ok {
				*val = true
			}
			if checkBingo(b) {
				winningBoard = b
				lastRoll = r
				break
			}
		}
	}
	return winningBoard, lastRoll
}

func calculateUncheckedSum(b *BingoBoard) int {
	sum := 0
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if !b.Board[y][x].Checked {
				sum += b.Board[y][x].Value
			}
		}
	}

	return sum
}

func calculateLoserBoard(rolls []int, boards []*BingoBoard) (*BingoBoard, int) {
	var loserBoard *BingoBoard

	for _, r := range rolls {
		temp := make([]*BingoBoard, 0, len(boards))
		if loserBoard != nil {
			break
		}
		for _, b := range boards {
			if val, ok := b.SearchIndex[r]; ok {
				*val = true
			}
			if len(boards) == 1 && checkBingo(b) {
				return boards[0], r

			} else if !checkBingo(b) {
				temp = append(temp, b)
			}
		}
		boards = temp
	}

	return nil, 0
}