package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	dat, _ := ioutil.ReadFile("./input.txt")
	lines := strings.Split(string(dat), "\n")

	masses := []int{}
	for _, line := range lines {
		mass, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		masses = append(masses, mass)
	}

	result := 0
	for _, mass := range masses {
		result += calculate(mass)
	}

	fmt.Println(result)

	result = 0
	for _, mass := range masses {
		x := mass
		for x > 0 {
			x = calculate(x)
			if x > 0 {
				result += x
			}
		}
	}

	fmt.Println(result)
}

func calculate(mass int) int {
	return int(math.Floor(float64(mass)/3.0) - 2)
}
