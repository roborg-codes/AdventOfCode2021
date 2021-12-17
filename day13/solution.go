package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	// InputFile = "./sample.txt"
	InputFile = "./input.txt"
)

func main() {
	paper, instructions := ParseInput(InputFile)

	PartOne(&paper, instructions[0])
	PartTwo(&paper, instructions[1:])
}

type Paper struct {
	Grid [][]bool
}

func (p *Paper) Dump() {
	for y := 0; y < len(p.Grid); y++ {
		for x := 0; x < len(p.Grid[y]); x++ {
			if p.Grid[y][x] {
				fmt.Printf("■")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
}

func (p *Paper) Iterate(f func(y, x int)) {
	for y := 0; y < len(p.Grid); y++ {
		for x := 0; x < len(p.Grid[y]); x++ {
			f(y, x)
		}
	}
}

type Axis string

const (
	x Axis = "x"
	y Axis = "y"
)

type Instruction struct {
	Axis Axis
	Pos  int
}

type Instructions []*Instruction

func (i Instructions) Dump() {
	for _, j := range i {
		fmt.Printf("Fold along %s=%d\n", j.Axis, j.Pos)
	}
}

func ParseInput(fname string) (Paper, Instructions) {
	bytes, err := ioutil.ReadFile(fname)
	if err != nil {
		panic(err)
	}
	dots := strings.Split(string(bytes), "\n")

	// Process points
	var i, maxY, maxX int
	// Map `points` for boundaries
	pointMap := make(map[int][]int)
	for ; dots[i] != ""; i++ {
		dotsSplit := strings.Split(dots[i], ",")
		x, _ := strconv.Atoi(dotsSplit[0])
		if x > maxX {
			maxX = x
		}
		y, _ := strconv.Atoi(dotsSplit[1])
		if y > maxY {
			maxY = y
		}
		pointMap[y] = append(pointMap[y], x)
	}
	// Fill in the paper
	var paper Paper
	paper.Grid = make([][]bool, maxY+1)
	for y := range pointMap {
		for _, x := range pointMap[y] {
			// allocate row if none already present
			if len(paper.Grid[y]) < 1 {
				paper.Grid[y] = make([]bool, maxX+1)
			}
			paper.Grid[y][x] = true
		}
	}
	// Fill out empty rows
	for y := 0; y < len(paper.Grid); y++ {
		if len(paper.Grid[y]) < 1 {
			paper.Grid[y] = make([]bool, maxX+1)
		}
	}

	// Process instructions
	dots = dots[i+1:]
	instructions := make(Instructions, 0)
	for i = 0; dots[i] != ""; i++ {
		ap := strings.SplitAfter(dots[i], " ")
		axis := Axis(ap[2:][0][0])
		pos, _ := strconv.Atoi(string(ap[2:][0][2:]))
		instructions = append(instructions, &Instruction{
			Axis: axis,
			Pos: pos,
		})
	}
	return paper, instructions
}

func AbsInt(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func (p *Paper) FoldAlongY(foldPos int) *Paper {
	for y := foldPos; y < len(p.Grid); y++ {
		for x := 0; x < len(p.Grid[y]); x++ {
			if !p.Grid[y][x] {
				continue
			}
			newY := AbsInt(y - foldPos*2)
			p.Grid[newY][x] = true
		}
	}
	p.Grid = p.Grid[:foldPos]
	return p
}

func (p *Paper) FoldAlongX(foldPos int) *Paper {
	for y := 0; y < len(p.Grid); y++ {
		for x := foldPos; x < len(p.Grid[y]); x++ {
			if !p.Grid[y][x] {
				continue
			}
			newX := AbsInt(x - foldPos*2)
			p.Grid[y][newX] = true
		}
	}
	for y := 0; y < len(p.Grid); y++ {
		p.Grid[y] = p.Grid[y][:foldPos]
	}
	return p
}

func (p *Paper) CountPoints() (count int) {
	for y := 0; y < len(p.Grid); y++ {
		for x := 0; x < len(p.Grid[y]); x++ {
			if p.Grid[y][x] {
				count++
			}
		}
	}
	return
}

func PartOne(paper *Paper, instruction *Instruction) {
	switch instruction.Axis {
	case y:
		paper.FoldAlongY(instruction.Pos)
	case x:
		paper.FoldAlongX(instruction.Pos)
	}

	fmt.Printf("Part one: %d\n", paper.CountPoints())
}

func PartTwo(paper *Paper, instructions Instructions) {
	for _, instr := range instructions {
		switch instr.Axis {
		case y:
			paper.FoldAlongY(instr.Pos)
		case x:
			paper.FoldAlongX(instr.Pos)
		}
	}
	fmt.Println("Part two:")
	paper.Dump()
}
