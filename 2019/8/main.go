package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("./input.txt")
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanRunes)

	const width int = 25
	const height int = 6

	i := 0
	minZeros := 100000
	minZerosLayer := make([]int, width*height)
	layer := make([]int, width*height)
	decoded := make([]int, 0, width*height)

	for s.Scan() {
		// End of row
		if i == (width * height) {
			zeros := 0
			for _, pixel := range layer {
				if pixel == 0 {
					zeros++
				}
			}

			if zeros < minZeros {
				minZeros = zeros
				copy(minZerosLayer, layer)
			}

			// Init for new row
			i = 0
		}

		value, _ := strconv.Atoi(s.Text())
		layer[i] = value

		if len(decoded) < len(layer) {
			decoded = append(decoded, layer[i])
		} else if decoded[i] == 2 {
			decoded[i] = value
		}

		i++
	}

	ones, twos := 0, 0
	for _, pixel := range minZerosLayer {
		if pixel == 1 {
			ones++
		} else if pixel == 2 {
			twos++
		}
	}

	fmt.Printf("Part 1: %d\n", ones*twos)

	fmt.Println("Part 2:")
	for i := 0; i < height; i++ {
		// fmt.Println(decoded[i*width : (i+1)*width])
		row := decoded[i*width : (i+1)*width]
		for _, pixel := range row {
			switch pixel {
			case 0:
				fmt.Print("  ")
			case 1:
				fmt.Print("# ")
			}
		}
		fmt.Print("\n")
	}
}
