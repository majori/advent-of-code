package main

import (
	"advent-of-code/common/aoc"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	input := aoc.GetInput(2020, 21).AsStringSlice()
	reg := regexp.MustCompile(`^(.*) \(contains (.*)\)$`)
	allIngredients := make(map[string]int)
	allAllergens := make(map[string][]string)

	for _, row := range input {
		match := reg.FindStringSubmatch(row)
		ingredients := strings.Split(match[1], " ")
		allergens := strings.Split(match[2], ", ")

		for _, ingredient := range ingredients {
			allIngredients[ingredient]++
		}

		for _, a := range allergens {
			if _, ok := allAllergens[a]; !ok {
				allAllergens[a] = ingredients
			} else {
				intersection := make([]string, 0)
				for _, i1 := range ingredients {
					for _, i2 := range allAllergens[a] {
						if i1 == i2 {
							intersection = append(intersection, i1)
						}
					}
				}
				allAllergens[a] = intersection
			}
		}
	}

	n := 0

	for ingredient, count := range allIngredients {
		found := false
		for _, ingredientsWithAllergen := range allAllergens {
			for _, ingredientWithAllergen := range ingredientsWithAllergen {
				if ingredient == ingredientWithAllergen {
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		if !found {
			n += count
		}
	}

	fmt.Println("Puzzle 1:", n)
}
