package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	inputFile = "./input.txt"
	// inputFile = "./sample.txt"

	boardStep = 5
	boardLen = boardStep * 5
)

type Numbers struct {
	List []int
}

func GetNumbers(input *string) *Numbers {
	splitInput := strings.Split(*input, ",")
	list := make([]int, len(splitInput))

	for i := range list {
		intVal, err := strconv.Atoi(splitInput[i])
		 if err != nil {
			fmt.Printf("[ERROR]: Converting ascii to int. %v\n", err)
		 }
		 list[i] = intVal
	}
	return &Numbers{
		List: list,
	}
}

type Cell struct {
	Marked bool
	Value  int
}

func (c *Cell) Mark(x int) {
	if c.Value == x {
		c.Marked = true
	}
}

func MakeCell(x int) *Cell {
	return &Cell{
		Value: x,
	}
}

type Board []Cell

func (b Board) MarkNumberOnBoard(x int) {
	for i := range b {
		b[i].Mark(x)
	}
}

func (b Board) checkCol(index int) bool {
	nextIndex := index+boardStep
	if nextIndex >= boardLen-5 && b[nextIndex].Marked {
		return true
	} else if b[nextIndex].Marked {
		return b.checkCol(nextIndex)
	}
	return false
}

func (b Board) cehckRow() Board {
	var seqCounter int
	for index := 0; index < boardLen; index++ {
		if seqCounter == 5 {
			return b
		}
		if index % 5 == 0 {
			seqCounter = 0
		}
		if b[index].Marked {
			seqCounter++
		}
	}
	return nil
}

func (b Board) CheckBingo() Board {
	var counter int

	// Columns
	for i := 0; i < 5; i ++ {
		if !b[i].Marked {
			continue
		}
		for j := i; j < boardLen; j += boardStep {
			if b[j].Marked {
				counter++
			} else {
				counter = 0
				break
			}
		}
		if counter == 5 {
			return b
		}
	}

	// Rows
	counter = 0
	for i := 0; i <= boardLen - boardStep; i += boardStep {
		for j := i; j < boardLen; j++ {
			if b[j].Marked {
				counter++
			} else {
				counter = 0
				break
			}
			if counter == 5 {
				return b
			}
		}
	}
	return nil
}

type Game struct {
	Board []Board
}

func (b Board) Dump() {
	fmt.Println("===State of Board========")
	for i, j := 0, 0; i < boardLen; i++ {
		if b[i].Marked {
			fmt.Printf(" [%*d] ", 2, b[i].Value)
		} else {
			fmt.Printf("  %*d  ", 2, b[i].Value)
		}
		j++
		if j == boardStep {
			j = 0
			fmt.Println()
		}
	}
	fmt.Println()
	fmt.Println("=========================")
}

func (g Game) Dump() {
	fmt.Println("===State of the Game================================")
	for _, b := range g.Board {
		b.Dump()
	}
	fmt.Println("====================================================")
}

func (g *Game) CheckBingo() (Board, int) {
	for i, board := range g.Board {
		if winnerBoard := board.CheckBingo(); winnerBoard != nil {
			return board, i
		}
	}

	return nil, -1
}

func (g *Game) MarkNumbers(number int) {
	for i := range g.Board {
		g.Board[i].MarkNumberOnBoard(number)
	}
}

func (g *Game) CalculateSumUnmarked(winnerBoard Board) int {
	var sumUnmarked int

	for i := range winnerBoard {
		if !winnerBoard[i].Marked {
			sumUnmarked += winnerBoard[i].Value
			continue
		}
	}

	return sumUnmarked
}

func PopulateGame(input []string) Game {
	var game Game
	var board Board
	var boardCellCount int

	for i := 0; i < len(input); i++ {
		if input[i] == "" {
			continue
		}
		splitInput := strings.Split(input[i], " ")
		for _, v := range splitInput {
			if v == "" {
				continue
			}
			strc, _ := strconv.Atoi(v)
			cell := MakeCell(strc)
			board = append(board, *cell)
		}
		boardCellCount += 5

		if boardCellCount == boardLen {
			game.Board = append(game.Board, board)
			board = Board{}
			boardCellCount = 0
		}
	}

	return game
}

func pop(g []Board, i int) []Board {
	g[i] = g[len(g)-1]
	return g[:len(g)-1]
}

func PartOne(numbers *Numbers, game Game) {
	var lastNumber int
	var winnerBoard Board

	for _, number := range numbers.List {
		game.MarkNumbers(number)

		if winnerBoard, _ = game.CheckBingo(); winnerBoard != nil {
			lastNumber = number
			break
		}
	}
	winnerBoard.Dump()
	fmt.Printf("Last drawn number: %d\n", lastNumber)
	fmt.Printf("Sum of unmarked: %d\n", game.CalculateSumUnmarked(winnerBoard))
	fmt.Printf("Part one final score: %d\n", game.CalculateSumUnmarked(winnerBoard) * lastNumber)
}

func PartTwo(numbers *Numbers, game Game) {
	var loserBoard Board = nil
	var lastNumber int

	for _, number := range numbers.List {
		lastNumber = number
		game.MarkNumbers(number)
		if w, i := game.CheckBingo(); w != nil && len(game.Board) == 1 {
			loserBoard = game.Board[i]
			break
		}
		PopAllMarkedBoards(&game)
	}
	loserBoard.Dump()
	fmt.Printf("Final number: %d\n", lastNumber)
	fmt.Printf("Sum of unmarked: %d\n", game.CalculateSumUnmarked(loserBoard))
	fmt.Printf("Part two final score:  %d\n", game.CalculateSumUnmarked(loserBoard) * lastNumber)
}

func ParseInput() (input []string) {
	inputRaw, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("[ERROR]: ReadFile incorrect. %v\n", err)
		return
	}
	input = strings.Split(string(inputRaw), "\n")

	return
}

func PopAllMarkedBoards(game *Game) {
	winner, i := game.CheckBingo()
	if winner == nil {
		return
	}
	for ;winner != nil;{
		game.Board = pop(game.Board, i)
		winner, i = game.CheckBingo()
	}
}

func main() {
	input := ParseInput()
	numbers := GetNumbers(&input[0])
	fmt.Printf("Numbers to draw: %v\n", numbers.List)
	input = input[1:]

	// Part one solution
	game := PopulateGame(input)
	PartOne(numbers, game)

	// Part two solution
	game = PopulateGame(input)
	PartTwo(numbers, game)
}
