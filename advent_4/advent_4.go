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

	foundHashes := make(map[int]bool, input.lengthsAmt)
	found := make([][]int, 0, input.lengthsAmt)

	stats := func() string {
		return fmt.Sprintf("Stats: hashes %d, found: %d", len(foundHashes), len(found))
	}
	hash := func(a []int) int {
		res := 0
		for _, aa := range a[:4] {
			res += aa
		}
		return res
	}
	alreadyFound := func(a []int) bool {
		// quick check in hash-array
		ah := hash(a)
		if foundHashes[ah] {
			for _, f := range found {
				if reflect.DeepEqual(a, f) {
					return true
				}
			}
		}
		return false

	}

	double := make([]int, 0, 2*input.lengthsAmt)
	double = append(double, input.lengths...)
	double = append(double, input.lengths...)

	// meaning max length of possible cycle
	maxPossibleAmount := 0
	// meaning starts of possible cycles
	firstNumberEntries := 0
	for _, l := range input.lengths {
		if input.lengths[0] == l {
			firstNumberEntries++
		}
	}
	maxPossibleAmount = input.lengthsAmt / firstNumberEntries

	for index := 0; index < input.lengthsAmt; index++ {
		permutation := double[index : index+input.lengthsAmt]

		if !alreadyFound(permutation) {
			foundHashes[hash(permutation)] = true
			found = append(found, permutation)
		}

		if index%1000 == 0 {
			log.Printf("[Index=%d] %s\n", index, stats())
		}

		if len(found) == maxPossibleAmount {
			break
		}
	}
	result.amount = len(found)
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
