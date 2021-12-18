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
	log.Printf("Answer: %d", result.value)
}

type advent7Result struct {
	value int64
}

func calcAdvent7File(inputFile advent7File) advent7Result {
	return advent7Result{}
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
		if err != nil { // todo refactor: these file openers should be
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
