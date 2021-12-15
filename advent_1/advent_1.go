package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	filename := "advent_1.sample"

	inputFile := readAdvent1File(filename)
	res := calcAdvent1(inputFile)
	log.Printf("Answer: %d clicks\n", res.clicksAmt)
}

type advent1Result struct {
	clicksAmt int
}

func calcAdvent1(inputFile advent1File) advent1Result {
	result := advent1Result{}

	result.clicksAmt = min(2*inputFile.downSwitchesCount+inputFile.upSwitchesCount,
		inputFile.downSwitchesCount+2*inputFile.upSwitchesCount)

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type advent1File struct {
	upSwitchesCount, downSwitchesCount int
}

func readAdvent1File(filename string) advent1File {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(fmt.Errorf("can't open file %q: %w", filename, err))
	}
	defer func() {
		if err2 := file.Close(); err2 != nil {
			fmt.Println(err2)
		}
	}()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	result := advent1File{}

	for index := 0; scanner.Scan(); index++ {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(fmt.Errorf("error converting number at idx %d: %w", index, err))
		}
		switch index {
		case 0:
			result.upSwitchesCount = number
		case 1:
			result.downSwitchesCount = number
		default:
			log.Fatalf("having strange number %d at index %d", number, index)
		}
	}
	return result
}
