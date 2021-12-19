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
	fmt.Printf("Answer is: %v", result.answer)
}

type advent9Result struct {
	answer int
}

func calcAdvent9Result(inputFile advent9File) advent9Result {
	return advent9Result{}
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
