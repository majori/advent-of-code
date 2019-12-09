package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("./test.txt")
	rawCodes := strings.Split(string(input), ",")

	codes := make([]int, len(rawCodes))
	for i, elem := range rawCodes {
		value, _ := strconv.Atoi(elem)
		codes[i] = value
	}

	amplify := func(phases []int, init int) (int, bool) {
		output := init
		var open bool

		for i := 0; i < len(phases); i++ {
			c := make(chan int)
			go compute(codes, c)
			c <- phases[i]
			c <- output

			for x := range c {
				output = x
			}
			// output, open = <-c

			if !open {
				return 0, true
			}
		}

		return output, false
	}

	max := 0
	// var bestPhase *[]int
	// for _, phases := range permutation([]int{0, 1, 2, 3, 4}) {
	// 	output, _ := amplify(phases, 0)
	// 	if output > max {
	// 		max = output
	// 		bestPhase = &phases
	// 	}
	// }

	// fmt.Printf("Part 1: %d (Phases %v)\n", max, *bestPhase)

	max = 0
	// permutations := permutation([]int{5, 6, 7, 8, 9})
	for i, phases := range [][]int{[]int{9, 8, 7, 6, 5}} {
		fmt.Printf("Phase %d\n", i)
		complete := false
		output, timeout := 0, 0
		for !complete {
			if timeout > 100000 {
				fmt.Println("TIMEOUT")
				return
			}

			// fmt.Println("ROUND")
			output, complete = amplify(phases, output)
			timeout++
		}

		if output > max {
			max = output
		}
	}

	fmt.Printf("Part 2: %d", max)
}

func compute(codes []int, ch chan int) {
	index := 0
	for true {
		code := codes[index]
		digits := []int{
			code % 10,
			(code / 10) % 10,
			(code / 100) % 10,
			(code / 1000) % 10,
			(code / 10000) % 10,
		}

		if digits[0] == 9 && digits[1] == 9 {
			close(ch)
			return
		}

		var param1Index, param2Index int
		if digits[2] == 0 {
			param1Index = codes[index+1]
		} else {
			param1Index = index + 1
		}

		if digits[3] == 0 {
			param2Index = codes[index+2]
		} else {
			param2Index = index + 2
		}

		// if digits[4] == 0 {
		// 	param3Index = codes[index+3]
		// } else {
		// 	param3Index = index + 3
		// }

		switch digits[0] {
		case 1:
			a := codes[param1Index]
			b := codes[param2Index]

			codes[codes[index+3]] = a + b
			index += 4
			break
		case 2:
			a := codes[param1Index]
			b := codes[param2Index]

			codes[codes[index+3]] = a * b
			index += 4
			break
		case 3:
			codes[param1Index] = <-ch
			index += 2
			break
		case 4:
			ch <- codes[param1Index]
			index += 2
			break
		case 5:
			if codes[param1Index] != 0 {
				index = codes[param2Index]
			} else {
				index += 3
			}
			break
		case 6:
			if codes[param1Index] == 0 {
				index = codes[param2Index]
			} else {
				index += 3
			}
			break
		case 7:
			if codes[param1Index] < codes[param2Index] {
				codes[codes[index+3]] = 1
			} else {
				codes[codes[index+3]] = 0
			}
			index += 4
			break
		case 8:
			if codes[param1Index] == codes[param2Index] {
				codes[codes[index+3]] = 1
			} else {
				codes[codes[index+3]] = 0
			}
			index += 4
			break
		}
	}
}

func permutation(xs []int) (permuts [][]int) {
	var rc func([]int, int)
	rc = func(a []int, k int) {
		if k == len(a) {
			permuts = append(permuts, append([]int{}, a...))
		} else {
			for i := k; i < len(xs); i++ {
				a[k], a[i] = a[i], a[k]
				rc(a, k+1)
				a[k], a[i] = a[i], a[k]
			}
		}
	}
	rc(xs, 0)

	return permuts
}
