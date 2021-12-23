package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const realFilename = "advent_13.test.txt"

func main() {
	inputFile := readAdvent13File(realFilename)
	result := calcAdvent13Result(inputFile)
	log.Printf("Answer: %s", result.answer)
}

type advent13Result struct {
	answer string
}

func calcAdvent13Result(inputFile advent13File) advent13Result {
	result := advent13Result{}
	return result
}

type advent13File struct {
	amount, ascNum int
	phones         []string
}

func readAdvent13File(filename string) advent13File {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(fmt.Errorf("error opening file %q: %w", filename, err))
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Println(fmt.Errorf("error closing file %q: %w", filename, err))
		}
	}()

	lineScanner := bufio.NewScanner(file)
	lineScanner.Split(bufio.ScanLines)
	result := advent13File{}
	for lineIdx := 0; lineScanner.Scan(); lineIdx++ {
		lineText := lineScanner.Text()
		switch lineIdx {
		case 0:
			wordScanner := bufio.NewScanner(strings.NewReader(lineText))
			wordScanner.Split(bufio.ScanWords)
			wordScanner.Scan()
			result.amount, _ = strconv.Atoi(wordScanner.Text())
			wordScanner.Scan()
			result.ascNum, _ = strconv.Atoi(wordScanner.Text())
			result.phones = make([]string, 0, result.amount)
		default:
			result.phones = append(result.phones, lineText)
		}
	}
	return result
}
