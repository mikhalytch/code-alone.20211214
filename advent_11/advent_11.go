package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const realFilename = "advent_11.test.txt"

func main() {
	inputFile := readAdvent11File(realFilename)
	result := calcAdvent11Result(inputFile)
	log.Printf("Answer: %d", result.getAnswer())
}

type advent11Result struct {
	sequence []int
}

func (ar *advent11Result) addAll(vs ...int) {
	for _, v := range vs {
		ar.add(v)
	}
}
func (ar *advent11Result) add(v int) {
	ar.sequence = append(ar.sequence, v)
}
func (ar advent11Result) getAnswer() int64 {
	sum := int64(0)
	for idx, v := range ar.sequence {
		number := int64(idx + 1)
		sum += number * int64(v)
	}
	return sum
}

func calcAdvent11Result(inputFile advent11File) advent11Result {
	result := advent11Result{}
	return result
}

type advent11File struct {
	linesAmount int
	lines       []advent11FileLine
}

type advent11FileLine struct {
	screamersAmt    int
	screamerNumbers []int
}

func readAdvent11File(filename string) advent11File {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(fmt.Errorf("error opening file %q: %w", filename, err))
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Println(fmt.Errorf("error closing file %q: %w", filename, err))
		}
	}()

	linesScanner := bufio.NewScanner(file)
	linesScanner.Split(bufio.ScanLines)
	result := advent11File{}

	for lineIdx := 0; linesScanner.Scan(); lineIdx++ {
		line := linesScanner.Text()
		switch lineIdx {
		case 0:
			number, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(fmt.Errorf("NaN %s at lineIdx %d: %w", line, lineIdx, err))
			}
			result.linesAmount = number
			result.lines = make([]advent11FileLine, 0, number)
		default:
			resultLine := advent11FileLine{}
			wordScanner := bufio.NewScanner(strings.NewReader(line))
			wordScanner.Split(bufio.ScanWords)
			for wordIdx := 0; wordScanner.Scan(); wordIdx++ {
				word := wordScanner.Text()
				number, err := strconv.Atoi(word)
				if err != nil {
					log.Fatal(fmt.Errorf("NaN %s at lineIdx:%d wordIdx:%d:%w", word, lineIdx, wordIdx, err))
				}
				switch wordIdx {
				case 0:
					resultLine.screamersAmt = number
					resultLine.screamerNumbers = make([]int, 0, number)
				default:
					resultLine.screamerNumbers = append(resultLine.screamerNumbers, number)
				}
			}
			result.lines = append(result.lines, resultLine)
		}
	}

	return result
}
