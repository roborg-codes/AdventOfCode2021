package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

const (
	// InputFile = "./sample.txt"
	InputFile = "./input.txt"
)

func ParseInput(fname string) ([]string, []string) {
	inputRaw, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalf("[ERROR] %v\n", err)
	}
	re := regexp.MustCompile(` \| |\n`)
	return func(s []string) ([]string, []string) {
		var l []string
		var r []string
		for i := range s {
			if i%2 == 0 {
				l = append(l, s[i])
			} else {
				r = append(r, s[i])
			}
		}
		return l[:len(l)-1], r
	}(re.Split(string(inputRaw), -1))
}

func PartOne(input []string) {
	type Digit int
	// n of segments to activate
	const (
		one   Digit = 2
		seven Digit = 3
		four  Digit = 4
		eight Digit = 7
	)
	var count int

	for _, ds := range input {
		seq := strings.Split(ds, " ")
		for i := range seq {
			switch Digit(len(seq[i])) {
			case one:
				count++
			case four:
				count++
			case seven:
				count++
			case eight:
				count++
			}
		}
	}
	fmt.Printf("Part one: %d\n", count)
}


func StrToInt(d string) int {
	v, err := strconv.Atoi(d)
	if err != nil {
		log.Fatalf("[ERROR]: %v\b", err)
	}
	return v
}

func Decode(signals []string) {

}


func DecodeSignals(signals, values []string) int {
	// Score of each segment
	segments := make(map[string]int)
	for _, signal := range signals {
		for _, sigsegment := range signal {
			segments[string(sigsegment)] += 1
		}
	}

	// Score of each segment based on frequency of each digit
	digits := make([]int, len(signals))
	for i, signal := range signals {
		for _, sigsegment := range signal {
			digits[i] += segments[string(sigsegment)]
		}
	}

	FrequencyMap := make(map[int]string, len(digits))
	for _, digit := range digits {
		switch digit {
		case 17:
			FrequencyMap[digit] = "1"
		case 25:
			FrequencyMap[digit] = "7"
		case 30:
			FrequencyMap[digit] = "4"
		case 34:
			FrequencyMap[digit] = "2"
		case 37:
			FrequencyMap[digit] = "5"
		case 39:
			FrequencyMap[digit] = "3"
		case 41:
			FrequencyMap[digit] = "6"
		case 42:
			FrequencyMap[digit] = "0"
		case 45:
			FrequencyMap[digit] = "9"
		case 49:
			FrequencyMap[digit] = "8"
		}
	}

	DisplayScores := make([]int, len(values))
	DisplayValue := make([]string, len(values))
	for i, signal := range values {
		for _, sigsegment := range signal {
			DisplayScores[i] += segments[string(sigsegment)]
		}
		DisplayValue[i] = FrequencyMap[DisplayScores[i]]
	}
	return StrToInt(strings.Join(DisplayValue, ""))
}

func PartTwo(patterns, values []string) {
	if len(patterns) != len(values) {
		diff := len(patterns) - len(values)
		log.Fatalf("[ERROR]: Slices are not equal. Difference: %d\n", diff)
	}
	InLen := len(patterns)

	var DigitsSum int
	for i := 0; i < InLen; i++ {
		DigitsSum += DecodeSignals(strings.Split(patterns[i], " "), strings.Split(values[i], " "))
	}
	fmt.Printf("Part two: %d\n", DigitsSum)
}

func main() {
	Patterns, DisplayValues := ParseInput(InputFile)
	PartOne(DisplayValues)
	PartTwo(Patterns, DisplayValues)
}
