package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	fileName := "advent_3.test"

	log.Printf("Reading file %q ...\n", fileName)
	inputFile := readAdvent3File(fileName)
	log.Printf("File %q contains %d ropes, %d total, and %d lengths\n",
		fileName, inputFile.ropesAmt, inputFile.requiredLength, len(inputFile.ropesLength))
	log.Printf("Processing...\n")

	calcResult := calcAdvent3(inputFile)
	log.Printf("Answer is %d (sum=%d at index=%d)", calcResult.index, calcResult.sum, calcResult.index)
}

type advent3Result struct {
	sum, index int
}

func calcAdvent3(input advent3file) advent3Result {
	sort.Sort(sort.Reverse(sort.IntSlice(input.ropesLength)))

	sum, idx := 0, 0
	for ; sum < input.requiredLength; idx++ {
		sum += input.ropesLength[idx]
	}
	return advent3Result{sum, idx - 1}
}

type advent3file struct {
	ropesAmt, requiredLength int
	ropesLength              []int
}

func readAdvent3File(fileName string) advent3file {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err2 := file.Close(); err2 != nil {
			fmt.Println(err2)
		}
	}()

	res := advent3file{}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanWords)

	idx := 0
	for fileScanner.Scan() {
		number, err := strconv.Atoi(fileScanner.Text())
		if err != nil {
			log.Fatal("Can't convert number at index", idx)
		}
		switch idx {
		case 0:
			res.ropesAmt = number
			res.ropesLength = make([]int, 0, number)
		case 1:
			res.requiredLength = number
		default:
			res.ropesLength = append(res.ropesLength, number)
		}
		idx++
	}
	return res
}
