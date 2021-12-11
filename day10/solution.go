package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

const (
	// InputFile = "./sample.txt"
	InputFile = "./input.txt"
)

func main() {
	input := ParseInput(InputFile)

	PartOne(input)
	PartTwo(input)
}

func ParseInput(fname string) []string {
	inputRaw, err := ioutil.ReadFile(fname)
	if err != nil {
		panic(err)
	}
	parsed := strings.Split(string(inputRaw), "\n")
	return parsed[:len(parsed)-1]
}

type Stack []rune

func (s Stack) Push(x rune) Stack {
	s = append(s, 0)
	copy(s[1:], s)
	s[0] = x
	return s
}

func (s Stack) Pop() (rune, Stack) {
	return s[0], s[1:]
}

func (s Stack) Dump()  {
	fmt.Println("Stack dump:")
	for i := range s {
		fmt.Printf("[%3d]: %v\n", i, string(s[i]))
	}
	fmt.Println("========================")
}

// Calculates score for the particular line (not discarded)
func (s Stack) CalculateCompletionScore() (score int) {
	ChunkTable := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}
	ScoreTable := map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}

	for i := range s {
		score *= 5
		score += ScoreTable[ChunkTable[s[i]]]
	}
	return score
}

func PartOne(input []string) {
	ScoreTable := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	var score int


	// 'opening' -> 'closing'
	ChunkTable := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}

	for _, line := range input {
		var stack Stack
		for _, chunk := range line {
			if _, ok := ChunkTable[chunk]; ok {
				// the brace encountered is opening
				// just push it on the stack
				stack = stack.Push(chunk)
			} else {
				// the brace is closing
				// compare to the last stack item
				var sc rune
				sc, stack = stack.Pop()
				if ChunkTable[sc] != chunk {
					score += ScoreTable[chunk]
					break
				}
			}
		}
	}
	fmt.Printf("Part one: %v\n", score)
}

func PartTwo(input []string) {
	ChunkTable := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}

	var scores []int
	for _, line := range input {
		var stack Stack
		var discarded bool
		for _, chunk := range line {
			if _, ok := ChunkTable[chunk]; ok {
				// the brace encountered is opening
				// just push it on the stack
				stack = stack.Push(chunk)
			} else {
				// the brace is closing
				// compare to the last stack item
				var sc rune
				sc, stack = stack.Pop()
				if ChunkTable[sc] != chunk {
					// discard line
					discarded = true
					break
				}
			}
		}
		if !discarded {
			scores = append(scores, stack.CalculateCompletionScore())
		}
	}
	sort.Ints(scores)
	fmt.Printf("Part two: %d\n", scores[len(scores)/2])
}
