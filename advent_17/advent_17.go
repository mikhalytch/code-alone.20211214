package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

const realFilename = "advent_17.test.txt"

func main() {
	inputFile := readAdvent17File(realFilename)
	result := calcAdvent17Result(inputFile)
	log.Printf("Answer: %d", result.answer)
}

type advent17Result struct {
	answer int
}

func calcAdvent17Result(inputFile advent17File) advent17Result {
	return advent17Result{}
}

type advent17File struct {
	symbols []rune
}

func readAdvent17File(filename string) advent17File {
	str, err := os.ReadFile(filename)
	if err != nil {
		log.Println(fmt.Errorf("error opening file %q: %w", filename, err))
	}
	return advent17File{bytes.Runes([]byte(strings.TrimSpace(string(str))))}
}
