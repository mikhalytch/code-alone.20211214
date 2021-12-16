package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
)

func main() {
	filename := "advent_4.test.txt"
	inputFile := readAdvent4File(filename)
	res := calcAdvent4Result(inputFile)
	log.Println("Answer:", res.amount)
}

type advent4Result struct {
	amount int
}

func calcAdvent4Result(input advent4File) advent4Result {
	result := advent4Result{0}

	found := make([][]int, 0)

	alreadyFound := func(a []int) bool {
		for _, f := range found {
			if reflect.DeepEqual(a, f) {
				return true
			}
		}
		return false
	}

	for index := 0; index < input.lengthsAmt; index++ {
		permutation := make([]int, 0, input.lengthsAmt)

		permutation = append(permutation, input.lengths[index:]...)
		permutation = append(permutation, input.lengths[:index]...)

		if !alreadyFound(permutation) {
			result.amount++
			found = append(found, permutation)
		}
	}

	return result
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
