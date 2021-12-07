package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	inputFile = "./input.txt"
	boardSize = 1000
	// inputFile = "./sample.txt"
	// boardSize = 10
)

type Entry struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func (e *Entry) Dump() {
	fmt.Printf("(%3d, %3d) -> (%3d, %3d)\n", e.x1, e.y1, e.x2, e.y2)
}

func DumpEntries(entries []Entry) {
	for i := range entries {
		entries[i].Dump()
	}
}

type Map [boardSize][boardSize]int

func Sign(n int) int {
	if n > 0 {
		return 1
	}
	if n < 0 {
		return -1
	}
	return 0
}

func (m *Map) DrawLinesPartOne(entries []Entry) {
	for _, entry := range entries {
		if entry.x1 == entry.x2 || entry.y1 == entry.y2 {
			n := int(math.Max(
				math.Abs(float64(entry.x2) - float64(entry.x1))+1,
				math.Abs(float64(entry.y2) - float64(entry.y1))+1),
			)

			for j := 0; j < n; j++ {
				x := entry.x1 + Sign(entry.x2 - entry.x1) * j
				y := entry.y1 + Sign(entry.y2 - entry.y1) * j
				m[y][x]++
			}
		}
	}
}

func (m *Map) DrawLinesPartTwo(entries []Entry) {
	for _, entry := range entries {
		n := int(math.Max(
			math.Abs(float64(entry.x2) - float64(entry.x1))+1,
			math.Abs(float64(entry.y2) - float64(entry.y1))+1),
		)

		for j := 0; j < n; j++ {
			x := entry.x1 + Sign(entry.x2 - entry.x1) * j
			y := entry.y1 + Sign(entry.y2 - entry.y1) * j
			m[y][x]++
		}
	}
}

func (m Map) Dump() {
	for y := 0; y < boardSize; y++ {
		for x := 0; x < boardSize; x++ {
			pos := m[y][x]
			if pos == 0 {
				fmt.Printf(".")
				continue
			}
			fmt.Printf("%d", pos)
		}
		fmt.Println()
	}
}

func (m Map) CountOverlaps() (count int) {
	for y := 0; y < boardSize; y++ {
		for x :=0; x < boardSize; x++ {
			pos := m[y][x]
			if pos >= 2 {
				count++
			}
		}
	}
	return
}

func errcheck(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func ParseInput(input []string) []Entry {
	var entries []Entry

	for _, strinput := range input {
		if strinput == "" {
			break
		}
		coordinates := strings.Split(strinput, " -> ")
		begin := strings.Split(coordinates[0], ",")
		end := strings.Split(coordinates[1], ",")
		x1, err := strconv.Atoi(begin[0])
		errcheck(err)
		x2, err := strconv.Atoi(end[0])
		errcheck(err)
		y1, err := strconv.Atoi(begin[1])
		errcheck(err)
		y2, err := strconv.Atoi(end[1])
		errcheck(err)

		entry := Entry{
			x1: x1,
			y1: y1,
			x2: x2,
			y2: y2,
		}
		entries = append(entries, entry)
	}

	return entries
}

func main() {
	inputRaw, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("[ERROR]: ReadFile incorrect. %v\n", err)
		return
	}
	input := strings.Split(string(inputRaw), "\n")
	entries := ParseInput(input)
	entryMap := Map{}

	entryMap.DrawLinesPartOne(entries)
	fmt.Printf("First part answer: %d\n", entryMap.CountOverlaps())

	entryMap = Map{}
	entryMap.DrawLinesPartTwo(entries)
	fmt.Printf("Second part answer: %d\n", entryMap.CountOverlaps())
}
