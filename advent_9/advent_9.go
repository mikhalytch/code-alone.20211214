package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	filename := "advent_9.test.txt"
	inputFile := readAdvent9File(filename)
	result := calcAdvent9Result(inputFile)
	fmt.Printf("Answer is: %v (and amount of so-long chains: %d)\n", result.answer, result.amount)
}

type chain struct {
	startingIndex int
	entries       map[int]bool
}

func (c *chain) add(new int) bool {
	if c.entries[new] {
		return false
	}
	c.entries[new] = true
	return true
}
func newChain(capacity int, startingIdx int) chain {
	c := chain{startingIdx, make(map[int]bool, capacity)}
	return c
}

type chainVariants struct {
	longestSoFar   chain
	amountThatLong int
}

func (cv *chainVariants) addChain(c chain) {
	thisL := len(c.entries)
	longestL := len(cv.longestSoFar.entries)
	if thisL > longestL {
		cv.longestSoFar = c
		cv.amountThatLong = 1
	} else if thisL == longestL {
		cv.amountThatLong += 1
	}
}

type longest struct {
	len    int
	amount int
}

func (cv *chainVariants) longestLen() longest {
	return longest{len(cv.longestSoFar.entries), cv.amountThatLong}
}

type advent9Result struct {
	answer int
	amount int
}

func calcAdvent9Result(inputFile advent9File) advent9Result {
	variants := chainVariants{}
	for currentInputNumberIdx := range inputFile.numbers {
		currentChain := newChain(0, currentInputNumberIdx)

		nextChainEntryInputNumberToTest := currentInputNumberIdx + 1 // start from current point @ input sequence
		for currentChain.add(nextChainEntryInputNumberToTest) {
			nextChainEntryInputNumberToTest = inputFile.numbers[nextChainEntryInputNumberToTest-1]
		}
		variants.addChain(currentChain)
	}

	longestLen := variants.longestLen()
	result := advent9Result{longestLen.len, longestLen.amount}

	return result
}

type advent9File struct {
	amount  int
	numbers []int
}

func readAdvent9File(filename string) advent9File {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(fmt.Errorf("unable to open file %q: %w", filename, err))
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Println(fmt.Errorf("unable to close file %q: %w", filename, err))
		}
	}()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	result := advent9File{}
	for wordIndex := 0; scanner.Scan(); wordIndex++ {
		text := scanner.Text()
		number, err := strconv.Atoi(text)
		if err != nil {
			log.Fatal(fmt.Errorf("NaN %q at index %d: %w", text, wordIndex, err))
		}
		switch wordIndex {
		case 0:
			result.amount = number
			result.numbers = make([]int, 0, number)
		default:
			result.numbers = append(result.numbers, number)
		}
	}

	return result
}
