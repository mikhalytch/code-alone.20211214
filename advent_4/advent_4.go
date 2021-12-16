package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	filename := "advent_4.test.txt"
	_ = readAdvent4File(filename)
}

type advent4File struct {
	lengthsAmt int
	lengths    []int
}

func readAdvent4File(filename string) advent4File {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(fmt.Errorf("unable to open file %q: %w", filename, err))
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Println(fmt.Errorf("unable to close file %q: %w", filename, err))
		}
	}()

	result := advent4File{}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for index := 0; scanner.Scan(); index++ {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(fmt.Errorf("unable to convert number at index %d: %w", index, err))
		}
		switch index {
		case 0:
			result.lengthsAmt = number
			result.lengths = make([]int, 0, number)
		default:
			result.lengths = append(result.lengths, number)
		}
	}

	return result
}
