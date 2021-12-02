package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const inputFile = "./input.txt"

type command int64

const (
	forward command = iota
	down
	up
)

func (c command) String() string {
	switch c {
	case forward:
		return "forward"
	case down:
		return "down"
	case up:
		return "up"
	}
	return "unknown"
}

type pos struct {
	Horizontal int
	Depth      int
	Aim        int
}

func processCommand(command string) (comm string, val int) {
	cs := strings.Split(command, " ")
	if len(cs) == 1 {
		return
	}
	comm = cs[0]
	val, err := strconv.Atoi(cs[1])
	errcheck(err)
	return
}

func (p *pos) ExecuteCommand(command string) {
	comm, val := processCommand(command)
	if comm == "" {
		return // EOF
	}

	switch comm {
	case forward.String():
		p.Horizontal += val
		return
	case down.String():
		p.Depth += val
		return
	case up.String():
		p.Depth -= val
		return
	}
}


func (p *pos) ExecuteCommandCorrect(command string) {
	comm, val := processCommand(command)
	if comm == "" {
		return // EOF
	}

	switch comm {
	case forward.String():
		p.Horizontal += val
		p.Depth += p.Aim * val
	case down.String():
		p.Aim += val
	case up.String():
		p.Aim -= val
	}
}

func (p pos) PrintState() {
	fmt.Printf("===New state================\n")
	fmt.Printf("Current .Horizontal = %d\n", p.Horizontal)
	fmt.Printf("Current .Depth      = %d\n", p.Depth)
	fmt.Printf("Current .Aim        = %d\n", p.Aim)
	fmt.Printf("============================\n")
}

func (p pos) PrintSolution(title string) {
	fmt.Println(title)
	fmt.Printf("Horizontal - %d\n", p.Horizontal)
	fmt.Printf("Depth      - %d\n", p.Depth)
	fmt.Printf("Solution   - %d\n", p.Horizontal*p.Depth)
}

func errcheck(err error) {
	if err != nil {
		log.Fatalf("[ERR]: %v\n", err)
	}
}

func main() {
	inputRaw, err := ioutil.ReadFile(inputFile)
	errcheck(err)
	input := strings.Split(string(inputRaw), "\n")

	var position pos

	for _, val := range input {
		position.ExecuteCommand(val)
	}
	position.PrintSolution("First Solution:")
	position.Horizontal, position.Depth, position.Aim = 0, 0, 0

	for _, val := range input {
		position.ExecuteCommandCorrect(val)
	}
	position.PrintSolution("Second Solution:")
}
