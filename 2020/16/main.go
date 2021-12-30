package main

import (
	"advent-of-code/common/aoc"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input := aoc.GetInput(2020, 16)

	identifiers := map[string]string{
		"depLoc":  "departure location",
		"depSta":  "departure station",
		"depPlat": "departure platform",
		"depTra":  "departure track",
		"depDat":  "departure date",
		"depTim":  "departure time",
		"arrLoc":  "arrival location",
		"arrSta":  "arrival station",
		"arrPla":  "arrival platform",
		"arrTra":  "arrival track",
		"cla":     "class",
		"dur":     "duration",
		"pri":     "price",
		"rou":     "route",
		"row":     "row",
		"sea":     "seat",
		"tra":     "train",
		"typ":     "type",
		"wag":     "wagon",
		"zon":     "zone",
	}

	fieldRanges := make(map[string][][2]int)
	for key, id := range identifiers {
		regex := regexp.MustCompile(fmt.Sprintf(`%s: (\d+)-(\d+) or (\d+)-(\d+)`, id))
		match := regex.FindAllStringSubmatch(input, -1)
		for i := 1; i < len(match[0]); i += 2 {
			lower, _ := strconv.Atoi(match[0][i])
			upper, _ := strconv.Atoi(match[0][i+1])
			fieldRanges[key] = append(fieldRanges[key], [2]int{lower, upper})
		}
	}

	nearbyTickets := make([][20]int, 0)

	nearbyRegex := regexp.MustCompile(`nearby tickets:\n(.*\n?)+`)
	nearbyBlock := nearbyRegex.FindString(input)
	rowRegex := regexp.MustCompile(`((?:\d+,?){20})`)
	for _, row := range rowRegex.FindAllString(nearbyBlock, -1) {
		ticket := [20]int{}
		for i, s := range strings.Split(row, ",") {
			ticket[i], _ = strconv.Atoi(s)
		}
		nearbyTickets = append(nearbyTickets, ticket)
	}

	errorRate := 0
	invalid := make([]int, 0)

	for i, ticket := range nearbyTickets {
	validation:
		for _, field := range ticket {
			for _, fieldRange := range fieldRanges {
				for _, r := range fieldRange {
					if field >= r[0] && field <= r[1] {
						continue validation
					}
				}
			}
			invalid = append(invalid, i)
			errorRate += field
			break
		}
	}

	fmt.Println("Puzzle 1:", errorRate)

	for i := len(invalid) - 1; i >= 0; i-- {
		nearbyTickets = append(nearbyTickets[:invalid[i]], nearbyTickets[invalid[i]+1:]...)
	}

	myRegex := regexp.MustCompile(`your ticket:\n((?:\d+,?){20})\n`)
	myTicket := [20]int{}
	for i, s := range strings.Split(myRegex.FindStringSubmatch(input)[1], ",") {
		myTicket[i], _ = strconv.Atoi(s)
	}

	// fieldNames := make(map[int]string)

	for fieldIndex := 0; fieldIndex < 20; fieldIndex++ {
		check := make(map[string]int)
		for _, ticket := range nearbyTickets {
			for key, fieldRange := range fieldRanges {
				if _, ok := check[key]; !ok {
					check[key] = -190
				}
				ok := false
				for _, r := range fieldRange {
					if ticket[fieldIndex] >= r[0] && ticket[fieldIndex] <= r[1] {
						ok = true
					}
				}
				if ok {
					check[key]++
				}
			}
		}

		count := 0
		for _, v := range check {
			if v == 0 {
				count++
			}
		}
		// fmt.Println(count, fieldIndex, check)
	}
	fmt.Println("Puzzle 2:", myTicket[2]*myTicket[4]*myTicket[7]*myTicket[11]*myTicket[14]*myTicket[16])
}
