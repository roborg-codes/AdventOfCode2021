package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

const (
	inputFile         = "./input.txt"
	lineLen           = 13
	inputLineLen      = lineLen - 1
	lineCount         = 1000
	mask         rate = 0b111111111111
	one          byte = 49
	zero         byte = 48
)

type rate uint
type pos int
type count int

func errcheck(err error) {
	if err != nil {
		log.Fatalf("[ERR]: %v\n", err)
	}
}

// Split input into slices representing each input line (w/o '\n')
func SplitBSlice(input []byte) (bs [][]byte) {
	sLen := len(input)

	for i, j := 0, 0; i < sLen; i += lineLen {
		j += lineLen
		if j > sLen {
			j = sLen
		}
		bs = append(bs, input[i:j-1])
	}
	return
}

func ReverseBitOrder(input rate) (output rate) {
	for i := 0; i < inputLineLen; i++ {
		output = (output << 1) | (input & 0b1)
		input >>= 1
	}
	return
}

func CalculateGammaRate(valMap map[pos]count) (gamma rate) {
	for key, valCount := range valMap {
		if valCount > lineCount/2 {
			// set bit to 1 at inversed position
			gamma |= gamma ^ (1 << key)
		}
	}
	gamma = ReverseBitOrder(gamma) // restore original order
	return
}

func CalculateEpsilonRate(gamma rate) (epsilon rate) {
	return gamma ^ mask // flip first 14 bits
}

func RemoveLineFromSSlice(bss [][]byte, index int, dropValue byte) [][]byte {
	if len(bss) == 1 {
		return bss
	}
	for line := len(bss) - 1; line >= 0; line-- {
		if bss[line][index] == dropValue {
			// drop line
			bss = append(bss[:line], bss[line+1:]...)
		}
	}
	return bss
}

// Logic for determining which value of index bit to consider
// Oxygen generator diagnostics default to 1 if
// there is equal amt of occurances of 1s and 0s
func getOxygenDropBit(bs [][]byte, index int) byte {
	var bitCount int
	for _, value := range bs {
		if value[index] == one {
			bitCount++
		} else {
			bitCount--
		}
	}
	fmt.Printf("[Oxygen] bitCount: %d\n", bitCount)
	if bitCount < 0 {
		// One is the is the least common, drop it
		return one
	}
	return zero
}

func CalculateOxygenGeneratorRate(bs [][]byte) (oxygen rate) {
	var dropValue byte

	for index := 0; index <= inputLineLen; index++ {
		if len(bs) == 1 {
			break
		} else if len(bs) < 1 {
			fmt.Printf("[ERROR]: bs is empty :9\n")
			return
		}
		// var index is index of bit to consider
		// var dropValue is byte used to match lines
		dropValue = getOxygenDropBit(bs, index)
		prv := len(bs)
		bs = RemoveLineFromSSlice(bs, index, dropValue)
		crv := len(bs)
		fmt.Printf("Iterating over index %d finished. %d items of `%d` deleted. %d items remaining.\n", index, prv-crv, dropValue, crv)
	}

	for key, value := range bs[0] {
		if value == one {
			oxygen |= oxygen ^ (1 << key)
			fmt.Printf("oxygen | oxygen & (1 << %b) = %b\n", key, oxygen)
		}
	}
	oxygen = ReverseBitOrder(oxygen)
	return
}

// Logic for determining which value of index bit to consider
// CO2 scrubberr diagnostics default to 0 if
// there is equal amt of occurances of 1s and 0s
func getScrubberDropBit(bs [][]byte, index int) byte {
	var bitCount int
	for _, value := range bs {
		if value[index] == one {
			bitCount++
		} else {
			bitCount--
		}
	}
	if bitCount >= 0 {
		// One is the most common or less prioritized, drop it
		return one
	}
	return zero
}

func CalculateCO2ScrubberRate(bs [][]byte) (scrubber rate) {
	var dropValue byte

	for index := 0; index <= inputLineLen; index++ {
		if len(bs) == 1 {
			break
		} else if len(bs) < 1 {
			fmt.Printf("[ERROR]: bs is empty :9\n")
			return
		}
		dropValue = getScrubberDropBit(bs, index)
		prv := len(bs)
		bs = RemoveLineFromSSlice(bs, index, dropValue)
		crv := len(bs)
		fmt.Printf("Iterating over index %d finished. %d items of `%d` deleted. %d items remaining.\n", index, prv-crv, dropValue, crv)
	}
	for key, value := range bs[0] {
		if value == one {
			scrubber |= scrubber ^ (1 << key)
			fmt.Printf("scrubber | scrubber & (1 << %b) = %b\n", key, scrubber)
		}
	}
	scrubber = ReverseBitOrder(scrubber)
	return
}

func main() {
	inputRaw, err := ioutil.ReadFile(inputFile)
	errcheck(err)
	bss := SplitBSlice(inputRaw)
	valMap := make(map[pos]count, inputLineLen)
	for _, line := range bss {
		for bytepos, byte := range line {
			if byte == one {
				valMap[pos(bytepos)]++
			}
		}
	}

	var gamma = CalculateGammaRate(valMap)
	var epsilon = CalculateEpsilonRate(gamma)
	var solution = gamma * epsilon

	fmt.Printf("GAMMA:            %012b\n", gamma)
	fmt.Printf("EPSILON:          %012b\n", epsilon)
	fmt.Printf("SOLUTION ONE:     %12d\n", solution)

	bss = SplitBSlice(inputRaw)
	var oxygen = CalculateOxygenGeneratorRate(bss)

	bss = SplitBSlice(inputRaw)
	var co2Scrubber = CalculateCO2ScrubberRate(bss)
	solution = oxygen * co2Scrubber

	fmt.Printf("OXYGEN GENERATOR: %012b\n", oxygen)
	fmt.Printf("CO2 SCRUBBER:     %012b\n", co2Scrubber)
	fmt.Printf("SOLUTION TWO:     %12d\n", solution)
}
