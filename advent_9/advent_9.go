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
	fmt.Printf("Answer is: %v\n", result.answer)
}

type chain struct {
	entries map[int]bool
}

func (c *chain) add(new int) bool {
	if c.entries[new] {
		return false
	}
	c.entries[new] = true
	return true
}
func newChain(capacity int) chain {
	c := chain{make(map[int]bool, capacity)}
	return c
}

type chainVariants struct {
	chains []chain
}

func (c *chainVariants) longest() []chain {
	result := make([]chain, 0)
	currentMaxLen := 0
	for _, c := range c.chains {
		l := len(c.entries)
		if l > currentMaxLen {
			result = []chain{c}
			currentMaxLen = l
		} else if l == currentMaxLen {
			result = append(result, c)
		}
	}

	return result
}

func (c *chainVariants) longestLen() int {
	longest := c.longest()
	if len(longest) == 0 {
		return 0
	}
	return len(longest[0].entries)
}

type advent9Result struct {
	answer int
}

func calcAdvent9Result(inputFile advent9File) advent9Result {
	variants := chainVariants{make([]chain, 0, 0)}
	for currentInputNumberIdx := range inputFile.numbers {
		currentChain := newChain(0)

		nextChainEntryInputNumberToTest := currentInputNumberIdx + 1 // start from current point @ input sequence
		for currentChain.add(nextChainEntryInputNumberToTest) {
			nextChainEntryInputNumberToTest = inputFile.numbers[nextChainEntryInputNumberToTest-1]
		}
		variants.chains = append(variants.chains, currentChain)
	}

	result := advent9Result{variants.longestLen()}
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
