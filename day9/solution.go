package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

const (
	// InputFile = "./sample.txt"
	InputFile = "./input.txt"
)

func main() {
	input := ParseInput(InputFile)

	PartOne(input)

	var inputPt2 = make(HeightMapSB)
	for i := range input {
		for j := range input[i] {
			pt := Point{
				Value: input[i][j],
				Visited: false,
			}
			inputPt2[i] = append(inputPt2[i], pt)
		}
	}

	PartTwo(inputPt2)
}

type HeightMap map[int][]int

func (hm HeightMap) Dump() {
	for k, v := range hm {
		fmt.Printf("%d: %v\n", k, v)
	}
}

type Point struct {
	Value int
	Visited bool
}

type HeightMapSB map[int][]Point

func (hm HeightMapSB) Dump() {
	for k, v := range hm {
		for _, p := range v {
			fmt.Printf("row[%d]: %t -> %d\n", k, p.Visited, p.Value)
		}
		fmt.Println("====================")
	}
}

func (hm HeightMapSB) GetBasin(col, row int) (BasinSize int) {
	pt := hm[row][col]
	if pt.Value == 9 || pt.Visited {
		return 0
	}

	hm[row][col].Visited = true
	BasinSize += 1

	var (
		top    int
		left   int
		right  int
		bottom int
	)

	// Handle edge cases
	if row != 0 {
		// check positions above
		top = hm[row-1][col].Value
	}
	if row != len(hm)-1 {
		// check positions below
		bottom = hm[row+1][col].Value
	}
	if col != 0 {
		// check positions on the left
		left = hm[row][col-1].Value
	}
	if col != len(hm[row])-1 {
		// check positions on the right
		right = hm[row][col+1].Value
	}

	if top > pt.Value {
		BasinSize += hm.GetBasin(col, row-1)
	}
	if bottom > pt.Value {
		BasinSize += hm.GetBasin(col, row+1)
	}
	if left > pt.Value {
		BasinSize += hm.GetBasin(col-1, row)
	}
	if right > pt.Value {
		BasinSize += hm.GetBasin(col+1, row)
	}

	return
}

func ParseInput(fname string) HeightMap {
	inputRaw, err := ioutil.ReadFile(fname)
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(inputRaw), "\n")

	Map := make(map[int][]int)
	for x := range input[:len(input)-1] {
		for _, point := range input[x] {
			intv, err := strconv.Atoi(string(point))
			if err != nil {
				panic(err)
			}
			Map[x] = append(Map[x], intv)
		}
	}
	return Map
}

func PartOne(hm HeightMap) {
	// hm.Dump()
	var count int

	for rI, row := range hm {
		for pI, point := range row {
			var (
				top    = math.MaxInt
				left   = math.MaxInt
				right  = math.MaxInt
				bottom = math.MaxInt
			)

			// Handle edge cases
			if rI != 0 {
				// skip checking positions above
				top = hm[rI-1][pI]
			}
			if rI != len(hm)-1 {
				// skip checking positions below
				bottom = hm[rI+1][pI]
			}
			if pI != 0 {
				// skip checking positions on the left
				left = row[pI-1]
			}
			if pI != len(row)-1 {
				// skip checking positions on the right
				right = row[pI+1]
			}

			if point < top && point < bottom &&
				point < left && point < right {

				count += point+1
			}
		}
	}
	fmt.Printf("Part one: %v\n", count)
}

func SubstSmallestInt(s []int, n int) {
	smallest := math.MaxInt
	for i := range s {
		if s[i] < smallest {
			smallest = s[i]
		}
	}
	for i := range s {
		if s[i] == smallest {
			s[i] = n
			return
		}
	}
}

func PartTwo(hm HeightMapSB) {
	var LargetsBasins = make([]int, 3)

	for rI, row := range hm {
		for pI, point := range row {
			var (
				top    = math.MaxInt
				left   = math.MaxInt
				right  = math.MaxInt
				bottom = math.MaxInt
			)

			if rI != 0 {
				top = hm[rI-1][pI].Value
			}
			if rI != len(hm)-1 {
				bottom = hm[rI+1][pI].Value
			}
			if pI != 0 {
				left = row[pI-1].Value
			}
			if pI != len(row)-1 {
				right = row[pI+1].Value
			}

			if point.Value < top && point.Value < bottom &&
				point.Value < left && point.Value < right {

				basin := hm.GetBasin(pI, rI)
				for i := range LargetsBasins {
					if basin > LargetsBasins[i] {
						SubstSmallestInt(LargetsBasins, basin)
						break
					}
				}
			}
		}
	}
	var total = 1
	for i := range LargetsBasins {
		total *= LargetsBasins[i]
	}
	fmt.Printf("Part two: %v\n", total)
}
