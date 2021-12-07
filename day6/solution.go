package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const (
	inputFile = "./input.txt"
	// inputFile = "./sample.txt"
)

func ParseInput(fname string) (input []int) {
	inputRaw, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalf("[ERROR] While reading file: %v\n", err)
	}

	for _, inputStr := range strings.Split(string(inputRaw), ",") {
		inputStr = strings.TrimSuffix(inputStr, "\n")
		in, err := strconv.Atoi(inputStr)
		if err != nil {
			log.Fatalf("[ERROR] While reading file: %v\n", err)
		}
		input = append(input, in)
	}
	return input
}

func PartOne(input []int) {
	days := 80
	for day := 0; day < days; day++ {
		lenAtDay := len(input)
		for i := 0; i < lenAtDay; i++ {
			switch input[i] {
			case 0:
				input[i] = 6
				input = append(input, 8)
			default:
				input[i]--
			}
		}
	}
	fmt.Printf("Amt of lanternfish after %d days: %d\n", 80, len(input))
}

func PartTwo(input []int)  {
	days := 256
	var fish = make([]int, 9)
	for _, d := range input {
		fish[d]++
	}

	for day := 0; day < days; day++ {
		nextFish := fish[0]
		for i := 0; i < len(fish)-1; i++ {
			fish[i] = fish[i+1]
		}
		fish[6] += nextFish
		fish[8] = nextFish
	}

	var ans int
	for _, v := range fish {
		ans += v
	}
	fmt.Printf("Amt of lanternfish after %d days: %d\n", days, ans)
}

func main() {
	input := ParseInput(inputFile)
	PartOne(input)

	input = ParseInput(inputFile)
	PartTwo(input)
}
