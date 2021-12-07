package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
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

func Range(min, max int) []int {
	nums := make([]int, 0, max-min+1)
	for i := min; i <= max; i++ {
		nums = append(nums, i)
	}
	return nums
}

func IntAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func MinMax(list []int) (min, max int) {
	min = math.MaxInt
	for i := range list {
		if list[i] > max {
			max = list[i]
		} else if list[i] < min {
			min = list[i]
		}
	}
	return
}

func PartOneAndTwo(input []int) {
	var fuel1, fuel2 int
	min, max := MinMax(input)
	posRange := Range(min, max)

	for _, pos := range posRange {
		var sum1, sum2 int
		for i := range input {
			d := IntAbs(pos-input[i])
			sum1 += d
			sum2 += d * (d + 1) / 2
		}
		if pos == min || sum1 < fuel1 {
			fuel1 = sum1
		}
		if pos == min || sum2 < fuel2 {
			fuel2 = sum2
		}
	}
	fmt.Printf("Part 1: %d\n", fuel1)
	fmt.Printf("Part 2: %d\n", fuel2)
}

func main() {
	input := ParseInput(inputFile)
	PartOneAndTwo(input)
}
