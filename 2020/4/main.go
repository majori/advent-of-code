package main

import (
	"advent-of-code/common/aoc"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var fields []string = []string{
	"byr", // Birth Year
	"iyr", // Issue Year
	"eyr", // Expiration Year
	"hgt", // Height
	"hcl", // Hair Color
	"ecl", // Eye Color
	"pid", // Passport ID
	"cid", // Country ID
}

func main() {
	in := aoc.GetInput(2020, 4)
	inList := strings.Split(in, "\n\n")

	reg := regexp.MustCompile(`(\w+):([^\s]+)`)

	passports := make([]map[string]string, 0, len(inList))

	for _, raw := range inList {
		entries := reg.FindAllStringSubmatch(raw, -1)

		passport := make(map[string]string)
		for _, entry := range entries {
			passport[entry[1]] = entry[2]
		}

		passports = append(passports, passport)
	}

	var valid int

passportsLoop:
	for i, passport := range passports {

		for _, field := range fields {
			if _, ok := passport[field]; !ok && field != "cid" {
				passports[i] = nil
				continue passportsLoop
			}
		}

		valid++
	}

	fmt.Println("Puzzle 1: ", valid)

	valid = 0

validate:
	for _, passport := range passports {
		if passport == nil {
			continue
		}

		for key, value := range passport {
			switch key {
			case "byr":
				year, err := strconv.Atoi(value)
				if !(1920 <= year && year <= 2002 && err == nil) {
					continue validate
				}

			case "iyr":
				year, err := strconv.Atoi(value)
				if !(2010 <= year && year <= 2020 && err == nil) {
					continue validate
				}

			case "eyr":
				year, err := strconv.Atoi(value)
				if !(2020 <= year && year <= 2030 && err == nil) {
					continue validate
				}

			case "hgt":
				reg := regexp.MustCompile(`^(\d+)(in|cm)$`)
				matches := reg.FindStringSubmatch(value)
				if len(matches) != 3 {
					continue validate
				}

				unit := matches[2]
				h, _ := strconv.Atoi(matches[1])
				if unit == "cm" && !(150 <= h && h <= 193) {
					continue validate
				} else if unit == "in" && !(59 <= h && h <= 76) {
					continue validate
				}

			case "hcl":
				reg := regexp.MustCompile(`^#(\w|\d){6}$`)
				if !reg.MatchString(value) {
					continue validate
				}

			case "ecl":
				allowed := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
				valid := false
				for _, elem := range allowed {
					if value == elem {
						valid = true
						break
					}
				}
				if !valid {
					continue validate
				}

			case "pid":
				reg := regexp.MustCompile(`^\d{9}$`)
				if !reg.MatchString(value) {
					continue validate
				}
			}
		}
		valid++
	}
	fmt.Println("Puzzle 2: ", valid)
}
