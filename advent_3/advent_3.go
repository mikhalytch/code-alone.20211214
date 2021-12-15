package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type advent3file struct {
	ropesAmt, requiredLength int
	ropesLength              []int
}

func main() {
	fileName := "advent_3.test"
	log.Printf("Reading file %q ...\n", fileName)
	input := readFile(fileName)
	log.Printf("File %q contains %d ropes, %d total, and %d lengths\n",
		fileName, input.ropesAmt, input.requiredLength, len(input.ropesLength))
	log.Printf("Sorting %d ints...\n", len(input.ropesLength))
	sort.Sort(sort.Reverse(sort.IntSlice(input.ropesLength)))

	sum, idx := 0, 0
	for ; sum < input.requiredLength; idx++ {
		sum += input.ropesLength[idx]
	}
	log.Printf("sum=%d at index=%d", sum, idx)
}

func readFile(fileName string) advent3file {
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
	}
	return res
}
