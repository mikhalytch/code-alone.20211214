package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := "advent_7.test.txt"
	inputFile := readAdvent7File(filename)
	result := calcAdvent7File(inputFile)
	log.Printf("Answer: %v (and existing-only: %v)", result.all, result.existing)
}

type advent7Result struct {
	all      int64
	existing int64
}

func numberExistsAnywhere(n64 int64, inputFile advent7File) bool {
	for _, line := range inputFile.lines {
		for _, num := range line.jars {
			if num == n64 {
				return true
			}
		}
	}
	return false
}
func calcAdvent7File(inputFile advent7File) advent7Result {
	result := advent7Result{0, 0}
	type zeroLineNumbered struct {
		idx  int
		line advent7FileLine
	}
	zeroLines := make([]zeroLineNumbered, 0)
	for lineIndex, fileLine := range inputFile.lines {
		if fileLine.jarsAmt == 0 {
			lineNumber := lineIndex + 1
			ln64 := int64(lineNumber)
			result.all += ln64
			if numberExistsAnywhere(ln64, inputFile) {
				zeroLines = append(zeroLines, zeroLineNumbered{lineNumber, fileLine})
				result.existing += ln64
			}
		}
	}
	return result
}

type advent7File struct {
	linesAmt int64
	lines    []advent7FileLine
}

type advent7FileLine struct {
	jarsAmt int64
	jars    []int64
}

func readAdvent7File(filename string) advent7File {
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

	linesScanner := bufio.NewScanner(file)
	linesScanner.Split(bufio.ScanLines)

	result := advent7File{}
	for lineIndex := 0; linesScanner.Scan(); lineIndex++ {
		line := linesScanner.Text()

		switch lineIndex {
		case 0:
			linesAmount := safeConvertToN64(line)
			result.linesAmt = linesAmount
			result.lines = make([]advent7FileLine, 0, linesAmount)
		default:
			wordScanner := bufio.NewScanner(strings.NewReader(line))
			wordScanner.Split(bufio.ScanWords)
			fileLine := advent7FileLine{}
			for wordIndex := 0; wordScanner.Scan(); wordIndex++ {
				n64 := safeConvertToN64(wordScanner.Text())
				switch wordIndex {
				case 0:
					fileLine.jarsAmt = n64
					fileLine.jars = make([]int64, 0, n64)
				default:
					fileLine.jars = append(fileLine.jars, n64)
				}
			}
			result.lines = append(result.lines, fileLine)
		}
	}

	return result
}

func safeConvertToN64(str string) int64 {
	number, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(fmt.Errorf("%v is not a number: %w", number, err))
	}
	n64 := int64(number)
	return n64
}
