package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("./input.txt")
	rawCodes := strings.Split(string(input), ",")

	codes := make([]int, len(rawCodes))
	for i, elem := range rawCodes {
		value, _ := strconv.Atoi(elem)
		codes[i] = value
	}

	// fmt.Println(codes)
	index := 0
	for codes[index] != 99999 {
		code := codes[index]
		digits := []int{
			code % 10,
			(code / 10) % 10,
			(code / 100) % 10,
			(code / 1000) % 10,
			(code / 10000) % 10,
		}

		if digits[0] == 9 && digits[1] == 9 {
			break
		}

		input := 5

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
			codes[param1Index] = input
			index += 2
			break
		case 4:
			fmt.Println(codes[param1Index])
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

		// instruction := code[len(code)-1]
	}
}
