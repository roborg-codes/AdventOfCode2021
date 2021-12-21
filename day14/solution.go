package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	// InputFile = "./sample.txt"
	InputFile = "./input.txt"

	PartOneStepN = 10
	PartTwoStepN = 40

	alphabet = 26
)

var (
	LastElement string
)

func main() {
	pt, it := ParseInput(InputFile)
	PartOne(pt, it)

	pt, _ = ParseInput(InputFile)
	PartTwo(pt, it)
}

type PolymerTemplate map[byte]map[byte]int

func (pt PolymerTemplate) Init() PolymerTemplate {
	pt = make(PolymerTemplate)
	for i := 0; i < alphabet; i++ {
		fP := ItA(i)
		pt[fP] = make(map[byte]int)
		for j := 0; j < alphabet; j++ {
			sP := ItA(j)
			pt[fP][sP] = 0
		}
	}
	return pt
}

func (pt PolymerTemplate) Copy() PolymerTemplate {
	dest := make(PolymerTemplate)

	for i := 0; i < alphabet; i++ {
		fP := ItA(i)
		dest[fP] = make(map[byte]int)
		for j := 0; j < alphabet; j++ {
			sP := ItA(j)
			dest[fP][sP] = pt[fP][sP]
		}
	}

	return dest
}

// Covert int to Ascii char
func ItA(x int) byte {
	return byte(x + 'A')
}

// ItA(x int) but as string
func ItAstr(x int) string {
	return string(ItA(x))
}

func (pt PolymerTemplate) Dump() {
	fmt.Printf("  ")
	for i := 0; i < alphabet; i++ {
		fmt.Printf("%*v", 3, ItAstr(i))
	}
	fmt.Println()

	for i := 0; i < alphabet; i++ {
		fmt.Printf("%*v", 2, ItAstr(i))
		for j := 0; j < alphabet; j++ {
			pFP := ItA(i)
			pSP := ItA(j)
			pN := pt[pFP][pSP]
			if pN < 1 {
				fmt.Printf("%*s", 3, ".")
			} else {
				fmt.Printf("%*v", 3, pN)
			}
		}
		fmt.Println()
	}
}

type InsertionTable map[string]byte

func (it InsertionTable) Dump() {
	for k, v := range it {
		fmt.Printf("%s -> %s\n", k, string(v))
	}
}

func (it InsertionTable) Lookup(a, b byte) byte {
	k := string(a) + string(b)
	if v, ok := it[k]; !ok {
		panic("entry not in insertion table")
	} else {
		return v
	}
}

func ParseInput(fname string) (PolymerTemplate, InsertionTable) {
	raw, err := ioutil.ReadFile(fname)
	if err != nil {
		panic(err)
	}
	inputByLine := strings.Split(string(raw), "\n")

	// Parsing template
	var pt PolymerTemplate
	pt = pt.Init()
	polymerString := inputByLine[0]
	for i := 0; i+1 < len(polymerString); i++ {
		polyFP := polymerString[i]
		polySP := polymerString[i+1]
		pt[polyFP][polySP]++
	}
	LastElement = string(polymerString[len(polymerString)-1])
	inputByLine = inputByLine[2:]

	// Parsing insertion table
	it := make(InsertionTable)
	for i := 0; inputByLine[i] != ""; i++ {
		line := inputByLine[i]
		lineValues := strings.Split(line, " -> ")
		k, v := lineValues[0], lineValues[1][0]
		it[k] = v
	}
	return pt, it
}

func CountElements(template PolymerTemplate) map[string]int {
	em := make(map[string]int)
	for i := 0; i < alphabet; i++ {
		for j := 0; j < alphabet; j++ {
			cFP := ItA(i)
			cSP := ItA(j)
			em[string(cFP)] += template[cFP][cSP]
		}
	}
	em[LastElement]++
	return em
}

func MinMaxElements(em map[string]int) (min, max int) {
	min = int(^uint(0) >> 1)

	for i := range em {
		if em[i] <= 0 {
			continue
		}

		if em[i] > max {
			max = em[i]
		}
		if em[i] < min {
			min = em[i]
		}
	}

	return
}

func PartOne(template PolymerTemplate, table InsertionTable) {
	for step := 1; step <= PartOneStepN; step++ {
		templateTmp := template.Copy()

		for i := 0; i < alphabet; i++ {
			for j := 0; j < alphabet; j++ {
				cFP := ItA(i)
				cSP := ItA(j)

				// Look up current values in template
				if template[cFP][cSP] <= 0 {
					continue
				}

				// lookup new values to insert
				value := table.Lookup(cFP, cSP)
				templateTmp[cFP][value] += template[cFP][cSP]
				templateTmp[value][cSP] += template[cFP][cSP]
				templateTmp[cFP][cSP]   -= template[cFP][cSP]
			}
		}
		template = templateTmp.Copy()
	}

	elementMap := CountElements(template)
	min, max := MinMaxElements(elementMap)

	fmt.Printf("Part one: %v\n", max - min)
}

func PartTwo(template PolymerTemplate, table InsertionTable) {
	for step := 1; step <= PartTwoStepN; step++ {
		templateTmp := template.Copy()

		for i := 0; i < alphabet; i++ {
			for j := 0; j < alphabet; j++ {
				cFP := ItA(i)
				cSP := ItA(j)

				if template[cFP][cSP] <= 0 {
					continue
				}

				value := table.Lookup(cFP, cSP)
				templateTmp[cFP][value] += template[cFP][cSP]
				templateTmp[value][cSP] += template[cFP][cSP]
				templateTmp[cFP][cSP]   -= template[cFP][cSP]
			}
		}
		template = templateTmp.Copy()
	}

	elementMap := CountElements(template)
	min, max := MinMaxElements(elementMap)

	fmt.Printf("Part two: %v\n", max - min)
}
