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

}

type advent5Result struct {
	code string
}

func calcAdvent5Result(inputFile advent5File) advent5Result {
	return advent5Result{}
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
