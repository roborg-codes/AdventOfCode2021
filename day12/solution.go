package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
)

const (
	// InputFile = "./sample.txt"
	// InputFile = "./sample2.txt"
	// InputFile = "./sample3.txt"
	InputFile = "./input.txt"
)

const (
	start CaveName = "start"
	end   CaveName = "end"
)

func main() {
	input := ParseInput(InputFile)
	PartOne(input)

	input = ParseInput(InputFile)
	PartTwo(input)
}

type CaveName string

type Cave struct {
	Value CaveName
	Small bool
}

func (c *Cave) String() string {
	return fmt.Sprintf("%v", c.Value)
}

func (c *Cave) IsSmall() bool {
	return c.Small
}

type CaveGraph struct {
	Edges    map[Cave][]*Cave
	Vertices []*Cave
}

func (cg *CaveGraph) Dump() {
	var s string
	for i := 0; i < len(cg.Vertices); i++ {
		s += cg.Vertices[i].String() + " -> "
		near := cg.Edges[*cg.Vertices[i]]
		for j := 0; j < len(near); j++ {
			s += near[j].String() + " | "
		}
		s += "\n"
	}
	fmt.Println(s)
}

func (cg *CaveGraph) AddCave(c *Cave) {
	cg.Vertices = append(cg.Vertices, c)
}

func (cg *CaveGraph) AddEdge(c1, c2 *Cave) {
	if cg.Edges == nil {
		cg.Edges = make(map[Cave][]*Cave)
	}
	if c1.Value == start {
		// no need to go back to the start
		cg.Edges[*c1] = append(cg.Edges[*c1], c2)
		return
	}
	cg.Edges[*c2] = append(cg.Edges[*c2], c1)
	cg.Edges[*c1] = append(cg.Edges[*c1], c2)
}

func (cg *CaveGraph) Start() *Cave {
	for i := 0; i < len(cg.Vertices); i++ {
		if cg.Vertices[i].Value == start {
			return cg.Vertices[i]
		}
	}
	return nil
}

func (cg *CaveGraph) End() *Cave {
	for i := 0; i < len(cg.Vertices); i++ {
		if cg.Vertices[i].Value == end {
			return cg.Vertices[i]
		}
	}
	return nil
}

func ParseInput(fname string) CaveGraph {
	InputRaw, err := ioutil.ReadFile(fname)
	if err != nil {
		panic(err)
	}
	MapEntries := strings.Split(string(InputRaw), "\n")

	CaveMap := make(map[string]*Cave)
	var CaveG CaveGraph
	for _, entry := range MapEntries[:len(MapEntries)-1] {
		caves := strings.Split(entry, "-")
		pointA, pointB := caves[0], caves[1]
		// Keep track of already added caves
		if _, ok := CaveMap[pointA]; !ok {
			CaveMap[pointA] = &Cave{
				Value: CaveName(pointA),
			}
			// Mark small ones
			if unicode.IsLower(rune(pointA[0])) {
				CaveMap[pointA].Small = true
			}
			CaveG.AddCave(CaveMap[pointA])
		}
		if _, ok := CaveMap[pointB]; !ok {
			CaveMap[pointB] = &Cave{
				Value: CaveName(pointB),
			}
			if unicode.IsLower(rune(pointB[0])) {
				CaveMap[pointB].Small = true
			}
			CaveG.AddCave(CaveMap[pointB])
		}
		// Connect two caves
		CaveG.AddEdge(CaveMap[pointA], CaveMap[pointB])
	}
	return CaveG
}

func (cg *CaveGraph) Traverse(source *Cave, paths []*Cave, visited map[*Cave]bool) int {
	var pathsN int

	visited[source] = true
	paths = append(paths, source)
	if source.Value == end {
		pathsN++
	} else {
		near := cg.Edges[*source]
		for n := range near {
			if !visited[near[n]] || visited[near[n]] && !near[n].Small {
				pathsN += cg.Traverse(near[n], paths, visited)
			}
		}
	}
	visited[source] = false
	return pathsN
}

func PartOne(CaveG CaveGraph) {
	startpoint := CaveG.Start()
	paths := make([]*Cave, 0)
	visited := make(map[*Cave]bool)

	solution := CaveG.Traverse(startpoint, paths, visited)

	fmt.Printf("Part one: %v\n", solution)
}

func OnPath(cave *Cave, path []*Cave) bool {
	for i := range path {
		if path[i].Value == cave.Value {
			return true
		}
	}
	return false
}

func (cg *CaveGraph) Traverse2(source *Cave, path []*Cave, visited map[*Cave]bool, doubleSmall bool) int {
	var pathsN int
	visited[source] = true
	path = append(path, source)

	if source.Value == end {
		pathsN++
		goto end
	}

	for _, cave := range cg.Edges[*source] {
		if cave.Value == start {
			continue
		}
		if OnPath(cave, path) && cave.Small {
			if !doubleSmall {
				pathsN += cg.Traverse2(cave, path, visited, true)
			}
			continue
		}
		if !visited[cave] || !cave.Small {
			pathsN += cg.Traverse2(cave, path, visited, doubleSmall)
			continue
		}
	}

end:
	visited[source] = false
	return pathsN
}

func PartTwo(CaveG CaveGraph) {
	startpoint := CaveG.Start()
	path := make([]*Cave, 0)
	visited := make(map[*Cave]bool)
	doubleSmall := false
	solution := CaveG.Traverse2(startpoint, path, visited, doubleSmall)

	fmt.Printf("Part two: %v\n", solution)
}
