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
func (ar *advent11Result) getAnswer() int64 {
	sum := int64(0)
	for idx, v := range ar.sequence {
		number := int64(idx + 1)
		sum += number * int64(v)
	}
	return sum
}

func newSeenScreamerNumbers(capacity int) seenScreamerNumbers {
	return make(seenScreamerNumbers, capacity)
}

type seenScreamerNumbers map[int]bool

func (ssn *seenScreamerNumbers) size() int {
	return len(*ssn)
}
func (ssn *seenScreamerNumbers) isSeen(number int) bool {
	return (*ssn)[number]
}
func (ssn *seenScreamerNumbers) areAllSeen(numbers ...int) bool {
	for _, n := range numbers {
		if !ssn.isSeen(n) {
			return false
		}
	}
	return true
}
func (ssn *seenScreamerNumbers) addNumber(num int) {
	(*ssn)[num] = true
}

func calcAdvent11Result(inputFile advent11File) advent11Result {
	totalScreamers := inputFile.linesAmount
	result := advent11Result{make([]int, totalScreamers)}

	result.sequence = make([]int, 0, totalScreamers)

	aggregator := newSeenScreamerNumbers(totalScreamers)

	for screamerIdx, loopIdx := 0, 0; aggregator.size() < totalScreamers; {
		screamerNumber := screamerIdx + 1

		screamerNumbersToBeRunBefore := inputFile.lines[screamerIdx].screamerNumbers
		if !aggregator.isSeen(screamerNumber) && aggregator.areAllSeen(screamerNumbersToBeRunBefore...) {
			aggregator.addNumber(screamerNumber)
			result.add(screamerNumber)
			screamerIdx = 0
		} else {
			screamerIdx++
		}

		if loopIdx == totalScreamers && aggregator.size() != totalScreamers {
			log.Fatalf("traversed to the last element (num:%d), and still not seen all screamers(seen:%d)!",
				screamerNumber, aggregator.size())
		}
		if screamerIdx == totalScreamers {
			screamerIdx = 0
			loopIdx++
		}
	}

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
