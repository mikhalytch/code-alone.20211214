package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	filename := "advent_10.test.txt"
	inputFile := readAdvent10File(filename)
	result := calcAdvent10Result(inputFile)
	log.Printf("Answer: %d\n", result.answer)
}

type advent10Result struct {
	answer int
}

func calcAdvent10Result(inputFile advent10File) advent10Result {
	result := advent10Result{}
	return result
}

type advent10File struct {
	amount int
	sizes  []int
}

func readAdvent10File(filename string) advent10File {
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
	result := advent10File{}
	for wordIndex := 0; scanner.Scan(); wordIndex++ {
		text := scanner.Text()
		number, err := strconv.Atoi(text)
		if err != nil {
			log.Fatal(fmt.Errorf("NaN %q at index %d: %w", text, wordIndex, err))
		}
		switch wordIndex {
		case 0:
			result.amount = number
			result.sizes = make([]int, 0, number)
		default:
			result.sizes = append(result.sizes, number)
		}
	}

	return result
}
