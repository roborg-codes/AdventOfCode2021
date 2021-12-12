package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	InputFile = "./input.txt"
	// InputFile = "./sample.txt"
	// InputFile = "./small_sample.txt"
)

func main() {
	input := ParseInput(InputFile)
	input.Dump()

	PartOne(input)

	input = ParseInput(InputFile)
	PartTwo(input)
}

func ParseInput(fname string) OctopusGrid {
	inputRaw, err := ioutil.ReadFile(fname)
	if err != nil {
		panic(err)
	}

	var input = make(OctopusGrid)
	for y, octopuses := range strings.Split(string(inputRaw), "\n") {
		if octopuses == "" {
			continue
		}
		input[y] = make(map[int]Octopus)
		for x, octopus := range octopuses {
			oStr, err := strconv.Atoi(string(octopus))
			if err != nil {
				panic(err)
			}
			// y -> key (row), x -> octopus map (col)
			input[y][x] = Octopus{Energy: int(oStr)}
		}
	}

	return input
}

type Octopus struct {
	Energy  int
	Flashed bool
}

type OctopusGrid map[int]map[int]Octopus

func (og OctopusGrid) Dump(step ...int) {
	fmt.Println("Grid dump:")
	if len(step) >= 1 {
		fmt.Printf("=step%3d====\n", step[0])
	} else {
		fmt.Println("============")
	}
	for y := 0; y < len(og); y++ {
		for x := 0; x < len(og[y]); x++ {
			 if og[y][x].Energy == 0 {
				fmt.Printf("\x1b[1m%d\x1b[0m", og[y][x].Energy)
				continue
			 }
			fmt.Printf("%d", og[y][x].Energy)
		}
		fmt.Println()
	}
	fmt.Println("============")
}

func (og OctopusGrid) Iterate(f func(y, x int)) {
	for y := 0; y < len(og); y++ {
		for x := 0; x < len(og[y]); x++ {
			f(y, x)
		}
	}
}

func (og OctopusGrid) Flash(y, x int) int {
	if octopus, ok := og[y][x]; ok {
		var flashes int
		if octopus.Energy < 9 {
			octopus.Energy++
			og[y][x] = octopus
			return flashes
		}
		octopus.Energy++
		if octopus.Energy > 9 && !octopus.Flashed {
			flashes += 1
			octopus.Flashed = true
			og[y][x] = octopus
			// octopus.Energy = 0 // <- might need to reset later

			flashes += og.Flash(y  , x-1)
			flashes += og.Flash(y  , x+1)
			flashes += og.Flash(y-1, x  )
			flashes += og.Flash(y+1, x  )
			flashes += og.Flash(y-1, x-1)
			flashes += og.Flash(y-1, x+1)
			flashes += og.Flash(y+1, x-1)
			flashes += og.Flash(y+1, x+1)
		}
		return flashes
	}
	return 0
}

func (og OctopusGrid) MakeStep() (flashes int) {
	// Increment everything
	og.Iterate(func(y, x int) {
		if o, ok := og[y][x]; ok {
			o.Energy++
			og[y][x] = o
		}
	})

	// Flash
	og.Iterate(func(y, x int) {
		if og[y][x].Energy > 9 && !og[y][x].Flashed {
			flashes += og.Flash(y, x)
		}
	})

	// Cleanup
	og.Iterate(func(y, x int) {
		if o, ok := og[y][x]; ok {
			if o.Flashed {
				o.Energy = 0
			}
			o.Flashed = false
			og[y][x] = o
		}
	})

	return
}

func PartOne(grid OctopusGrid) {
	var flashes int

	for i := 0; i < 195; i++ {
		// make a step
		flashes += grid.MakeStep()
	}

	fmt.Printf("Part one: %v\n", flashes)
}

// 100 flashes == sync

func PartTwo(grid OctopusGrid) {
	var flashes int
	var step int
	for ; flashes != 10*10; step++ {
		flashes = grid.MakeStep()
		grid.Dump(step)
	}
	fmt.Printf("Part two: %v\n", step)
}
