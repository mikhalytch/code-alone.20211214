package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	filename := "advent_6.test.txt"
	inputFile := readAdvent6File(filename)
	result := calcAdvent6Result(&inputFile)
	log.Printf("Answer: %s", result.print())
}

type advent6Result advent6File // types are same
func (a6r *advent6Result) print() string {
	return string(a6r.directions)
}

func calcAdvent6Result(inputFile *advent6File) advent6Result {
	return advent6Result{}
}

type advent6File struct {
	directions []rune
}

func readAdvent6File(filename string) advent6File {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(fmt.Errorf("unable to open file %q for reading: %w", filename, err))
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	result := advent6File{}
	result.directions = make([]rune, 0)
	for scanner.Scan() {
		result.directions = append(result.directions, bytes.Runes([]byte(scanner.Text()))[0])
	}

	return result
}
