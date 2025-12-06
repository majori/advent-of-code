package main

import (
	"advent-of-code/common/aoc"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	task := aoc.Init(2024, 2)
	input := task.AsStringSlice()

	reports := make([][]int, len(input))

	for i, row := range input {
		for _, elem := range strings.Split(row, " ") {
			y, _ := strconv.Atoi(elem)
			reports[i] = append(reports[i], y)
		}
	}

	safeCount := 0
	fixedSafeCount := 0

	for _, report := range reports {
		isSafe := validator(report)
		if isSafe {
			safeCount++
			continue
		}

		for i := range report {
			fixedReport := make([]int, len(report))
			copy(fixedReport, report)
			fixedReport = append(fixedReport[:i], fixedReport[i+1:]...)
			isSafe = validator(fixedReport)
			if isSafe {
				fixedSafeCount++
				break
			}
		}
	}

	fmt.Println("Puzzle 1: ", safeCount)
	fmt.Println("Puzzle 2: ", safeCount+fixedSafeCount)
}

func validator(report []int) bool {
	increasing := false
	if report[0] < report[1] {
		increasing = true
	}

	for i := 0; i < len(report)-1; i++ {
		if increasing && report[i] >= report[i+1] {
			return false
		}

		if !increasing && report[i] <= report[i+1] {
			return false
		}

		diff := report[i] - report[i+1]
		if diff < 0 {
			diff = -diff
		}

		if diff < 1 || diff > 3 {
			return false
		}
	}

	return true
}
