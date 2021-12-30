package main

import (
	"advent-of-code/common/aoc"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input := aoc.GetInputRows(2020, 19)

	rules := make(map[int]string)
	messages := make([]string, 0)

	parseRules := true
	for _, row := range input {
		if row == "" {
			parseRules = false
			continue
		}

		if parseRules {
			parts := strings.Split(row, ":")
			index, _ := strconv.Atoi(parts[0])
			rules[index] = strings.TrimSpace(parts[1])
		} else {
			messages = append(messages, row)
		}
	}

	var parseRule func(rule string) string
	parseRule = func(rule string) string {
		parsed := ""
		parts := strings.Split(rule, " ")
		isOr := false
		for _, part := range parts {
			if part == "|" {
				isOr = true
				parsed = "(" + parsed
				parsed += "|"
			} else if strings.HasPrefix(part, "\"") {
				return strings.Trim(part, "\"")
			} else {
				d, _ := strconv.Atoi(part)
				parsed += parseRule(rules[d])
			}
		}
		if isOr {
			parsed += ")"
		}
		return parsed
	}

	pattern := parseRule(rules[0])
	r := regexp.MustCompile(fmt.Sprintf("^%s$", pattern))

	valid := 0
	for _, message := range messages {
		if r.MatchString(message) {
			valid++
		}
	}

	fmt.Println("Puzzle 1:", valid)
	fmt.Println("")
	fmt.Println(parseRule(rules[42]))
	fmt.Println("")
	fmt.Println(parseRule(rules[31]))
}
