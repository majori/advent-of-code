package main

import (
	"advent-of-code/common/aoc"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	Items                    []int
	Operation                func(int, int) int
	OpB                      int
	TestModulus              int
	ThrowToMonkeyIfTestTrue  int
	ThrowToMonkeyIfTestFalse int
	ItemsInspected           int
}

func main() {
	task := aoc.Init(2022, 11)
	input := task.AsString()

	monkeys := make([]Monkey, 0)
	reg := regexp.MustCompile(`Monkey (\d):\n. Starting items: (.*)\n. Operation: new = (.+) (.) (.+)\n  Test: divisible by (\d+)\n.   If true: throw to monkey (\d+)\n.   If false: throw to monkey (\d+)`)
	for _, raw := range reg.FindAllStringSubmatch(input, -1) {
		monkey := Monkey{}

		itemsRaw := strings.Split(raw[2], ", ")
		monkey.Items = make([]int, len(itemsRaw))
		for i := range itemsRaw {
			monkey.Items[i], _ = strconv.Atoi(itemsRaw[i])
		}

		if raw[4] == "*" {
			monkey.Operation = multiply
		} else {
			monkey.Operation = sum
		}

		if raw[5] != "old" {
			monkey.OpB, _ = strconv.Atoi(raw[5])
		}

		monkey.TestModulus, _ = strconv.Atoi(raw[6])
		monkey.ThrowToMonkeyIfTestTrue, _ = strconv.Atoi(raw[7])
		monkey.ThrowToMonkeyIfTestFalse, _ = strconv.Atoi(raw[8])

		monkeys = append(monkeys, monkey)
	}

	for i := 0; i < 20; i++ {
		Round(monkeys)
	}

	inspected := make([]int, len(monkeys))
	for _, monkey := range monkeys {
		inspected = append(inspected, monkey.ItemsInspected)
	}
	sort.Ints(inspected)
	fmt.Println(inspected)
	fmt.Println("Puzzle 1:", inspected[len(inspected)-1]*inspected[len(inspected)-2])
}

func Turn(monkeys []Monkey, id int) {
	monkey := &monkeys[id]
	for _, item := range monkey.Items {
		var worry int
		if monkey.OpB != 0 {
			worry = monkey.Operation(item, monkey.OpB)
		} else {
			worry = monkey.Operation(item, item)
		}

		x := (worry / 3)
		var targetMonkey *Monkey
		if (x % monkey.TestModulus) == 0 {
			targetMonkey = &monkeys[monkey.ThrowToMonkeyIfTestTrue]
		} else {
			targetMonkey = &monkeys[monkey.ThrowToMonkeyIfTestFalse]
		}
		targetMonkey.Items = append(targetMonkey.Items, x)
		monkey.ItemsInspected++
	}
	monkey.Items = make([]int, 0)
}

func Round(monkeys []Monkey) {
	for i := range monkeys {
		Turn(monkeys, i)
	}
}

type Operator func(int, int) int

func multiply(a, b int) int {
	return a * b
}

func sum(a, b int) int {
	return a + b
}
