package main

import "log"

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
}

func readAdvent17File(filename string) advent17File {
	return advent17File{}
}
