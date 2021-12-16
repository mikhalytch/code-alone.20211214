package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := "advent_5.test.txt"
	inputFile := readAdvent5File(filename)
	result := calcAdvent5Result(inputFile)
	log.Printf("Result (%dth code): %s (and 5th code: %s)\n",
		inputFile.codeNumber, result.requiredCode, result.fifthCode)
}

type advent5Result struct {
	fifthCode    string
	requiredCode string
}

type indexesCalculator struct {
	values    []int
	maxValues []int // exclusive
}

func (ic *indexesCalculator) inc() {
	depth := len(ic.values)

	for level := depth - 1; level >= 0; level-- {
		currLevelValue := ic.values[level]
		newLevelValue := currLevelValue + 1
		if newLevelValue >= ic.maxValues[level] {
			ic.values[level] = 0 // should increase upper level
		} else {
			ic.values[level] = newLevelValue
			break
		}
	}
}
func (ic *indexesCalculator) currentNumber(linesOfNumbers [][]rune) string {
	result := make([]rune, 0)
	for idx, v := range ic.values {
		r := linesOfNumbers[idx][v]
		result = append(result, r)
	}
	return string(result)
}

func calcAdvent5Result(inputFile advent5File) advent5Result {
	requiredIndex := inputFile.codeNumber - 1
	const required5thIndex = 4
	result := advent5Result{}

	cipherSize := inputFile.cipherSize
	calculator := indexesCalculator{make([]int, cipherSize /*init with 0-s*/), make([]int, 0, cipherSize)}
	for _, line := range inputFile.linesOfNumbers {
		calculator.maxValues = append(calculator.maxValues, len(line))
	}
	for currentIndex := 0; currentIndex <= requiredIndex; currentIndex++ {
		if currentIndex == required5thIndex {
			result.fifthCode = calculator.currentNumber(inputFile.linesOfNumbers)
		}
		if currentIndex == requiredIndex {
			result.requiredCode = calculator.currentNumber(inputFile.linesOfNumbers)
		}
		calculator.inc()
	}

	return result
}

type advent5File struct {
	cipherSize, codeNumber int
	linesOfNumbers         [][]rune
}

func readAdvent5File(filename string) advent5File {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(fmt.Errorf("unable to read file %q: %w", filename, err))
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Println(fmt.Errorf("unable to close file %q: %w", filename, err))
		}
	}()

	lineScanner := bufio.NewScanner(file)
	lineScanner.Split(bufio.ScanLines)

	result := advent5File{}
	for lineIndex := 0; lineScanner.Scan(); lineIndex++ {
		wordScanner := bufio.NewScanner(strings.NewReader(lineScanner.Text()))
		wordScanner.Split(bufio.ScanWords)
		for wordIndex := 0; wordScanner.Scan(); wordIndex++ {
			switch lineIndex {
			case 0:
				number, err := strconv.Atoi(wordScanner.Text())
				if err != nil {
					log.Fatal(fmt.Errorf("unable convert to number at rowIdx=%d wordIdx=%d: %w", lineIndex, wordIndex, err))
				}
				switch wordIndex {
				case 0:
					result.cipherSize = number
					//init rune matrix
					result.linesOfNumbers = make([][]rune, 0, number)
					for matrixRow := 0; matrixRow < number; matrixRow++ {
						result.linesOfNumbers = append(result.linesOfNumbers, make([]rune, 0))
					}
				default:
					result.codeNumber = number
				}
			default:
				result.linesOfNumbers[lineIndex-1] =
					append(result.linesOfNumbers[lineIndex-1], bytes.Runes([]byte(wordScanner.Text()))[0])
			}
		}
	}

	return result
}
