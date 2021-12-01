package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const inputFile = "./input.txt"

type window struct {
	A int
	B int
	C int
}

func (w window) Sum() int {
	return w.A + w.B + w.C
}

func errcheck(err error) {
	if err != nil {
		log.Fatalf("[ERR]: %v\n", err)
	}
}

func splitInput(input string) []int {
	strInp := strings.Split(input, "\n")
	ints := make([]int, len(strInp))

	for i, s := range strInp {
		ints[i], _ = strconv.Atoi(s)
	}
	return ints
}

func main() {
	inputRaw, err := ioutil.ReadFile(inputFile)
	errcheck(err)
	input := splitInput(string(inputRaw))
	var incCount int

	var prevVal int
	for i, val := range input {
		if i == 0 {
			prevVal = val
			continue // no previous measurment
		}
		if val > prevVal {
			incCount++
		}
		prevVal = val
	}
	fmt.Printf("Increases of values: %d\n", incCount)
	incCount = 0


	var prevWin window
	var currWin window
	for i, val := range input {
		if i+1 == len(input) { break } // index out of range, no comparison possible
		if i < 2 {
			if i == 1 { prevWin.A, prevWin.B, prevWin.C = input[i-1], val, input[i+1] }
			continue
		}

		// Start at index 3
		currWin.A, currWin.B, currWin.C = input[i-1], val, input[i+1]
		if currWin.Sum() > prevWin.Sum() {
			incCount++
		}
		prevWin.A, prevWin.B, prevWin.C = input[i-1], val, input[i+1]
	}
	fmt.Printf("Increases of 3-value windows: %d\n", incCount)
}
